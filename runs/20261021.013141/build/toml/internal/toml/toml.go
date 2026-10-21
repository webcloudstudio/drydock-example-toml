// Package toml parses the TOML values needed by the decoder boundary.
package toml

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Value struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func Decode(document []byte) (map[string]any, error) {
	if !utf8.Valid(document) {
		return nil, fmt.Errorf("invalid UTF-8")
	}
	if strings.Contains(strings.ReplaceAll(string(document), "\r\n", ""), "\r") {
		return nil, fmt.Errorf("carriage return is not allowed")
	}
	for _, r := range string(document) {
		if (r < 0x20 && r != '\t' && r != '\n' && r != '\r') || r == 0x7f {
			return nil, fmt.Errorf("prohibited control character")
		}
	}
	root, current := map[string]any{}, map[string]any{}
	state := parserState{explicit: map[uintptr]bool{}, frozen: map[uintptr]bool{}, dotted: map[uintptr]bool{}}
	current = root
	lines := strings.Split(strings.ReplaceAll(string(document), "\r\n", "\n"), "\n")
	var pending string
	start := 0
	for i, line := range lines {
		if pending == "" {
			start = i + 1
		}
		pending += line
		clean := stripComment(pending)
		if strings.Trim(clean, " \t\n") == "" {
			pending = ""
			continue
		}
		if !complete(clean) {
			pending += "\n"
			continue
		}
		line = strings.Trim(clean, " \t")
		pending = ""
		if strings.Contains(line, "\n") && (strings.HasPrefix(line, "[") || strings.HasPrefix(line, "[[")) {
			return nil, at(start, fmt.Errorf("newline in table header"))
		}
		if strings.HasPrefix(line, "[[") && strings.HasSuffix(line, "]]") {
			p, e := parseKeys(strings.TrimSpace(line[2 : len(line)-2]))
			if e != nil {
				return nil, at(start, e)
			}
			current, e = state.appendTable(root, p)
			if e != nil {
				return nil, at(start, e)
			}
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			p, e := parseKeys(strings.TrimSpace(line[1 : len(line)-1]))
			if e != nil {
				return nil, at(start, e)
			}
			current, e = state.table(root, p)
			if e != nil {
				return nil, at(start, e)
			}
			continue
		}
		pos := assign(line)
		if pos < 0 {
			return nil, at(start, fmt.Errorf("expected key/value assignment"))
		}
		keys, e := parseKeys(strings.Trim(line[:pos], " \t"))
		if e != nil {
			return nil, at(start, e)
		}
		v, e := parseValue(strings.TrimSpace(line[pos+1:]))
		if e != nil {
			return nil, at(start, e)
		}
		if e = state.put(current, keys, v); e != nil {
			return nil, at(start, e)
		}
		state.freezeInline(v)
	}
	if strings.TrimSpace(pending) != "" {
		return nil, at(start, fmt.Errorf("unterminated value"))
	}
	return root, nil
}

type parserState struct {
	explicit map[uintptr]bool
	frozen   map[uintptr]bool
	dotted   map[uintptr]bool
}

func mapID(m map[string]any) uintptr { return reflect.ValueOf(m).Pointer() }

