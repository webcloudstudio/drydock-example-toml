// Package toml provides the TOML decoder used by the command-line filter.
package toml

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type scalar struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Decode parses a TOML document and returns its tagged representation.
func Decode(input string) (map[string]any, error) {
	if !utf8.ValidString(input) {
		return nil, fmt.Errorf("invalid UTF-8")
	}
	var err error
	input, err = normalizeLexicalBoundaries(input)
	if err != nil {
		return nil, err
	}
	result := map[string]any{}
	current := result
	definedTables := map[uintptr]bool{}
	sealedTables := map[uintptr]bool{}
	inlineTables := map[uintptr]bool{}
	staticArrays := map[uintptr]bool{}
	pos := 0
	line := 1
	for pos < len(input) {
		start := pos
		for pos < len(input) && input[pos] != '\n' {
			pos++
		}
		if pos < len(input) {
			pos++
		}
		physical := input[start:pos]
		if strings.TrimSpace(physical) == "" || strings.HasPrefix(strings.TrimSpace(physical), "#") {
			line++
			continue
		}
		trimmedLine := strings.TrimSpace(physical)
		if strings.HasPrefix(trimmedLine, "[") {
			if strings.HasPrefix(trimmedLine, "[[") {
				if !strings.HasSuffix(strings.TrimSpace(trimmedLine), "]]") {
					return nil, fmt.Errorf("line %d: invalid array table", line)
				}
				path, err := parseKeyPath(strings.TrimSpace(trimmedLine[2 : len(strings.TrimSpace(trimmedLine))-2]))
				if err != nil {
					return nil, fmt.Errorf("line %d: %w", line, err)
				}
				container, err := tableContainer(result, path[:len(path)-1], staticArrays, inlineTables)
				if err != nil {
					return nil, fmt.Errorf("line %d: %w", line, err)
				}
				name := path[len(path)-1]
				array, exists := container[name].([]any)
				if exists {
					if staticArrays[reflect.ValueOf(array).Pointer()] {
						return nil, fmt.Errorf("line %d: table conflicts with static array", line)
					}
					// An array of tables may be extended repeatedly. A regular
					// table or scalar at the same path cannot become an array.
				} else if _, exists := container[name]; exists {
					return nil, fmt.Errorf("line %d: table conflicts with value", line)
				}
				child := map[string]any{}
				array = append(array, child)
				container[name] = array
				current = child
				line += strings.Count(physical, "\n")
				continue
			}
			if !strings.HasSuffix(trimmedLine, "]") {
				return nil, fmt.Errorf("line %d: invalid table", line)
			}
			path, err := parseKeyPath(strings.TrimSpace(trimmedLine[1 : len(trimmedLine)-1]))
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", line, err)
			}
			current, err = tablePath(result, path, staticArrays, sealedTables, inlineTables)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", line, err)
			}
			identity := reflect.ValueOf(current).Pointer()
			if definedTables[identity] {
				return nil, fmt.Errorf("line %d: table redefinition", line)
			}
			definedTables[identity] = true
			line += strings.Count(physical, "\n")
			continue
		}
		eq := assignmentEquals(physical)
		if eq < 0 {
			return nil, fmt.Errorf("line %d: expected key-value pair", line)
		}
		keys, err := parseKeyPath(strings.TrimSpace(physical[:eq]))
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		valueText := input[start+eq+1:]
		lineValue := physical[eq+1:]
		if strings.Contains(lineValue, "{") && !strings.Contains(lineValue, "}") && !strings.Contains(lineValue, "[") && !strings.Contains(lineValue, "\"\"\"") && !strings.Contains(lineValue, "'''") {
			return nil, fmt.Errorf("line %d: newline in inline table", line)
		}
		value, consumed, err := parseValue(valueText, input, start+eq+1)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		lineEnd := strings.IndexByte(valueText, '\n')
		if lineEnd < 0 {
			lineEnd = len(valueText)
		}
		if consumed > lineEnd {
			lineEnd = consumed
		}
		trailing := strings.TrimSpace(valueText[consumed:lineEnd])
		if trailing != "" && !isComment(trailing) {
			return nil, fmt.Errorf("line %d: trailing characters after value", line)
		}
		if err := setPath(current, keys, value, definedTables, sealedTables, inlineTables); err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		markStaticArrays(value, staticArrays)
		markClosedTables(value, inlineTables)
		pos = start + eq + 1 + consumed
		line += strings.Count(input[start:pos], "\n")
	}
	return result, nil
}

