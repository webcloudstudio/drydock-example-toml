// Package toml parses TOML documents into the decoder's tagged JSON model.
package toml

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Value is a scalar TOML value represented in the JSON contract.
type Value struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Document is the root TOML table.
type Document map[string]any

type parser struct {
	src       string
	pos       int
	root      Document
	current   []string
	defined   map[string]bool
	assigned  map[string]bool
	arrayTabs map[string]bool
	inline    map[uintptr]bool
	line      int
}

// Decode parses a TOML document and returns its tagged representation.
func Decode(input string) (Document, error) {
	if !utf8.ValidString(input) {
		return nil, fmt.Errorf("invalid UTF-8")
	}
	for i, r := range input {
		if r == '\r' && (i+1 == len(input) || input[i+1] != '\n') {
			return nil, fmt.Errorf("invalid carriage return")
		}
		if (r < 0x20 && r != '\t' && r != '\n' && r != '\r') || r == 0x7f {
			return nil, fmt.Errorf("invalid control character")
		}
	}
	p := &parser{src: input, root: make(Document), defined: make(map[string]bool), assigned: make(map[string]bool), arrayTabs: make(map[string]bool), inline: make(map[uintptr]bool), line: 1}
	if err := p.document(); err != nil {
		return nil, fmt.Errorf("line %d: %w", p.line, err)
	}
	return p.root, nil
}

func (p *parser) document() error {
	for {
		p.skipSpace(true)
		if p.pos >= len(p.src) {
			return nil
		}
		if p.src[p.pos] == '[' {
			isArray := strings.HasPrefix(p.src[p.pos:], "[[")
			if err := p.header(isArray); err != nil {
				return err
			}
			continue
		}
		keys, err := p.keys('=', true)
		if err != nil {
			return err
		}
		if len(keys) == 0 {
			return fmt.Errorf("invalid key")
		}
		if err := p.assign(keys); err != nil {
			return err
		}
	}
}

func (p *parser) header(array bool) error {
	if array {
		p.pos += 2
	} else {
		p.pos++
	}
	start := p.pos
	close := "]"
	if array {
		close = "]]"
	}
	end := findClosing(p.src, p.pos, close[0])
	if end < 0 || (array && (end+1 >= len(p.src) || p.src[end+1] != ']')) {
		return fmt.Errorf("unterminated table header")
	}
	text := trimHorizontal(p.src[start:end])
	p.pos = end + len(close)
	keys, err := parseKeyList(text)
	if err != nil || len(keys) == 0 {
		return fmt.Errorf("invalid table name")
	}
	if err := p.finishLine(); err != nil {
		return err
	}
	if array {
		if err := p.addArrayTable(keys); err != nil {
			return err
		}
	} else if err := p.openTable(keys); err != nil {
		return err
	}
	return nil
}

func (p *parser) assign(keys []string) error {
	if p.pos >= len(p.src) || p.src[p.pos] != '=' {
		return fmt.Errorf("expected key and value")
	}
	p.pos++
	for p.pos < len(p.src) && (p.src[p.pos] == ' ' || p.src[p.pos] == '\t') {
		p.pos++
	}
	if p.pos < len(p.src) && (p.src[p.pos] == '\n' || p.src[p.pos] == '\r') {
		return fmt.Errorf("missing value")
	}
	v, err := p.value(false)
	if err != nil {
		return err
	}
	if err := p.finishLine(); err != nil {
		return err
	}
	base, err := p.tableAt(p.current, false)
	if err != nil {
		return err
	}
	for i, key := range keys {
		if i == len(keys)-1 {
			full := append(append([]string{}, p.current...), keys...)
			if p.assigned[pathString(full)] {
				return fmt.Errorf("duplicate key %q", key)
			}
			if _, ok := base[key]; ok {
				return fmt.Errorf("duplicate key %q", key)
			}
			base[key] = v
			p.assigned[pathString(full)] = true
			continue
		}
		prefix := append(append([]string{}, p.current...), keys[:i+1]...)
		if p.defined[pathString(prefix)] {
			return fmt.Errorf("cannot extend defined table with dotted key")
		}
		n, ok := base[key]
		if !ok {
			n = make(map[string]any)
			base[key] = n
		} else if isInline(n) {
			return fmt.Errorf("cannot extend inline table")
		}
		m, ok := n.(map[string]any)
		if !ok {
			return fmt.Errorf("cannot extend value as table")
		}
		base = m
	}
	return nil
}