func (s *parserState) freezeInline(v any) {
	switch x := v.(type) {
	case map[string]any:
		s.frozen[mapID(x)] = true
		for _, child := range x {
			s.freezeInline(child)
		}
	case []any:
		for _, child := range x {
			s.freezeInline(child)
		}
	}
}
func at(n int, e error) error { return fmt.Errorf("line %d: %w", n, e) }
func stripComment(s string) string {
	out := []byte(s)
	var q byte
	multi := false
	esc := false
	comment := false
	for i := 0; i < len(out); i++ {
		c := out[i]
		if comment {
			if c == '\n' {
				comment = false
			} else {
				out[i] = ' '
			}
			continue
		}
		if q != 0 {
			if multi && i+2 < len(s) && s[i:i+3] == strings.Repeat(string(q), 3) {
				run := i
				for run < len(s) && s[run] == q {
					run++
				}
				q = 0
				multi = false
				i = run - 1
				continue
			}
			if q == '"' && c == '\\' && !esc {
				esc = true
				continue
			}
			if !multi && ((q == '"' && c == q && !esc) || (q == '\'' && c == q)) {
				q = 0
			}
			esc = false
			continue
		}
		if c == '"' || c == '\'' {
			q = c
			if i+2 < len(s) && s[i:i+3] == strings.Repeat(string(c), 3) {
				multi = true
				i += 2
			}
		} else if c == '#' {
			comment = true
			out[i] = ' '
		}
	}
	return string(out)
}
func complete(s string) bool {
	var q byte
	multi := false
	esc := false
	d := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if q != 0 {
			if q == '"' && c == '\\' && !esc {
				esc = true
				continue
			}
			if multi && i+2 < len(s) && s[i:i+3] == strings.Repeat(string(q), 3) && !esc {
				run := i
				for run < len(s) && s[run] == q {
					run++
				}
				q = 0
				multi = false
				i = run - 1
				continue
			}
			if !multi && ((q == '"' && c == q && !esc) || (q == '\'' && c == q)) {
				q = 0
			}
			esc = false
			continue
		}
		if c == '"' || c == '\'' {
			q = c
			if i+2 < len(s) && s[i:i+3] == strings.Repeat(string(c), 3) {
				multi = true
				i += 2
			}
		} else if c == '[' || c == '{' {
			d++
		} else if c == ']' || c == '}' {
			d--
		}
	}
	return q == 0 && d == 0
}
func assign(s string) int {
	var q byte
	esc := false
	d := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if q != 0 {
			if q == '"' && c == '\\' && !esc {
				esc = true
				continue
			}
			if (q == '"' && c == q && !esc) || (q == '\'' && c == q) {
				q = 0
			}
			esc = false
			continue
		}
		if c == '"' || c == '\'' {
			q = c
		} else if c == '[' || c == '{' {
			d++
		} else if c == ']' || c == '}' {
			d--
		} else if c == '=' && d == 0 {
			return i
		}
	}
	return -1
}
func parseKeys(s string) ([]string, error) {
	var r []string
	for {
		s = strings.Trim(s, " \t")
		if s == "" {
			if len(r) == 0 {
				return nil, fmt.Errorf("empty key")
			}
			return r, nil
		}
		var k string
		if s[0] == '"' || s[0] == '\'' {
			q := s[0]
			end := 1
			for end < len(s) && s[end] != q {
				if q == '"' && s[end] == '\\' {
					end++
					if end >= len(s) {
						return nil, fmt.Errorf("unterminated key")
					}
				}
				end++
			}
			if end >= len(s) {
				return nil, fmt.Errorf("unterminated key")
			}
			rawKey := s[1:end]
			if strings.ContainsAny(rawKey, "\r\n") {
				return nil, fmt.Errorf("newline in key")
			}
			k = rawKey
			if q == '"' {
				var e error
				k, e = decodeBasic(k)
				if e != nil {
					return nil, e
				}
			}
			s = s[end+1:]
		} else {
			i := strings.IndexByte(s, '.')
			if i < 0 {
				i = len(s)
			}
			k = strings.Trim(s[:i], " \t")
			if k == "" {
				return nil, fmt.Errorf("empty key")
			}
			for _, c := range k {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-') {
					return nil, fmt.Errorf("invalid bare key")
				}
			}
			s = s[i:]
		}
		r = append(r, k)
		s = strings.Trim(s, " \t")
		if s == "" {
			return r, nil
		}
		if s[0] != '.' {
			return nil, fmt.Errorf("invalid key separator")
		}
		s = strings.Trim(s[1:], " \t")
		if s == "" {
			return nil, fmt.Errorf("empty key")
		}
	}
}

func (s *parserState) resolveParent(root map[string]any, p []string) (map[string]any, error) {
	m := root
	for _, k := range p {
		if x, ok := m[k]; ok {
			switch v := x.(type) {
			case map[string]any:
				if s.frozen[mapID(v)] {
					return nil, fmt.Errorf("table is closed")
				}
				m = v
			case []any:
				if len(v) == 0 {
					return nil, fmt.Errorf("empty array of tables")
				}
				var ok bool
				m, ok = v[len(v)-1].(map[string]any)
				if !ok {
					return nil, fmt.Errorf("key is not a table")
				}
				if s.frozen[mapID(m)] {
					return nil, fmt.Errorf("table is closed")
				}
			default:
				return nil, fmt.Errorf("key is not a table")
			}
		} else {
			n := map[string]any{}
			m[k] = n
			m = n
		}
	}
	return m, nil
}