func assignmentEquals(line string) int {
	quote := byte(0)
	for i := 0; i < len(line); i++ {
		if quote != 0 {
			if quote == '"' && line[i] == '\\' {
				i++
				continue
			}
			if line[i] == quote {
				quote = 0
			}
			continue
		}
		if line[i] == '\'' || line[i] == '"' {
			quote = line[i]
			continue
		}
		if line[i] == '=' {
			return i
		}
	}
	return -1
}

func parseKeyPath(text string) ([]string, error) {
	var keys []string
	for len(strings.TrimSpace(text)) > 0 {
		text = strings.TrimSpace(text)
		var key string
		var rest string
		quoted := false
		switch text[0] {
		case '\'', '"':
			quoted = true
			end, value, err := parseQuoted(text, 0, text[0] == '"')
			if err != nil {
				return nil, fmt.Errorf("invalid key: %w", err)
			}
			key, rest = value, text[end:]
		default:
			i := 0
			for i < len(text) && (text[i] >= 'a' && text[i] <= 'z' || text[i] >= 'A' && text[i] <= 'Z' || text[i] >= '0' && text[i] <= '9' || text[i] == '_' || text[i] == '-') {
				i++
			}
			if i == 0 {
				return nil, fmt.Errorf("invalid or empty key")
			}
			key, rest = text[:i], text[i:]
		}
		if key == "" && !quoted {
			return nil, fmt.Errorf("invalid or empty key")
		}
		keys = append(keys, key)
		rest = strings.TrimSpace(rest)
		if rest == "" {
			return keys, nil
		}
		if rest[0] != '.' {
			return nil, fmt.Errorf("invalid key separator")
		}
		text = rest[1:]
	}
	return nil, fmt.Errorf("empty key")
}

func parseValue(text, full string, absolute int) (any, int, error) {
	return parseValueMode(text, full, absolute, false)
}