func (p *parser) openTable(path []string) error {
	key := pathString(path)
	if p.defined[key] {
		return fmt.Errorf("redefined table")
	}
	if len(path) > 0 {
		parent, err := p.tableAt(path[:len(path)-1], true)
		if err != nil {
			return err
		}
		name := path[len(path)-1]
		if old, ok := parent[name]; ok {
			if _, array := old.([]any); array {
				return fmt.Errorf("table/array conflict")
			}
			for assigned := range p.assigned {
				parts := strings.Split(assigned, "\x00")
				if len(path) > 1 && len(parts) > len(path) && pathString(parts[:len(path)]) == key {
					return fmt.Errorf("redefined table")
				}
			}
		}
	}
	if _, err := p.tableAt(path, true); err != nil {
		return err
	}
	p.defined[key] = true
	p.current = path
	return nil
}

func (p *parser) addArrayTable(path []string) error {
	parent, err := p.tableAt(path[:len(path)-1], true)
	if err != nil {
		return err
	}
	name := path[len(path)-1]
	if old, ok := parent[name]; ok {
		if _, isArray := old.([]any); !isArray || !p.arrayTabs[pathString(path)] {
			return fmt.Errorf("table/array conflict")
		}
	} else {
		parent[name] = []any{}
	}
	a := parent[name].([]any)
	item := make(map[string]any)
	parent[name] = append(a, item)
	p.arrayTabs[pathString(path)] = true
	prefix := pathString(path)
	for k := range p.assigned {
		if k == prefix || strings.HasPrefix(k, prefix+"\x00") {
			delete(p.assigned, k)
		}
	}
	for k := range p.defined {
		if k == prefix || strings.HasPrefix(k, prefix+"\x00") {
			delete(p.defined, k)
		}
	}
	p.current = path
	return nil
}

func (p *parser) tableAt(path []string, create bool) (map[string]any, error) {
	cur := map[string]any(p.root)
	sofar := []string{}
	for i, key := range path {
		sofar = append(sofar, key)
		n, ok := cur[key]
		if !ok {
			if !create {
				return nil, fmt.Errorf("table not found")
			}
			m := make(map[string]any)
			cur[key] = m
			cur = m
			continue
		}
		if isInline(n) {
			return nil, fmt.Errorf("cannot extend inline table")
		}
		if a, ok := n.([]any); ok {
			if !p.arrayTabs[pathString(sofar)] {
				return nil, fmt.Errorf("array is not a table")
			}
			if len(a) == 0 {
				return nil, fmt.Errorf("empty table array")
			}
			m, ok := a[len(a)-1].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid table array")
			}
			cur = m
			continue
		}
		m, ok := n.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("value used as table")
		}
		if i == len(path)-1 {
			cur = m
		} else {
			cur = m
		}
	}
	return cur, nil
}

func (p *parser) value(inline bool) (any, error) {
	p.skipSpace(!inline)
	if p.pos >= len(p.src) {
		return nil, fmt.Errorf("missing value")
	}
	switch p.src[p.pos] {
	case '"':
		if strings.HasPrefix(p.src[p.pos:], "\"\"\"") {
			v, n, e := parseMultiline(p.src, p.pos, '"')
			p.pos = n
			return Value{"string", v}, e
		}
		v, n, e := parseQuoted(p.src, p.pos, '"')
		p.pos = n
		return Value{"string", v}, e
	case '\'':
		if strings.HasPrefix(p.src[p.pos:], "'''") {
			v, n, e := parseMultiline(p.src, p.pos, '\'')
			p.pos = n
			return Value{"string", v}, e
		}
		v, n, e := parseQuoted(p.src, p.pos, '\'')
		p.pos = n
		return Value{"string", v}, e
	case '[':
		return p.array()
	case '{':
		return p.inlineTable()
	}
	start := p.pos
	for p.pos < len(p.src) && !strings.ContainsRune("\t\r\n,#]}", rune(p.src[p.pos])) {
		p.pos++
	}
	raw := strings.TrimSpace(p.src[start:p.pos])
	return parseScalarOrDate(raw)
}