func (s *parserState) table(root map[string]any, p []string) (map[string]any, error) {
	if len(p) == 0 {
		return nil, fmt.Errorf("empty table")
	}
	m, e := s.resolveParent(root, p[:len(p)-1])
	if e != nil {
		return nil, e
	}
	k := p[len(p)-1]
	if x, ok := m[k]; ok {
		v, good := x.(map[string]any)
		if !good || s.frozen[mapID(v)] || s.explicit[mapID(v)] || s.dotted[mapID(v)] {
			return nil, fmt.Errorf("table is already defined")
		}
		s.explicit[mapID(v)] = true
		return v, nil
	}
	n := map[string]any{}
	m[k] = n
	s.explicit[mapID(n)] = true
	return n, nil
}

func (s *parserState) appendTable(root map[string]any, p []string) (map[string]any, error) {
	if len(p) == 0 {
		return nil, fmt.Errorf("empty table")
	}
	m, e := s.resolveParent(root, p[:len(p)-1])
	if e != nil {
		return nil, e
	}
	k := p[len(p)-1]
	a, ok := m[k].([]any)
	if m[k] != nil && !ok {
		return nil, fmt.Errorf("key is not an array of tables")
	}
	if ok && len(a) == 0 {
		return nil, fmt.Errorf("key is already defined")
	}
	n := map[string]any{}
	m[k] = append(a, n)
	return n, nil
}

func (s *parserState) put(m map[string]any, p []string, v any) error {
	if len(p) == 0 {
		return fmt.Errorf("empty key")
	}
	for _, k := range p[:len(p)-1] {
		x, ok := m[k]
		if ok {
			n, good := x.(map[string]any)
			if !good || s.frozen[mapID(n)] || s.explicit[mapID(n)] {
				return fmt.Errorf("key is already defined")
			}
			m = n
			continue
		}
		n := map[string]any{}
		m[k] = n
		s.dotted[mapID(n)] = true
		m = n
	}
	k := p[len(p)-1]
	if _, ok := m[k]; ok {
		return fmt.Errorf("duplicate key %q", k)
	}
	m[k] = v
	return nil
}

func put(m map[string]any, p []string, v any) error {
	if len(p) == 0 {
		return fmt.Errorf("empty key")
	}
	for _, k := range p[:len(p)-1] {
		x, ok := m[k]
		if ok {
			var good bool
			m, good = x.(map[string]any)
			if !good {
				return fmt.Errorf("key is already defined")
			}
		} else {
			n := map[string]any{}
			m[k] = n
			m = n
		}
	}
	k := p[len(p)-1]
	if _, ok := m[k]; ok {
		return fmt.Errorf("duplicate key %q", k)
	}
	m[k] = v
	return nil
}

type valueParser struct {
	s string
	i int
}