func parseValueMode(text, full string, absolute int, allowNewline bool) (any, int, error) {
	if strings.HasPrefix(strings.TrimLeft(text, " \t"), "\n") || strings.HasPrefix(strings.TrimLeft(text, " \t"), "\r") {
		return nil, 0, fmt.Errorf("missing value")
	}
	trimmed := strings.TrimLeft(text, " \t\r\n")
	offset := len(text) - len(trimmed)
	if trimmed == "" {
		return nil, 0, fmt.Errorf("missing value")
	}
	if strings.HasPrefix(trimmed, "\"\"\"") || strings.HasPrefix(trimmed, "'''") {
		literal := trimmed[0] == '\''
		value, end, err := parseMultiline(trimmed, literal)
		return value, offset + end, err
	}
	if trimmed[0] == '{' {
		end, err := matchingDelimiter(trimmed, '{', '}')
		if err != nil {
			return nil, 0, fmt.Errorf("unterminated inline table")
		}
		body := strings.TrimSpace(trimmed[1:end])
		if inlineTableHasNewline(body) {
			return nil, 0, fmt.Errorf("newline in inline table")
		}
		if strings.ContainsAny(body, "\r\n") && !strings.Contains(body, "\"\"\"") && !strings.Contains(body, "'''") && !strings.Contains(body, "[") {
			return nil, 0, fmt.Errorf("newline in inline table")
		}
		table := map[string]any{}
		if body != "" {
			closed := map[uintptr]bool{}
			for _, item := range splitTopLevel(body, ',') {
				eq := strings.IndexByte(item, '=')
				if eq < 0 {
					return nil, 0, fmt.Errorf("invalid inline table")
				}
				keys, err := parseKeyPath(strings.TrimSpace(item[:eq]))
				if err != nil {
					return nil, 0, err
				}
				value, _, err := parseValue(strings.TrimSpace(item[eq+1:]), "", 0)
				if err != nil {
					return nil, 0, err
				}
				if err := setPath(table, keys, value, closed); err != nil {
					return nil, 0, err
				}
				markClosedTables(value, closed)
			}
		}
		return table, offset + end + 1, nil
	}
	if trimmed[0] == '[' {
		end, err := matchingDelimiter(trimmed, '[', ']')
		if err != nil {
			return nil, 0, err
		}
		body := trimmed[1:end]
		values := []any{}
		for pos := 0; ; {
			for pos < len(body) && (body[pos] == ' ' || body[pos] == '\t' || body[pos] == '\r' || body[pos] == '\n') {
				pos++
			}
			if pos == len(body) {
				break
			}
			value, used, valueErr := parseValueMode(body[pos:], full, absolute+offset+1+pos, true)
			if valueErr != nil {
				return nil, 0, valueErr
			}
			values = append(values, value)
			pos += used
			for pos < len(body) && (body[pos] == ' ' || body[pos] == '\t' || body[pos] == '\r' || body[pos] == '\n') {
				pos++
			}
			if pos == len(body) {
				break
			}
			if body[pos] != ',' {
				return nil, 0, fmt.Errorf("invalid array")
			}
			pos++
		}
		return values, offset + end + 1, nil
	}
	if trimmed[0] == '\'' || trimmed[0] == '"' {
		end, value, err := parseQuoted(trimmed, 0, trimmed[0] == '"')
		return scalar{Type: "string", Value: value}, offset + end, err
	}
	end := len(trimmed)
	if !allowNewline {
		if newline := strings.IndexAny(trimmed, "\r\n"); newline >= 0 {
			end = newline
		}
	} else {
		if comma := strings.IndexByte(trimmed, ','); comma >= 0 {
			end = comma
		}
	}
	for i, r := range trimmed {
		if r == '#' && i < end {
			end = i
			break
		}
	}
	word := strings.TrimSpace(trimmed[:end])
	if word == "true" || word == "false" {
		return scalar{Type: "bool", Value: word}, offset + end, nil
	}
	if typ, canonical := temporalValue(word); typ != "" {
		return scalar{Type: typ, Value: canonical}, offset + end, nil
	}
	if isInteger(word) {
		canonical, err := canonicalInteger(word)
		if err != nil {
			return nil, 0, err
		}
		return scalar{Type: "integer", Value: canonical}, offset + end, nil
	}
	if isFloat(word) {
		return scalar{Type: "float", Value: word}, offset + end, nil
	}
	return nil, 0, fmt.Errorf("invalid value %q", word)
}