func (p *parser) array() (any, error) {
	p.pos++
	result := []any{}
	for {
		p.skipSpace(true)
		if p.pos >= len(p.src) {
			return nil, fmt.Errorf("unterminated array")
		}
		if p.src[p.pos] == ']' {
			p.pos++
			return result, nil
		}
		v, err := p.value(false)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
		p.skipSpace(true)
		if p.pos < len(p.src) && p.src[p.pos] == ',' {
			p.pos++
			continue
		}
		if p.pos < len(p.src) && p.src[p.pos] == ']' {
			p.pos++
			return result, nil
		}
		return nil, fmt.Errorf("expected comma or closing bracket")
	}
}

func (p *parser) inlineTable() (any, error) {
	p.pos++
	m := make(map[string]any)
	for {
		p.skipSpace(false)
		if p.pos >= len(p.src) {
			return nil, fmt.Errorf("unterminated inline table")
		}
		if p.src[p.pos] == '}' {
			p.pos++
			inlinePointers[reflect.ValueOf(m).Pointer()] = true
			return m, nil
		}
		keys, err := p.keys('=', false)
		if err != nil {
			return nil, err
		}
		if len(keys) == 0 {
			return nil, fmt.Errorf("invalid key")
		}
		if p.pos >= len(p.src) || p.src[p.pos] != '=' {
			return nil, fmt.Errorf("expected equals")
		}
		p.pos++
		v, err := p.value(true)
		if err != nil {
			return nil, err
		}
		base := m
		for i, k := range keys {
			if i == len(keys)-1 {
				if _, ok := base[k]; ok {
					return nil, fmt.Errorf("duplicate key")
				}
				base[k] = v
				break
			}
			n, ok := base[k]
			if !ok {
				n = map[string]any{}
				base[k] = n
			} else if isInline(n) {
				return nil, fmt.Errorf("cannot extend inline table")
			}
			var yes bool
			base, yes = n.(map[string]any)
			if !yes {
				return nil, fmt.Errorf("invalid dotted inline key")
			}
		}
		p.skipSpace(false)
		if p.pos < len(p.src) && p.src[p.pos] == ',' {
			p.pos++
			p.skipSpace(false)
			if p.pos < len(p.src) && p.src[p.pos] == '}' {
				return nil, fmt.Errorf("trailing comma in inline table")
			}
			continue
		}
		if p.pos < len(p.src) && p.src[p.pos] == '}' {
			p.pos++
			inlinePointers[reflect.ValueOf(m).Pointer()] = true
			return m, nil
		}
		return nil, fmt.Errorf("expected comma or closing brace")
	}
}

func (p *parser) keys(separator byte, line bool) ([]string, error) {
	start := p.pos
	quote := byte(0)
	depth := 0
	for p.pos < len(p.src) {
		c := p.src[p.pos]
		if quote != 0 {
			if quote == '"' && c == '\\' {
				p.pos += 2
				continue
			}
			if c == quote {
				quote = 0
			}
			p.pos++
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			p.pos++
			continue
		}
		if c == '[' || c == '{' {
			depth++
		}
		if c == ']' || c == '}' {
			depth--
		}
		if c == separator && depth == 0 {
			break
		}
		if line && c == '\n' {
			break
		}
		p.pos++
	}
	text := trimHorizontal(p.src[start:p.pos])
	return parseKeyList(text)
}