func parseValue(s string) (any, error) {
	p := &valueParser{s: s}
	v, e := p.value()
	if e == nil && strings.TrimSpace(p.s[p.i:]) != "" {
		e = fmt.Errorf("trailing characters")
	}
	return v, e
}
func (p *valueParser) ws() {
	for p.i < len(p.s) && (p.s[p.i] == ' ' || p.s[p.i] == '\t' || p.s[p.i] == '\n' || p.s[p.i] == '\r') {
		p.i++
	}
}
func (p *valueParser) value() (any, error) {
	p.ws()
	if p.i >= len(p.s) {
		return nil, fmt.Errorf("missing value")
	}
	if p.s[p.i] == '"' || p.s[p.i] == '\'' {
		return p.string()
	}
	if p.s[p.i] == '[' {
		return p.array()
	}
	if p.s[p.i] == '{' {
		return p.inline()
	}
	start := p.i
	for p.i < len(p.s) && !strings.ContainsRune(",]}\r\n", rune(p.s[p.i])) {
		p.i++
	}
	return scalar(strings.TrimSpace(p.s[start:p.i]))
}
func (p *valueParser) string() (any, error) {
	q := p.s[p.i]
	multi := p.i+2 < len(p.s) && p.s[p.i:p.i+3] == strings.Repeat(string(q), 3)
	if multi {
		p.i += 3
	} else {
		p.i++
	}
	start := p.i
	for {
		if p.i >= len(p.s) {
			return nil, fmt.Errorf("unterminated string")
		}
		if multi && p.s[p.i] == q {
			run := p.i
			for run < len(p.s) && p.s[run] == q {
				run++
			}
			if run-p.i >= 3 {
				if run-p.i >= 6 {
					return nil, fmt.Errorf("too many quote characters")
				}
				raw := p.s[start:p.i] + strings.Repeat(string(q), run-p.i-3)
				p.i = run
				if q == '\'' {
					return Value{"string", trimMultiline(raw)}, nil
				}
				v, e := decodeBasic(trimMultiline(raw))
				if e != nil {
					return nil, e
				}
				return Value{"string", v}, nil
			}
		}
		if q == '"' && p.s[p.i] == '\\' && p.i+1 < len(p.s) {
			p.i += 2
			continue
		}
		if !multi && p.s[p.i] == q {
			raw := p.s[start:p.i]
			p.i++
			if q == '\'' {
				return Value{"string", raw}, nil
			}
			v, e := decodeBasic(raw)
			if e != nil {
				return nil, e
			}
			return Value{"string", v}, nil
		}
		if !multi && p.s[p.i] == '\n' {
			return nil, fmt.Errorf("newline in string")
		}
		p.i++
	}
}
func trimMultiline(s string) string {
	if strings.HasPrefix(s, "\n") {
		return s[1:]
	}
	return s
}
func decodeBasic(raw string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(raw); i++ {
		if raw[i] != '\\' {
			b.WriteByte(raw[i])
			continue
		}
		i++
		if i >= len(raw) {
			return "", fmt.Errorf("invalid escape")
		}
		if raw[i] == ' ' || raw[i] == '\t' || raw[i] == '\n' {
			j := i
			for j < len(raw) && (raw[j] == ' ' || raw[j] == '\t' || raw[j] == '\n') {
				j++
			}
			if i < len(raw) && raw[i] == '\n' {
				i = j - 1
				continue
			}
			if j > i && j <= len(raw) {
				if strings.Contains(raw[i:j], "\n") {
					i = j - 1
					continue
				}
				return "", fmt.Errorf("invalid escape")
			}
		}
		switch raw[i] {
		case 'b':
			b.WriteByte('\b')
		case 't':
			b.WriteByte('\t')
		case 'n':
			b.WriteByte('\n')
		case 'f':
			b.WriteByte('\f')
		case 'r':
			b.WriteByte('\r')
		case '"', '\\':
			b.WriteByte(raw[i])
		case 'u', 'U':
			n := 4
			if raw[i] == 'U' {
				n = 8
			}
			if i+n >= len(raw) {
				return "", fmt.Errorf("invalid unicode escape")
			}
			x, e := strconv.ParseUint(raw[i+1:i+1+n], 16, 32)
			if e != nil || x > utf8.MaxRune || (x >= 0xd800 && x <= 0xdfff) {
				return "", fmt.Errorf("invalid unicode escape")
			}
			b.WriteRune(rune(x))
			i += n
		default:
			return "", fmt.Errorf("invalid escape")
		}
	}
	return b.String(), nil
}
func (p *valueParser) array() (any, error) {
	p.i++
	a := make([]any, 0)
	p.ws()
	if p.i < len(p.s) && p.s[p.i] == ']' {
		p.i++
		return a, nil
	}
	for {
		v, e := p.value()
		if e != nil {
			return nil, e
		}
		a = append(a, v)
		p.ws()
		if p.i >= len(p.s) {
			return nil, fmt.Errorf("unterminated array")
		}
		if p.s[p.i] == ']' {
			p.i++
			return a, nil
		}
		if p.s[p.i] != ',' {
			return nil, fmt.Errorf("expected comma")
		}
		p.i++
		p.ws()
		if p.i < len(p.s) && p.s[p.i] == ']' {
			p.i++
			return a, nil
		}
	}
}
func (p *valueParser) inline() (any, error) {
	if inlineHasNewline(p.s, p.i) {
		return nil, fmt.Errorf("newline in inline table")
	}
	p.i++
	m := map[string]any{}
	closed := map[uintptr]bool{}
	p.ws()
	if p.i < len(p.s) && p.s[p.i] == '}' {
		p.i++
		return m, nil
	}
	for {
		start := p.i
		for p.i < len(p.s) && p.s[p.i] != '=' {
			p.i++
		}
		if p.i >= len(p.s) {
			return nil, fmt.Errorf("invalid inline table")
		}
		keys, e := parseKeys(strings.TrimSpace(p.s[start:p.i]))
		if e != nil {
			return nil, e
		}
		p.i++
		v, e := p.value()
		if e != nil {
			return nil, e
		}
		if e = putInline(m, keys, v, closed); e != nil {
			return nil, e
		}
		closeInlineMaps(v, closed)
		p.ws()
		if p.i >= len(p.s) {
			return nil, fmt.Errorf("unterminated inline table")
		}
		if p.s[p.i] == '}' {
			p.i++
			return m, nil
		}
		if p.s[p.i] != ',' {
			return nil, fmt.Errorf("expected comma")
		}
		p.i++
	}
}