var (
	datePattern    = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timePattern    = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}(?:\.\d+)?$`)
	localDateTime  = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[Tt ]\d{2}:\d{2}:\d{2}(?:\.\d+)?$`)
	offsetDateTime = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[Tt ]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:[Zz]|[+-]\d{2}:\d{2})$`)
)

func temporalValue(value string) (string, string) {
	normalized := strings.Replace(strings.Replace(value, "t", "T", 1), "z", "Z", 1)
	normalized = truncateFraction(normalized)
	switch {
	case offsetDateTime.MatchString(normalized):
		if !validOffset(normalized) {
			return "", ""
		}
		parsed, err := time.Parse("2006-01-02T15:04:05Z07:00", strings.Replace(normalized, " ", "T", 1))
		if err == nil {
			return "datetime", parsed.Format(time.RFC3339Nano)
		}
	case localDateTime.MatchString(normalized):
		parsed, err := time.ParseInLocation("2006-01-02T15:04:05", strings.Replace(normalized, " ", "T", 1), time.UTC)
		if err == nil {
			return "datetime-local", parsed.Format("2006-01-02T15:04:05.999999999")
		}
	case datePattern.MatchString(normalized):
		parsed, err := time.Parse("2006-01-02", normalized)
		if err == nil {
			return "date-local", parsed.Format("2006-01-02")
		}
	case timePattern.MatchString(normalized):
		parsed, err := time.Parse("15:04:05", normalized)
		if err == nil {
			return "time-local", parsed.Format("15:04:05.999999999")
		}
	}
	return "", ""
}

func truncateFraction(value string) string {
	dot := strings.IndexByte(value, '.')
	if dot < 0 {
		return value
	}
	end := dot + 1
	for end < len(value) && value[end] >= '0' && value[end] <= '9' {
		end++
	}
	if end-dot-1 <= 9 {
		return value
	}
	return value[:dot+10] + value[end:]
}

func pathKey(path []string) string { return strings.Join(path, "\x00") }

// tableContainer resolves the parent of a header, selecting the latest element
// whenever a parent is an array of tables.
func tableContainer(root map[string]any, path []string, staticArrays, inline map[uintptr]bool) (map[string]any, error) {
	current := root
	for _, key := range path {
		value, ok := current[key]
		if !ok {
			next := map[string]any{}
			current[key] = next
			current = next
			continue
		}
		switch next := value.(type) {
		case map[string]any:
			if inline[reflect.ValueOf(next).Pointer()] {
				return nil, fmt.Errorf("table conflicts with inline table")
			}
			current = next
		case []any:
			if staticArrays[reflect.ValueOf(next).Pointer()] {
				return nil, fmt.Errorf("table conflicts with static array")
			}
			if len(next) == 0 {
				return nil, fmt.Errorf("invalid empty table array")
			}
			last, ok := next[len(next)-1].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid table path")
			}
			current = last
		default:
			return nil, fmt.Errorf("table conflicts with value")
		}
	}
	return current, nil
}

func tablePath(root map[string]any, path []string, staticArrays map[uintptr]bool, sealed, inline map[uintptr]bool) (map[string]any, error) {
	if len(path) == 0 {
		return root, nil
	}
	parent, err := tableContainer(root, path[:len(path)-1], staticArrays, inline)
	if err != nil {
		return nil, err
	}
	last := path[len(path)-1]
	if value, ok := parent[last]; ok {
		if _, isArray := value.([]any); isArray {
			return nil, fmt.Errorf("table conflicts with array")
		}
		if table, ok := value.(map[string]any); ok {
			if sealed[reflect.ValueOf(table).Pointer()] || inline[reflect.ValueOf(table).Pointer()] {
				return nil, fmt.Errorf("table was defined by dotted key")
			}
			return table, nil
		}
		return nil, fmt.Errorf("table conflicts with value")
	}
	table := map[string]any{}
	parent[last] = table
	return table, nil
}

func matchingDelimiter(text string, open, close byte) (int, error) {
	depth := 0
	quote := byte(0)
	for i := 0; i < len(text); i++ {
		if quote != 0 {
			if quote == '"' && text[i] == '\\' {
				i++
				continue
			}
			if text[i] == quote {
				quote = 0
			}
			continue
		}
		if text[i] == '\'' || text[i] == '"' {
			quote = text[i]
			continue
		}
		switch text[i] {
		case open:
			depth++
		case close:
			depth--
			if depth == 0 {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("unterminated array")
}

func isInteger(word string) bool {
	text := word
	if len(text) > 0 && (text[0] == '+' || text[0] == '-') {
		text = text[1:]
	}
	if text == "" {
		return false
	}
	base := 10
	signed := len(text) != len(word)
	if len(text) >= 2 && text[0] == '0' {
		switch text[1] {
		case 'x':
			base, text = 16, text[2:]
		case 'o':
			base, text = 8, text[2:]
		case 'b':
			base, text = 2, text[2:]
		}
	}
	if signed && base != 10 {
		return false
	}
	if text == "" || !validDigits(text, base) {
		return false
	}
	if base == 10 && len(text) > 1 && text[0] == '0' {
		return false
	}
	return true
}

func validDigits(text string, base int) bool {
	if text == "" || text[0] == '_' || text[len(text)-1] == '_' {
		return false
	}
	previousUnderscore := false
	for _, r := range text {
		if r == '_' {
			if previousUnderscore {
				return false
			}
			previousUnderscore = true
			continue
		}
		value := -1
		switch {
		case r >= '0' && r <= '9':
			value = int(r - '0')
		case r >= 'a' && r <= 'f':
			value = int(r-'a') + 10
		case r >= 'A' && r <= 'F':
			value = int(r-'A') + 10
		}
		if value < 0 || value >= base {
			return false
		}
		previousUnderscore = false
	}
	return true
}

func canonicalInteger(word string) (string, error) {
	text := strings.ReplaceAll(word, "_", "")
	sign := ""
	if text[0] == '+' || text[0] == '-' {
		sign, text = text[:1], text[1:]
	}
	base := 10
	if len(text) >= 2 && text[0] == '0' {
		switch text[1] {
		case 'x':
			base, text = 16, text[2:]
		case 'o':
			base, text = 8, text[2:]
		case 'b':
			base, text = 2, text[2:]
		}
	}
	value, err := strconv.ParseUint(text, base, 64)
	if err != nil || (sign != "-" && value > uint64(^uint64(0)>>1)) || (sign == "-" && value > uint64(1)<<63) {
		return "", fmt.Errorf("integer out of range")
	}
	if value == 0 {
		return "0", nil
	}
	canonical := strconv.FormatUint(value, 10)
	if sign == "-" {
		return "-" + canonical, nil
	}
	return canonical, nil
}

func isFloat(word string) bool {
	if word == "inf" || word == "+inf" || word == "-inf" || word == "nan" || word == "+nan" || word == "-nan" {
		return true
	}
	text := word
	if len(text) > 0 && (text[0] == '+' || text[0] == '-') {
		text = text[1:]
	}
	if text == "" {
		return false
	}
	exponent := strings.IndexAny(text, "eE")
	if exponent >= 0 {
		if strings.IndexAny(text[exponent+1:], "eE") >= 0 {
			return false
		}
		exp := text[exponent+1:]
		if len(exp) > 0 && (exp[0] == '+' || exp[0] == '-') {
			exp = exp[1:]
		}
		if !validDigits(exp, 10) {
			return false
		}
		text = text[:exponent]
	}
	dot := strings.IndexByte(text, '.')
	if dot >= 0 {
		if strings.IndexByte(text[dot+1:], '.') >= 0 || !validDigits(text[:dot], 10) || !validDigits(text[dot+1:], 10) {
			return false
		}
		if len(text[:dot]) > 1 && text[0] == '0' {
			return false
		}
		return true
	}
	return exponent >= 0 && validDigits(text, 10)
}

func parseQuoted(s string, start int, basic bool) (int, string, error) {
	quote := s[start]
	var b strings.Builder
	for i := start + 1; i < len(s); i++ {
		if s[i] == quote {
			return i + 1, b.String(), nil
		}
		if s[i] == '\n' || s[i] == '\r' || (s[i] < 0x20 && s[i] != '\t') || s[i] == 0x7f {
			return 0, "", fmt.Errorf("control character in string")
		}
		if basic && s[i] == '\\' {
			value, end, err := decodeEscape(s, i)
			if err != nil {
				return 0, "", err
			}
			b.WriteString(value)
			i = end - 1
			continue
		}
		b.WriteByte(s[i])
	}
	return 0, "", fmt.Errorf("unterminated string")
}

func parseMultiline(s string, literal bool) (scalar, int, error) {
	delim := s[:3]
	end := -1
	for i := 3; i < len(s); {
		if !literal && s[i] == '\\' {
			if i+1 >= len(s) {
				break
			}
			if s[i+1] == '\r' && i+2 < len(s) && s[i+2] == '\n' {
				i += 3
			} else {
				i += 2
			}
			continue
		}
		if s[i] != delim[0] {
			i++
			continue
		}
		run := 0
		for i+run < len(s) && s[i+run] == delim[0] {
			run++
		}
		if run >= 3 {
			if run > 5 {
				return scalar{}, 0, fmt.Errorf("invalid multiline delimiter")
			}
			end = i + run - 3
			break
		}
		i += run
	}
	if end < 0 {
		return scalar{}, 0, fmt.Errorf("unterminated multiline string")
	}
	body := s[3:end]
	if strings.HasPrefix(body, "\r\n") {
		body = body[2:]
	} else if strings.HasPrefix(body, "\n") {
		body = body[1:]
	}
	if literal {
		if strings.ContainsAny(body, "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0b\x0c\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f") {
			return scalar{}, 0, fmt.Errorf("control character in string")
		}
		return scalar{Type: "string", Value: body}, end + 3, nil
	}
	decoded, err := decodeBasic(body)
	return scalar{Type: "string", Value: decoded}, end + 3, err
}

func decodeBasic(s string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			j := i + 1
			for j < len(s) && (s[j] == ' ' || s[j] == '\t') {
				j++
			}
			if j < len(s) && (s[j] == '\n' || s[j] == '\r') {
				i = j
				if s[i] == '\r' && i+1 < len(s) && s[i+1] == '\n' {
					i++
				}
				for i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\t' || s[i+1] == '\n' || s[i+1] == '\r') {
					i++
				}
				continue
			}
		}
		if s[i] == '\\' {
			value, end, err := decodeEscape(s, i)
			if err != nil {
				return "", err
			}
			b.WriteString(value)
			i = end - 1
			continue
		}
		if (s[i] < 0x20 && s[i] != '\n' && s[i] != '\t' && s[i] != '\r') || s[i] == 0x7f {
			return "", fmt.Errorf("control character in string")
		}
		b.WriteByte(s[i])
	}
	return b.String(), nil
}

func decodeEscape(s string, i int) (string, int, error) {
	if i+1 >= len(s) {
		return "", 0, fmt.Errorf("invalid escape")
	}
	short := map[byte]string{'b': "\b", 't': "\t", 'n': "\n", 'f': "\f", 'r': "\r", '"': "\"", '\\': "\\"}
	if value, ok := short[s[i+1]]; ok {
		return value, i + 2, nil
	}
	width := 0
	if s[i+1] == 'u' {
		width = 4
	}
	if s[i+1] == 'U' {
		width = 8
	}
	if width == 0 || i+2+width > len(s) {
		return "", 0, fmt.Errorf("invalid escape")
	}
	code, err := strconv.ParseUint(s[i+2:i+2+width], 16, 32)
	if err != nil || code > 0x10ffff || code >= 0xd800 && code <= 0xdfff {
		return "", 0, fmt.Errorf("invalid Unicode escape")
	}
	return string(rune(code)), i + 2 + width, nil
}

func escaped(s string, i int) bool {
	n := 0
	for i > 0 && s[i-1] == '\\' {
		n++
		i--
	}
	return n%2 == 1
}
func isComment(s string) bool { return strings.HasPrefix(s, "#") }

// normalizeLexicalBoundaries removes comments while preserving quoted content and line endings.
func normalizeLexicalBoundaries(input string) (string, error) {
	var b strings.Builder
	b.Grow(len(input))
	quote := byte(0)
	multiline := false
	for i := 0; i < len(input); i++ {
		c := input[i]
		if quote != 0 {
			b.WriteByte(c)
			if quote == '"' && c == '\\' && i+1 < len(input) {
				i++
				b.WriteByte(input[i])
				continue
			}
			if multiline {
				if i+2 < len(input) && input[i:i+3] == string([]byte{quote, quote, quote}) {
					b.WriteByte(input[i+1])
					b.WriteByte(input[i+2])
					i += 2
					quote = 0
					multiline = false
				}
			} else if c == quote {
				quote = 0
			}
			continue
		}
		if c == '\'' || c == '"' {
			quote = c
			if i+2 < len(input) && input[i:i+3] == string([]byte{c, c, c}) {
				b.WriteByte(input[i+1])
				b.WriteByte(input[i+2])
				i += 2
				multiline = true
			}
			b.WriteByte(c)
			continue
		}
		if c == '#' {
			for i < len(input) && input[i] != '\n' && input[i] != '\r' {
				if input[i] != '\t' && input[i] < 0x20 || input[i] == 0x7f {
					return "", fmt.Errorf("control character in comment")
				}
				b.WriteByte(' ')
				i++
			}
			if i < len(input) {
				b.WriteByte(input[i])
			}
			continue
		}
		if c == '\r' && (i+1 == len(input) || input[i+1] != '\n') {
			return "", fmt.Errorf("bare carriage return")
		}
		if (c < 0x20 && c != '\t' && c != '\n' && c != '\r') || c == 0x7f {
			return "", fmt.Errorf("control character in document")
		}
		if c >= utf8.RuneSelf {
			r, width := utf8.DecodeRuneInString(input[i:])
			if unicode.IsSpace(r) {
				return "", fmt.Errorf("invalid whitespace in document")
			}
			for j := 1; j < width; j++ {
				b.WriteByte(input[i+j])
			}
			i += width - 1
		}
		b.WriteByte(c)
	}
	return b.String(), nil
}

func setPath(root map[string]any, keys []string, value any, trackers ...map[uintptr]bool) error {
	var defined map[uintptr]bool
	var sealed map[uintptr]bool
	var closed map[uintptr]bool
	if len(trackers) > 0 {
		defined = trackers[0]
	}
	if len(trackers) > 1 {
		sealed = trackers[1]
	}
	if len(trackers) > 2 {
		closed = trackers[2]
	}
	current := root
	for _, key := range keys[:len(keys)-1] {
		existing, ok := current[key]
		if !ok {
			child := map[string]any{}
			current[key] = child
			if sealed != nil {
				sealed[reflect.ValueOf(child).Pointer()] = true
			}
			current = child
			continue
		}
		child, ok := existing.(map[string]any)
		if !ok {
			return fmt.Errorf("cannot extend scalar key %q", key)
		}
		if defined != nil && defined[reflect.ValueOf(child).Pointer()] {
			return fmt.Errorf("cannot extend explicitly defined table %q", key)
		}
		if closed != nil && closed[reflect.ValueOf(child).Pointer()] {
			return fmt.Errorf("cannot extend inline table %q", key)
		}
		current = child
	}
	last := keys[len(keys)-1]
	if _, exists := current[last]; exists {
		return fmt.Errorf("duplicate key %q", last)
	}
	current[last] = value
	return nil
}

func splitTopLevel(text string, separator byte) []string {
	var parts []string
	start, depth := 0, 0
	quote := byte(0)
	for i := 0; i < len(text); i++ {
		if quote != 0 {
			if quote == '"' && text[i] == '\\' {
				i++
			} else if text[i] == quote {
				quote = 0
			}
			continue
		}
		if text[i] == '\'' || text[i] == '"' {
			quote = text[i]
			continue
		}
		switch text[i] {
		case '{', '[':
			depth++
		case '}', ']':
			depth--
		case separator:
			if depth == 0 {
				parts = append(parts, text[start:i])
				start = i + 1
			}
		}
	}
	return append(parts, text[start:])
}

func inlineTableHasNewline(text string) bool {
	depth := 0
	quote := byte(0)
	triple := false
	for i := 0; i < len(text); i++ {
		if quote != 0 {
			if triple && i+2 < len(text) && text[i:i+3] == string([]byte{quote, quote, quote}) {
				i += 2
				quote = 0
				triple = false
				continue
			}
			if quote == '"' && text[i] == '\\' {
				i++
			} else if text[i] == quote {
				quote = 0
			}
			continue
		}
		if text[i] == '\'' || text[i] == '"' {
			quote = text[i]
			if i+2 < len(text) && text[i:i+3] == string([]byte{text[i], text[i], text[i]}) {
				triple = true
				i += 2
			}
		} else if text[i] == '[' {
			depth++
		} else if text[i] == ']' {
			depth--
		} else if (text[i] == '\n' || text[i] == '\r') && depth == 0 {
			return true
		}
	}
	return false
}

func markStaticArrays(value any, arrays map[uintptr]bool) {
	if values, ok := value.([]any); ok {
		arrays[reflect.ValueOf(values).Pointer()] = true
	}
}

func markClosedTables(value any, closed map[uintptr]bool) {
	switch value := value.(type) {
	case map[string]any:
		closed[reflect.ValueOf(value).Pointer()] = true
		for _, child := range value {
			markClosedTables(child, closed)
		}
	case []any:
		for _, child := range value {
			markClosedTables(child, closed)
		}
	}
}

func validOffset(value string) bool {
	if len(value) < 6 || (value[len(value)-6] != '+' && value[len(value)-6] != '-') {
		return true
	}
	hour, _ := strconv.Atoi(value[len(value)-5 : len(value)-3])
	minute, _ := strconv.Atoi(value[len(value)-2:])
	return hour <= 23 && minute <= 59
}