func parseKeyList(s string) ([]string, error) {
	var out []string
	for len(trimHorizontal(s)) > 0 {
		s = trimHorizontal(s)
		var key string
		var n int
		var e error
		if s[0] == '"' || s[0] == '\'' {
			key, n, e = parseQuoted(s, 0, s[0])
		} else {
			n = 0
			for n < len(s) && s[n] != '.' && s[n] != '=' && s[n] != ' ' && s[n] != '\t' {
				n++
			}
			key = s[:n]
			if !validBareKey(key) {
				e = fmt.Errorf("invalid key")
			}
		}
		if e != nil {
			return nil, e
		}
		out = append(out, key)
		s = s[n:]
		s = trimHorizontal(s)
		if s == "" {
			break
		}
		if s[0] != '.' {
			return nil, fmt.Errorf("invalid dotted key")
		}
		s = s[1:]
		if trimHorizontal(s) == "" {
			return nil, fmt.Errorf("invalid dotted key")
		}
	}
	return out, nil
}

func trimHorizontal(s string) string { return strings.Trim(s, " \t") }

func (p *parser) finishLine() error {
	p.skipSpace(false)
	if p.pos < len(p.src) && p.src[p.pos] == '#' {
		for p.pos < len(p.src) && p.src[p.pos] != '\n' {
			p.pos++
		}
	}
	if p.pos < len(p.src) && p.src[p.pos] != '\n' {
		return fmt.Errorf("trailing content")
	}
	return nil
}
func (p *parser) skipSpace(newlines bool) {
	for p.pos < len(p.src) {
		switch p.src[p.pos] {
		case ' ', '\t', '\r':
			p.pos++
		case '\n':
			if !newlines {
				return
			}
			p.pos++
			p.line++
		case '#':
			if !newlines {
				return
			}
			for p.pos < len(p.src) && p.src[p.pos] != '\n' {
				p.pos++
			}
		default:
			return
		}
	}
}
func isInline(v any) bool {
	m, ok := v.(map[string]any)
	return ok && false || ok && inlinePointers[reflect.ValueOf(m).Pointer()]
}

var inlinePointers = map[uintptr]bool{}

func pathString(p []string) string { return strings.Join(p, "\x00") }
func findClosing(s string, start int, c byte) int {
	q := byte(0)
	for i := start; i < len(s); i++ {
		if q != 0 {
			if q == '"' && s[i] == '\\' {
				i++
				continue
			}
			if s[i] == q {
				q = 0
			}
			continue
		}
		if s[i] == '"' || s[i] == '\'' {
			q = s[i]
			continue
		}
		if s[i] == c {
			return i
		}
	}
	return -1
}

var dateRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var timeRe = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}(?:\.\d+)?$`)
var dateTimeRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[Tt ]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:[Zz]|[+-]\d{2}:\d{2})?$`)

func parseScalarOrDate(raw string) (Value, error) {
	if dateRe.MatchString(raw) {
		if _, e := time.Parse("2006-01-02", raw); e != nil {
			return Value{}, fmt.Errorf("invalid date")
		}
		return Value{"date-local", raw}, nil
	}
	if timeRe.MatchString(raw) {
		clean := truncateFraction(raw)
		if _, e := time.Parse("15:04:05.999999999", clean); e != nil {
			return Value{}, fmt.Errorf("invalid time")
		}
		return Value{"time-local", clean}, nil
	}
	if dateTimeRe.MatchString(raw) {
		if v, e := parseDateTime(raw); e == nil {
			return v, nil
		}
	}
	return parseScalar(raw)
}
func truncateFraction(s string) string {
	if i := strings.IndexByte(s, '.'); i >= 0 {
		j := i + 1
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}
		if j-i-1 > 9 {
			s = s[:i+10] + s[j:]
		}
	}
	return s
}
func parseDateTime(raw string) (Value, error) {
	if z := strings.LastIndexAny(raw, "+-"); z > 10 && strings.Contains(raw[z:], ":") {
		off := raw[z+1:]
		if len(off) != 5 || off[2] != ':' || off[0] > '2' || off[0] < '0' || off[1] > '9' || off[1] < '0' || off[3] > '5' || off[3] < '0' || off[4] > '9' || off[4] < '0' {
			return Value{}, fmt.Errorf("invalid datetime")
		}
	}
	s := strings.Replace(raw, " ", "T", 1)
	s = strings.Replace(s, "t", "T", 1)
	if strings.HasSuffix(s, "z") {
		s = s[:len(s)-1] + "Z"
	}
	s = truncateFraction(s)
	layouts := []string{"2006-01-02T15:04:05Z07:00", "2006-01-02T15:04:05.999999999Z07:00", "2006-01-02T15:04:05", "2006-01-02T15:04:05.999999999"}
	for _, l := range layouts {
		if t, e := time.Parse(l, s); e == nil {
			if strings.HasSuffix(s, "Z") || strings.Contains(s[19:], "+") || strings.Contains(s[19:], "-") {
				return Value{"datetime", t.Format(time.RFC3339Nano)}, nil
			}
			return Value{"datetime-local", t.Format("2006-01-02T15:04:05.999999999")}, nil
		}
	}
	return Value{}, fmt.Errorf("invalid datetime")
}