func closeInlineMaps(v any, closed map[uintptr]bool) {
	switch x := v.(type) {
	case map[string]any:
		closed[mapID(x)] = true
		for _, child := range x {
			closeInlineMaps(child, closed)
		}
	case []any:
		for _, child := range x {
			closeInlineMaps(child, closed)
		}
	}
}

func putInline(m map[string]any, p []string, v any, closed map[uintptr]bool) error {
	if len(p) == 0 {
		return fmt.Errorf("empty key")
	}
	for _, k := range p[:len(p)-1] {
		x, ok := m[k]
		if ok {
			n, good := x.(map[string]any)
			if !good || closed[mapID(n)] {
				return fmt.Errorf("key is already defined")
			}
			m = n
			continue
		}
		n := map[string]any{}
		m[k] = n
		m = n
	}
	k := p[len(p)-1]
	if _, ok := m[k]; ok {
		return fmt.Errorf("duplicate key %q", k)
	}
	m[k] = v
	return nil
}

func inlineHasNewline(s string, open int) bool {
	depth := 0
	arrayDepth := 0
	var quote byte
	escaped := false
	for i := open; i < len(s); i++ {
		c := s[i]
		if (c == '\n' || c == '\r') && arrayDepth == 0 && quote == 0 {
			return true
		}
		if quote != 0 {
			if i+2 < len(s) && s[i:i+3] == strings.Repeat(string(quote), 3) {
				i += 2
				quote = 0
				continue
			}
			if quote == '"' && c == '\\' && !escaped {
				escaped = true
				continue
			}
			if c == quote && !escaped {
				quote = 0
			}
			escaped = false
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			if i+2 < len(s) && s[i:i+3] == strings.Repeat(string(c), 3) {
				i += 2
			}
		} else if c == '{' {
			depth++
		} else if c == '[' {
			arrayDepth++
		} else if c == ']' {
			arrayDepth--
		} else if c == '}' {
			depth--
			if depth == 0 {
				return false
			}
		}
	}
	return true
}

var (
	datePattern           = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timePattern           = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}(?:\.\d+)?$`)
	localDateTimePattern  = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[Tt ]\d{2}:\d{2}:\d{2}(?:\.\d+)?$`)
	offsetDateTimePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[Tt ]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:Z|z|[+-]\d{2}:\d{2})$`)
)

func temporalValue(s string) (any, bool, error) {
	kind := ""
	canonical := s
	switch {
	case offsetDateTimePattern.MatchString(s):
		kind, canonical = "datetime", normalizeTemporal(s)
	case localDateTimePattern.MatchString(s):
		kind, canonical = "datetime-local", normalizeTemporal(s)
	case datePattern.MatchString(s):
		kind = "date-local"
	case timePattern.MatchString(s):
		kind = "time-local"
	default:
		return nil, false, nil
	}
	if kind == "datetime" && !validOffset(canonical) {
		return nil, true, fmt.Errorf("invalid datetime offset")
	}
	layout := map[string]string{
		"datetime": "2006-01-02T15:04:05.999999999Z07:00", "datetime-local": "2006-01-02T15:04:05.999999999",
		"date-local": "2006-01-02", "time-local": "15:04:05.999999999",
	}[kind]
	if _, err := time.Parse(layout, canonical); err != nil {
		return nil, true, fmt.Errorf("invalid %s", kind)
	}
	return Value{Type: kind, Value: canonical}, true, nil
}

func validOffset(s string) bool {
	if strings.HasSuffix(s, "Z") || strings.HasSuffix(s, "z") {
		return true
	}
	if len(s) < 6 {
		return false
	}
	offset := s[len(s)-6:]
	hour, errHour := strconv.Atoi(offset[1:3])
	minute, errMinute := strconv.Atoi(offset[4:6])
	return (offset[0] == '+' || offset[0] == '-') && offset[3] == ':' &&
		errHour == nil && errMinute == nil && hour <= 23 && minute <= 59
}

func normalizeTemporal(s string) string {
	s = strings.Replace(s, " ", "T", 1)
	s = strings.Replace(s, "t", "T", 1)
	if strings.HasSuffix(s, "z") {
		s = s[:len(s)-1] + "Z"
	}
	dot := strings.IndexByte(s, '.')
	if dot < 0 {
		return s
	}
	end := dot + 1
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if end-dot-1 > 9 {
		s = s[:dot+10] + s[end:]
	}
	return s
}
func scalar(s string) (any, error) {
	if v, matched, err := temporalValue(s); matched {
		return v, err
	}
	if s == "true" || s == "false" {
		return Value{"bool", s}, nil
	}
	unsigned := strings.TrimPrefix(strings.TrimPrefix(s, "+"), "-")
	if strings.HasPrefix(unsigned, "0x") || strings.HasPrefix(unsigned, "0o") || strings.HasPrefix(unsigned, "0b") {
		return intValue(s)
	}
	if s == "inf" || s == "+inf" || s == "-inf" || s == "nan" || s == "+nan" || s == "-nan" || strings.ContainsAny(s, ".eE") {
		return floatValue(s)
	}
	return intValue(s)
}
func validDigits(s string, base int) bool {
	if s == "" || s[0] == '_' || s[len(s)-1] == '_' || strings.Contains(s, "__") {
		return false
	}
	for _, c := range s {
		if c == '_' {
			continue
		}
		ok := c >= '0' && c <= '9'
		if base == 16 {
			ok = ok || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
		}
		if base == 8 {
			ok = c >= '0' && c <= '7'
		}
		if base == 2 {
			ok = c == '0' || c == '1'
		}
		if !ok {
			return false
		}
	}
	return true
}
func intValue(s string) (any, error) {
	sign := ""
	body := s
	if strings.HasPrefix(body, "+") || strings.HasPrefix(body, "-") {
		sign = body[:1]
		body = body[1:]
	}
	base := 10
	if len(body) > 2 && body[0] == '0' {
		if body[1] == 'x' {
			base = 16
		}
		if body[1] == 'o' {
			base = 8
		}
		if body[1] == 'b' {
			base = 2
		}
	}
	digits := body
	if base != 10 {
		if sign != "" {
			return nil, fmt.Errorf("invalid integer")
		}
		digits = body[2:]
	}
	if base == 10 && len(digits) > 1 && digits[0] == '0' {
		return nil, fmt.Errorf("invalid integer")
	}
	if !validDigits(digits, base) || !validUnderscores(digits, base) {
		return nil, fmt.Errorf("invalid integer")
	}
	clean := strings.ReplaceAll(s, "_", "")
	if base != 10 {
		clean = sign + strings.ReplaceAll(digits, "_", "")
	}
	if _, e := strconv.ParseInt(clean, base, 64); e != nil {
		return nil, fmt.Errorf("invalid integer")
	}
	n, _ := strconv.ParseInt(clean, base, 64)
	return Value{"integer", strconv.FormatInt(n, 10)}, nil
}
func floatValue(s string) (any, error) {
	clean := strings.ReplaceAll(s, "_", "")
	if !validUnderscores(s, 10) {
		return nil, fmt.Errorf("invalid float")
	}
	if s == "inf" || s == "+inf" || s == "-inf" || s == "nan" || s == "+nan" || s == "-nan" {
		return Value{"float", s}, nil
	}
	number := strings.TrimPrefix(strings.TrimPrefix(clean, "+"), "-")
	dot := strings.IndexByte(number, '.')
	if !strings.ContainsAny(clean, ".eE") || strings.HasSuffix(number, ".") || strings.HasPrefix(number, ".") || (dot >= 0 && (dot+1 == len(number) || number[dot+1] == 'e' || number[dot+1] == 'E')) {
		return nil, fmt.Errorf("invalid float")
	}
	if len(number) > 1 && number[0] == '0' && number[1] >= '0' && number[1] <= '9' {
		return nil, fmt.Errorf("invalid float")
	}
	if _, e := strconv.ParseFloat(clean, 64); e != nil {
		return nil, fmt.Errorf("invalid float")
	}
	return Value{"float", s}, nil
}

func validUnderscores(s string, base int) bool {
	for i := range s {
		if s[i] != '_' {
			continue
		}
		if i == 0 || i+1 == len(s) || !digitForBase(s[i-1], base) || !digitForBase(s[i+1], base) {
			return false
		}
	}
	return !strings.Contains(s, "__")
}

func digitForBase(c byte, base int) bool {
	if c >= '0' && c <= '9' {
		return int(c-'0') < base
	}
	return base == 16 && ((c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F'))
}