func parseScalar(raw string) (Value, error) {
	if raw == "true" || raw == "false" {
		return Value{"bool", raw}, nil
	}
	if raw == "inf" || raw == "+inf" || raw == "-inf" || raw == "nan" || raw == "+nan" || raw == "-nan" {
		return Value{"float", raw}, nil
	}
	clean := strings.ReplaceAll(raw, "_", "")
	if raw == "" || !validUnderscores(raw) {
		return Value{}, fmt.Errorf("invalid scalar")
	}
	if strings.HasPrefix(clean, "0x") || strings.HasPrefix(clean, "0o") || strings.HasPrefix(clean, "0b") {
		base := 16
		if clean[1] == 'o' {
			base = 8
		}
		if clean[1] == 'b' {
			base = 2
		}
		n, e := strconv.ParseUint(clean[2:], base, 64)
		if e != nil || n > math.MaxInt64 || !validBaseUnderscores(raw, base) {
			return Value{}, fmt.Errorf("invalid integer")
		}
		return Value{"integer", strconv.FormatUint(n, 10)}, nil
	}
	if strings.ContainsAny(clean, ".eE") {
		if _, e := strconv.ParseFloat(clean, 64); e != nil || !validFloat(raw) {
			return Value{}, fmt.Errorf("invalid float")
		}
		return Value{"float", raw}, nil
	}
	n, e := strconv.ParseInt(clean, 10, 64)
	digits := strings.TrimPrefix(strings.TrimPrefix(clean, "+"), "-")
	if e != nil || !validDecimalUnderscores(raw) || (len(digits) > 1 && digits[0] == '0') {
		return Value{}, fmt.Errorf("invalid integer")
	}
	return Value{"integer", strconv.FormatInt(n, 10)}, nil
}
func parseQuoted(s string, pos int, q byte) (string, int, error) {
	var b strings.Builder
	for i := pos + 1; i < len(s); i++ {
		if s[i] == q {
			return b.String(), i + 1, nil
		}
		if s[i] == '\n' || s[i] == '\r' {
			return "", i, fmt.Errorf("unterminated string")
		}
		if q == '\'' {
			b.WriteByte(s[i])
			continue
		}
		if s[i] == '\\' {
			v, n, e := escape(s, i+1)
			if e != nil {
				return "", i, e
			}
			b.WriteString(v)
			i = n
			continue
		}
		_, n := utf8.DecodeRuneInString(s[i:])
		b.WriteString(s[i : i+n])
		i += n - 1
	}
	return "", len(s), fmt.Errorf("unterminated string")
}
func escape(s string, i int) (string, int, error) {
	if i >= len(s) {
		return "", i, fmt.Errorf("invalid escape")
	}
	short := map[byte]string{'b': "\b", 't': "\t", 'n': "\n", 'f': "\f", 'r': "\r", '"': "\"", '\\': "\\"}
	if v, ok := short[s[i]]; ok {
		return v, i, nil
	}
	n := 4
	if s[i] == 'U' {
		n = 8
	} else if s[i] != 'u' {
		return "", i, fmt.Errorf("invalid escape")
	}
	if i+n >= len(s) {
		return "", i, fmt.Errorf("invalid Unicode escape")
	}
	code, e := strconv.ParseUint(s[i+1:i+n+1], 16, 32)
	if e != nil || code > utf8.MaxRune || code >= 0xd800 && code <= 0xdfff {
		return "", i, fmt.Errorf("invalid Unicode escape")
	}
	return string(rune(code)), i + n, nil
}
func parseMultiline(s string, pos int, q byte) (string, int, error) {
	i := pos + 3
	if i < len(s) && s[i] == '\r' {
		i++
	}
	if i < len(s) && s[i] == '\n' {
		i++
	}
	var b strings.Builder
	for i < len(s) {
		if s[i] == q {
			run := i
			for run < len(s) && s[run] == q {
				run++
			}
			if run-i >= 3 {
				if run-i >= 6 {
					return "", i, fmt.Errorf("invalid multiline quote sequence")
				}
				b.WriteString(strings.Repeat(string(q), run-i-3))
				return b.String(), run, nil
			}
		}
		if q == '"' && s[i] == '\\' {
			j := i + 1
			for j < len(s) && (s[j] == ' ' || s[j] == '\t') {
				j++
			}
			if j < len(s) && (s[j] == '\n' || s[j] == '\r') {
				i = j
				if s[i] == '\r' {
					i++
				}
				if i < len(s) && s[i] == '\n' {
					i++
				}
				for i < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
					i++
				}
				continue
			}
			v, n, e := escape(s, i+1)
			if e != nil {
				return "", i, e
			}
			b.WriteString(v)
			i = n + 1
			continue
		}
		if s[i] == '\r' {
			if i+1 < len(s) && s[i+1] == '\n' {
				i++
			}
			b.WriteByte('\n')
			i++
			continue
		}
		if s[i] < 0x20 && s[i] != '\n' && s[i] != '\t' || s[i] == 0x7f {
			return "", i, fmt.Errorf("invalid control character in string")
		}
		b.WriteByte(s[i])
		i++
	}
	return "", len(s), fmt.Errorf("unterminated string")
}
func validBareKey(k string) bool {
	if k == "" {
		return false
	}
	for _, r := range k {
		if !(r == '_' || r == '-' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9') {
			return false
		}
	}
	return true
}
func validUnderscores(s string) bool {
	for i := range s {
		if s[i] == '_' && (i == 0 || i+1 == len(s) || !isAlphaNum(s[i-1]) || !isAlphaNum(s[i+1])) {
			return false
		}
	}
	return true
}
func isAlphaNum(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}
func validDecimalUnderscores(s string) bool {
	for i := range s {
		if s[i] == '_' && (i == 0 || i+1 == len(s) || s[i-1] < '0' || s[i-1] > '9' || s[i+1] < '0' || s[i+1] > '9') {
			return false
		}
	}
	return true
}
func validBaseUnderscores(s string, base int) bool {
	if len(s) < 3 || s[2] == '_' {
		return false
	}
	digit := func(c byte) bool {
		switch base {
		case 2:
			return c == '0' || c == '1'
		case 8:
			return c >= '0' && c <= '7'
		default:
			return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
		}
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '_' && (i+1 == len(s) || !digit(s[i-1]) || !digit(s[i+1])) {
			return false
		}
		if s[i] != '_' && !digit(s[i]) {
			return false
		}
	}
	return true
}
func validFloat(s string) bool {
	clean := strings.TrimLeft(s, "+-")
	if e := strings.IndexAny(clean, "eE"); e >= 0 {
		clean = clean[:e]
	}
	if strings.HasSuffix(clean, ".") || strings.HasPrefix(clean, ".") || strings.HasPrefix(clean, "0") && len(clean) > 1 && clean[1] != '.' {
		return false
	}
	return validDecimalUnderscores(s)
}
