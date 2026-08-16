// Package toml contains the TOML decoder and its tagged value model.
package toml

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Value is the JSON representation of a TOML scalar value.
type Value struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Document is the JSON representation of a TOML document.
type Document map[string]any

// Decode parses a TOML document and returns its tagged representation.
func Decode(input string) (Document, error) {
	if !utf8.ValidString(input) {
		return nil, fmt.Errorf("invalid UTF-8")
	}
	if err := validateInputBytes(input); err != nil {
		return nil, err
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")
	state := parserState{document: Document{}, explicit: map[uintptr]bool{}, dotted: map[uintptr]bool{}, closed: map[uintptr]bool{}, arrays: map[arrayKey]bool{}}
	current := map[string]any(state.document)
	lines := strings.Split(input, "\n")
	for lineNumber := 0; lineNumber < len(lines); lineNumber++ {
		line := lines[lineNumber]
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if strings.HasPrefix(trimmed, "[[") {
			name, err := headerName(trimmed, true)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
			}
			item, err := state.defineArrayTable(name)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
			}
			current = item
			continue
		}
		if strings.HasPrefix(trimmed, "[") {
			name, err := headerName(trimmed, false)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
			}
			current, err = state.defineTable(name)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
			}
			continue
		}
		eq := assignmentIndex(line)
		if eq < 0 {
			return nil, fmt.Errorf("line %d: expected key-value pair", lineNumber+1)
		}
		key := strings.TrimSpace(line[:eq])
		if key == "" {
			return nil, fmt.Errorf("line %d: invalid key-value pair", lineNumber+1)
		}
		raw := strings.TrimLeft(line[eq+1:], " \t")
		var value any
		var err error
		if strings.HasPrefix(raw, "\"\"\"") || strings.HasPrefix(raw, "'''") {
			quote := raw[:3]
			parts := []string{raw}
			for !hasMultilineEnd(strings.Join(parts, "\n"), quote) {
				lineNumber++
				if lineNumber >= len(lines) {
					return nil, fmt.Errorf("line %d: unterminated multiline string", lineNumber)
				}
				parts = append(parts, lines[lineNumber])
			}
			raw = strings.Join(parts, "\n")
			value, _, err = decodeString(raw)
		} else {
			for hasLineContinuation(raw) || delimitersOpen(raw) {
				lineNumber++
				if lineNumber >= len(lines) {
					return nil, fmt.Errorf("line %d: unterminated line continuation", lineNumber)
				}
				raw += "\n" + lines[lineNumber]
			}
			raw = stripComment(raw)
			cleaned := strings.TrimSpace(raw)
			var end int
			value, end, err = decodeValue(cleaned)
			if err == nil && strings.TrimSpace(cleaned[end:]) != "" {
				err = fmt.Errorf("text after value")
			}
		}
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
		}
		keyParts, keyErr := parseKey(key)
		if keyErr != nil {
			return nil, fmt.Errorf("line %d: %w", lineNumber+1, keyErr)
		}
		if len(keyParts) > 1 {
			parent, err := state.assignmentParent(current, keyParts[:len(keyParts)-1])
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
			}
			leaf := keyParts[len(keyParts)-1]
			if _, exists := parent[leaf]; exists {
				return nil, fmt.Errorf("line %d: duplicate key %q", lineNumber+1, key)
			}
			parent[keyParts[len(keyParts)-1]] = value
			state.markClosed(value)
			continue
		}
		if _, exists := current[keyParts[0]]; exists {
			return nil, fmt.Errorf("line %d: duplicate key %q", lineNumber+1, key)
		}
		current[keyParts[0]] = value
		state.markClosed(value)
	}
	return state.document, nil
}

func (p *parserState) markClosed(value any) {
	switch v := value.(type) {
	case map[string]any:
		p.explicit[mapID(v)] = true
		p.closed[mapID(v)] = true
		for _, child := range v {
			p.markClosed(child)
		}
	case []any:
		for _, child := range v {
			p.markClosed(child)
		}
	}
}

type arrayKey struct {
	parent uintptr
	name   string
}

type parserState struct {
	document Document
	explicit map[uintptr]bool
	dotted   map[uintptr]bool
	closed   map[uintptr]bool
	arrays   map[arrayKey]bool
}

func mapID(m map[string]any) uintptr {
	return reflect.ValueOf(m).Pointer()
}

func (p *parserState) defineTable(path []string) (map[string]any, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("invalid table path")
	}
	parent, leaf, err := p.parentFor(path)
	if err != nil {
		return nil, err
	}
	value, exists := parent[leaf]
	if !exists {
		value = map[string]any{}
		parent[leaf] = value
	}
	m, ok := value.(map[string]any)
	if !ok || p.explicit[mapID(m)] || p.dotted[mapID(m)] {
		return nil, fmt.Errorf("table redefinition")
	}
	p.explicit[mapID(m)] = true
	return m, nil
}

func (p *parserState) defineArrayTable(path []string) (map[string]any, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("invalid table path")
	}
	parent, leaf, err := p.parentFor(path)
	if err != nil {
		return nil, err
	}
	key := arrayKey{parent: mapID(parent), name: leaf}
	value, exists := parent[leaf]
	if !exists {
		item := map[string]any{}
		parent[leaf] = []any{item}
		p.arrays[key] = true
		p.explicit[mapID(item)] = true
		return item, nil
	}
	arr, ok := value.([]any)
	if !ok || !p.arrays[key] {
		return nil, fmt.Errorf("table redefinition")
	}
	item := map[string]any{}
	parent[leaf] = append(arr, item)
	p.explicit[mapID(item)] = true
	return item, nil
}

func (p *parserState) parentFor(path []string) (map[string]any, string, error) {
	cur := map[string]any(p.document)
	for _, part := range path[:len(path)-1] {
		value, exists := cur[part]
		if !exists {
			m := map[string]any{}
			cur[part] = m
			cur = m
			continue
		}
		switch v := value.(type) {
		case map[string]any:
			if p.closed[mapID(v)] {
				return nil, "", fmt.Errorf("table redefinition")
			}
			cur = v
		case []any:
			if !p.arrays[arrayKey{parent: mapID(cur), name: part}] || len(v) == 0 {
				return nil, "", fmt.Errorf("table redefinition")
			}
			item, ok := v[len(v)-1].(map[string]any)
			if !ok {
				return nil, "", fmt.Errorf("table redefinition")
			}
			cur = item
		default:
			return nil, "", fmt.Errorf("table redefinition")
		}
	}
	return cur, path[len(path)-1], nil
}

func (p *parserState) assignmentParent(current map[string]any, path []string) (map[string]any, error) {
	cur := current
	for _, part := range path {
		value, exists := cur[part]
		if !exists {
			m := map[string]any{}
			cur[part] = m
			p.dotted[mapID(m)] = true
			cur = m
			continue
		}
		switch v := value.(type) {
		case map[string]any:
			if p.closed[mapID(v)] || p.explicit[mapID(v)] {
				return nil, fmt.Errorf("table redefinition")
			}
			cur = v
		case []any:
			return nil, fmt.Errorf("table redefinition")
		default:
			return nil, fmt.Errorf("table redefinition")
		}
	}
	return cur, nil
}

// validateInputBytes applies the document-wide lexical restrictions before line
// parsing can accidentally treat a forbidden byte as whitespace or content.
func validateInputBytes(input string) error {
	lineHasToken := false
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch {
		case c == '\n':
			lineHasToken = false
			continue
		case c == '\t' || c == ' ':
			continue
		case c == '\r':
			if i+1 < len(input) && input[i+1] == '\n' {
				i++
				lineHasToken = false
				continue
			}
			return fmt.Errorf("invalid carriage return")
		case c < 0x20 || c == 0x7f:
			return fmt.Errorf("invalid control character")
		}
		if c >= utf8.RuneSelf {
			r, width := utf8.DecodeRuneInString(input[i:])
			if r == '\u3000' && !lineHasToken || r == '\ufeff' || r == '\ufffd' {
				return fmt.Errorf("invalid non-ASCII whitespace or encoding marker")
			}
			i += width - 1
		}
		lineHasToken = true
	}
	return nil
}

func hasMultilineEnd(s, quote string) bool {
	_, end, err := decodeString(s)
	return err == nil && end > 0 && end <= len(s) && validCommentSuffix(s[end:])
}

func validCommentSuffix(s string) bool {
	s = strings.TrimLeft(s, " \t")
	return s == "" || strings.HasPrefix(s, "#")
}

func decodeValue(raw string) (any, int, error) {
	if strings.HasPrefix(raw, "[") {
		return decodeArray(raw)
	}
	if strings.HasPrefix(raw, "\"") || strings.HasPrefix(raw, "'") {
		return decodeString(raw)
	}
	if strings.HasPrefix(raw, "{") {
		return decodeInlineTable(raw)
	}
	return decodeScalar(raw)
}

func decodeArray(raw string) ([]any, int, error) {
	if len(raw) < 2 || raw[0] != '[' {
		return nil, 0, fmt.Errorf("invalid array")
	}
	items := []any{}
	i := 1
	for {
		for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t' || raw[i] == '\n') {
			i++
		}
		if i < len(raw) && raw[i] == ']' {
			return items, i + 1, nil
		}
		if i >= len(raw) {
			return nil, 0, fmt.Errorf("unterminated array")
		}
		end := elementEnd(raw[i:])
		itemRaw := strings.TrimSpace(raw[i : i+end])
		v, n, err := decodeValue(itemRaw)
		if err != nil {
			return nil, 0, err
		}
		if n != len(itemRaw) {
			return nil, 0, fmt.Errorf("text before array separator")
		}
		i += end
		items = append(items, v)
		for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t' || raw[i] == '\n') {
			i++
		}
		if i < len(raw) && raw[i] == ',' {
			i++
			continue
		}
		if i < len(raw) && raw[i] == ']' {
			return items, i + 1, nil
		}
		return nil, 0, fmt.Errorf("expected comma or array end")
	}
}

func elementEnd(s string) int {
	depth := 0
	quote := byte(0)
	triple := false
	escaped := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if quote == '"' && escaped {
				escaped = false
				continue
			}
			if quote == '"' && c == '\\' {
				escaped = true
				continue
			}
			if triple && i+2 < len(s) && s[i] == quote && s[i+1] == quote && s[i+2] == quote {
				i += 2
				quote = 0
				triple = false
				continue
			}
			if !triple && c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			triple = i+2 < len(s) && s[i+1] == c && s[i+2] == c
			if triple {
				i += 2
			}
			continue
		}
		if c == '[' || c == '{' {
			depth++
			continue
		}
		if c == ']' || c == '}' {
			if depth == 0 {
				return i
			}
			depth--
			continue
		}
		if c == ',' && depth == 0 {
			return i
		}
	}
	return len(s)
}

func decodeInlineTable(raw string) (map[string]any, int, error) {
	if len(raw) < 2 || raw[0] != '{' {
		return nil, 0, fmt.Errorf("invalid inline table")
	}
	if inlineHasSeparatorNewline(raw) {
		return nil, 0, fmt.Errorf("newline in inline table")
	}
	result := map[string]any{}
	closed := map[uintptr]bool{}
	i := 1
	for {
		for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t') {
			i++
		}
		if i < len(raw) && raw[i] == '\n' {
			return nil, 0, fmt.Errorf("newline in inline table")
		}
		if i < len(raw) && raw[i] == '}' {
			return result, i + 1, nil
		}
		start := i
		eq := inlineEquals(raw[i:])
		if eq < 0 {
			return nil, 0, fmt.Errorf("invalid inline table key")
		}
		i += eq
		parts, err := parseKey(strings.TrimSpace(raw[start:i]))
		if err != nil {
			return nil, 0, err
		}
		i++
		for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t') {
			i++
		}
		if i < len(raw) && raw[i] == '\n' {
			return nil, 0, fmt.Errorf("newline in inline table")
		}
		end := elementEnd(raw[i:])
		if end == 0 {
			return nil, 0, fmt.Errorf("missing inline table value")
		}
		itemRaw := strings.TrimSpace(raw[i : i+end])
		v, n, err := decodeValue(itemRaw)
		if err != nil {
			return nil, 0, err
		}
		if n != len(itemRaw) {
			return nil, 0, fmt.Errorf("text after inline table value")
		}
		i += end
		if err := putInlineValue(result, parts, v, closed); err != nil {
			return nil, 0, err
		}
		markInlineClosed(v, closed)
		for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t') {
			i++
		}
		if i < len(raw) && raw[i] == '\n' {
			return nil, 0, fmt.Errorf("newline in inline table")
		}
		if i < len(raw) && raw[i] == ',' {
			i++
			for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t') {
				i++
			}
			if i < len(raw) && raw[i] == '}' {
				return nil, 0, fmt.Errorf("trailing comma in inline table")
			}
			continue
		}
		if i < len(raw) && raw[i] == '}' {
			return result, i + 1, nil
		}
		return nil, 0, fmt.Errorf("expected comma or inline table end")
	}
}

func inlineHasSeparatorNewline(s string) bool {
	depth := 0
	var quote byte
	triple := false
	for i := 1; i < len(s); i++ {
		if quote != 0 {
			if triple && i+2 < len(s) && s[i] == quote && s[i+1] == quote && s[i+2] == quote {
				i += 2
				quote = 0
				triple = false
			}
			continue
		}
		if s[i] == '"' || s[i] == '\'' {
			quote = s[i]
			triple = i+2 < len(s) && s[i+1] == s[i] && s[i+2] == s[i]
			if triple {
				i += 2
			}
			continue
		}
		switch s[i] {
		case '[', '{':
			depth++
		case ']', '}':
			if depth > 0 {
				depth--
			}
		case '\n':
			if depth == 0 {
				return true
			}
		}
	}
	return false
}

func inlineEquals(s string) int {
	var quote byte
	escaped := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if quote == '"' && escaped {
				escaped = false
				continue
			}
			if quote == '"' && c == '\\' {
				escaped = true
				continue
			}
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
		} else if c == '=' {
			return i
		}
	}
	return -1
}

func markInlineClosed(value any, closed map[uintptr]bool) {
	switch v := value.(type) {
	case map[string]any:
		closed[mapID(v)] = true
		for _, child := range v {
			markInlineClosed(child, closed)
		}
	case []any:
		for _, child := range v {
			markInlineClosed(child, closed)
		}
	}
}

func putInlineValue(result map[string]any, parts []string, value any, closed map[uintptr]bool) error {
	cur := result
	for _, part := range parts[:len(parts)-1] {
		existing, ok := cur[part]
		if !ok {
			nested := map[string]any{}
			cur[part] = nested
			cur = nested
			continue
		}
		nested, ok := existing.(map[string]any)
		if !ok {
			return fmt.Errorf("duplicate inline key")
		}
		if closed[mapID(nested)] {
			return fmt.Errorf("inline table is closed")
		}
		cur = nested
	}
	leaf := parts[len(parts)-1]
	if _, exists := cur[leaf]; exists {
		return fmt.Errorf("duplicate inline key")
	}
	cur[leaf] = value
	return nil
}

func decodeString(raw string) (Value, int, error) {
	if strings.HasPrefix(raw, "\"\"\"") {
		return decodeQuoted(raw, '"', true)
	}
	if strings.HasPrefix(raw, "'''") {
		return decodeQuoted(raw, '\'', true)
	}
	if strings.HasPrefix(raw, "\"") {
		return decodeQuoted(raw, '"', false)
	}
	if strings.HasPrefix(raw, "'") {
		return decodeQuoted(raw, '\'', false)
	}
	return Value{}, 0, fmt.Errorf("invalid string")
}

func decodeQuoted(raw string, quote byte, multiline bool) (Value, int, error) {
	start := 1
	if multiline {
		start = 3
		if len(raw) > 3 && raw[3] == '\n' {
			start++
		}
	}
	var out strings.Builder
	for i := start; i < len(raw); {
		if multiline && raw[i] == quote {
			run := 0
			for i+run < len(raw) && raw[i+run] == quote {
				run++
			}
			if run >= 3 {
				if run > 5 {
					return Value{}, 0, fmt.Errorf("too many delimiter quotes")
				}
				out.WriteString(strings.Repeat(string(quote), run-3))
				return Value{Type: "string", Value: out.String()}, i + run, nil
			}
		}
		if !multiline && raw[i] == quote {
			return Value{Type: "string", Value: out.String()}, i + 1, nil
		}
		if raw[i] == quote && !multiline {
			return Value{}, 0, fmt.Errorf("unexpected quote in string")
		}
		if raw[i] == '\n' && !multiline {
			return Value{}, 0, fmt.Errorf("newline in basic string")
		}
		if raw[i] == '\\' && quote == '"' {
			if i+1 >= len(raw) {
				return Value{}, 0, fmt.Errorf("unterminated escape")
			}
			j := i + 1
			for j < len(raw) && (raw[j] == ' ' || raw[j] == '\t') {
				j++
			}
			if j < len(raw) && raw[j] == '\n' {
				i = j + 1
				for i < len(raw) && (raw[i] == ' ' || raw[i] == '\t' || raw[i] == '\n') {
					i++
				}
				continue
			}
			decoded, n, err := escape(raw[i:])
			if err != nil {
				return Value{}, 0, err
			}
			out.WriteString(decoded)
			i += n
			continue
		}
		if !validStringCharacter(raw[i], multiline) {
			return Value{}, 0, fmt.Errorf("invalid control character in string")
		}
		if raw[i] < utf8.RuneSelf {
			out.WriteByte(raw[i])
			i++
		} else {
			r, n := utf8.DecodeRuneInString(raw[i:])
			if r == utf8.RuneError && n == 1 {
				return Value{}, 0, fmt.Errorf("invalid UTF-8 in string")
			}
			out.WriteString(raw[i : i+n])
			i += n
		}
	}
	return Value{}, 0, fmt.Errorf("unterminated string")
}

func hasLineContinuation(raw string) bool {
	raw = strings.TrimRight(raw, " \t")
	if len(raw) == 0 || raw[len(raw)-1] != '\\' {
		return false
	}
	n := 0
	for i := len(raw) - 1; i >= 0 && raw[i] == '\\'; i-- {
		n++
	}
	return n%2 == 1
}

func validStringCharacter(c byte, multiline bool) bool {
	if c == '\t' || c >= 0x20 {
		return c != 0x7f || multiline
	}
	return multiline && c == '\n'
}

func escape(raw string) (string, int, error) {
	switch raw[1] {
	case 'b':
		return "\b", 2, nil
	case 't':
		return "\t", 2, nil
	case 'n':
		return "\n", 2, nil
	case 'f':
		return "\f", 2, nil
	case 'r':
		return "\r", 2, nil
	case '"':
		return "\"", 2, nil
	case '\\':
		return "\\", 2, nil
	case 'u':
		return unicodeEscape(raw, 4)
	case 'U':
		return unicodeEscape(raw, 8)
	default:
		return "", 0, fmt.Errorf("invalid escape sequence")
	}
}

func unicodeEscape(raw string, digits int) (string, int, error) {
	if len(raw) < digits+2 {
		return "", 0, fmt.Errorf("incomplete Unicode escape")
	}
	code, err := strconv.ParseUint(raw[2:digits+2], 16, 32)
	if err != nil || code > utf8.MaxRune || (code >= 0xd800 && code <= 0xdfff) {
		return "", 0, fmt.Errorf("invalid Unicode scalar escape")
	}
	return string(rune(code)), digits + 2, nil
}

func stripComment(raw string) string {
	quote := byte(0)
	escaped := false
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\n' {
			quote = 0
			escaped = false
			continue
		}
		if quote == '"' && escaped {
			escaped = false
			continue
		}
		if quote == '"' && raw[i] == '\\' {
			escaped = true
			continue
		}
		if raw[i] == '"' || raw[i] == '\'' {
			if quote == 0 {
				quote = raw[i]
			} else if quote == raw[i] {
				quote = 0
			}
		} else if raw[i] == '#' && quote == 0 {
			j := i
			for j < len(raw) && raw[j] != '\n' {
				j++
			}
			raw = raw[:i] + raw[j:]
			i--
		}
	}
	return raw
}

func decodeScalar(raw string) (Value, int, error) {
	switch raw {
	case "true", "false":
		return Value{Type: "bool", Value: raw}, len(raw), nil
	case "inf", "+inf", "-inf", "nan", "+nan", "-nan":
		return Value{Type: "float", Value: raw}, len(raw), nil
	}
	if isDateTimeToken(raw) {
		value, err := parseDateTime(raw)
		if err != nil {
			return Value{}, 0, err
		}
		return value, len(raw), nil
	}
	if isIntegerToken(raw) {
		value, err := normalizeInteger(raw)
		if err != nil {
			return Value{}, 0, err
		}
		return Value{Type: "integer", Value: value}, len(raw), nil
	}
	if isFloatToken(raw) {
		return Value{Type: "float", Value: raw}, len(raw), nil
	}
	return Value{}, 0, fmt.Errorf("invalid scalar")
}

func isDateTimeToken(s string) bool {
	if len(s) == 10 && s[4] == '-' && s[7] == '-' {
		return allDigits(s[:4]) && allDigits(s[5:7]) && allDigits(s[8:])
	}
	if len(s) >= 8 && s[2] == ':' && s[5] == ':' {
		return allDigits(s[:2]) && allDigits(s[3:5]) && allDigits(s[6:8])
	}
	return len(s) >= 19 && s[4] == '-' && s[7] == '-' && (s[10] == 'T' || s[10] == 't' || s[10] == ' ')
}

func parseDateTime(s string) (Value, error) {
	if len(s) == 10 {
		year, month, day, err := dateParts(s)
		if err != nil {
			return Value{}, err
		}
		if !validDate(year, month, day) {
			return Value{}, fmt.Errorf("invalid date")
		}
		return Value{Type: "date-local", Value: s}, nil
	}
	if len(s) >= 8 && s[2] == ':' && s[5] == ':' {
		hour, minute, second, fraction, err := timeParts(s)
		if err != nil {
			return Value{}, err
		}
		if !validClock(hour, minute, second) {
			return Value{}, fmt.Errorf("invalid time")
		}
		return Value{Type: "time-local", Value: formatLocalTime(hour, minute, second, fraction)}, nil
	}

	year, month, day, err := dateParts(s[:10])
	if err != nil || len(s) < 19 || (s[10] != 'T' && s[10] != 't' && s[10] != ' ') {
		return Value{}, fmt.Errorf("invalid datetime")
	}
	timeEnd := len(s)
	offsetStart := -1
	for i := 19; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' || s[i] == 'Z' || s[i] == 'z' {
			timeEnd = i
			offsetStart = i
			break
		}
	}
	hour, minute, second, fraction, err := timeParts(s[11:timeEnd])
	if err != nil || !validDate(year, month, day) || !validClock(hour, minute, second) {
		return Value{}, fmt.Errorf("invalid datetime")
	}
	base := time.Date(year, time.Month(month), day, hour, minute, second, fraction, time.UTC)
	rest := ""
	if offsetStart >= 0 {
		rest = s[offsetStart:]
	}
	if rest == "" {
		return Value{Type: "datetime-local", Value: formatLocalDateTime(base)}, nil
	}
	if rest == "Z" || rest == "z" {
		return Value{Type: "datetime", Value: base.Format(time.RFC3339Nano)}, nil
	}
	if len(rest) != 6 || (rest[0] != '+' && rest[0] != '-') || rest[3] != ':' || !allDigits(rest[1:3]) || !allDigits(rest[4:]) {
		return Value{}, fmt.Errorf("invalid datetime offset")
	}
	offsetHour, _ := strconv.Atoi(rest[1:3])
	offsetMinute, _ := strconv.Atoi(rest[4:])
	if offsetHour > 23 || offsetMinute > 59 {
		return Value{}, fmt.Errorf("invalid datetime offset")
	}
	offset := (offsetHour*60 + offsetMinute) * 60
	if rest[0] == '-' {
		offset = -offset
	}
	zone := time.FixedZone("", offset)
	local := time.Date(year, time.Month(month), day, hour, minute, second, fraction, zone)
	return Value{Type: "datetime", Value: local.Format(time.RFC3339Nano)}, nil
}

func dateParts(s string) (int, int, int, error) {
	if len(s) != 10 || s[4] != '-' || s[7] != '-' || !allDigits(s[:4]) || !allDigits(s[5:7]) || !allDigits(s[8:]) {
		return 0, 0, 0, fmt.Errorf("invalid date")
	}
	year, _ := strconv.Atoi(s[:4])
	month, _ := strconv.Atoi(s[5:7])
	day, _ := strconv.Atoi(s[8:])
	return year, month, day, nil
}

func timeParts(s string) (int, int, int, int, error) {
	if len(s) < 8 || s[2] != ':' || s[5] != ':' || !allDigits(s[:2]) || !allDigits(s[3:5]) || !allDigits(s[6:8]) {
		return 0, 0, 0, 0, fmt.Errorf("invalid time")
	}
	hour, _ := strconv.Atoi(s[:2])
	minute, _ := strconv.Atoi(s[3:5])
	second, _ := strconv.Atoi(s[6:8])
	fraction := 0
	if len(s) > 8 {
		if s[8] != '.' || len(s) == 9 || !allDigits(s[9:]) {
			return 0, 0, 0, 0, fmt.Errorf("invalid fractional seconds")
		}
		frac := s[9:]
		if len(frac) > 9 {
			frac = frac[:9]
		}
		fraction, _ = strconv.Atoi(frac + strings.Repeat("0", 9-len(frac)))
	}
	return hour, minute, second, fraction, nil
}

func validDate(year, month, day int) bool {
	if year < 0 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}
	got := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return got.Year() == year && int(got.Month()) == month && got.Day() == day
}

func validClock(hour, minute, second int) bool {
	return hour >= 0 && hour <= 23 && minute >= 0 && minute <= 59 && second >= 0 && second <= 59
}

func formatLocalDateTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999999999")
}

func formatLocalTime(hour, minute, second, fraction int) string {
	return formatLocalDateTime(time.Date(0, 1, 1, hour, minute, second, fraction, time.UTC))[11:]
}

func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func delimitersOpen(s string) bool {
	depth := 0
	quote := byte(0)
	escaped := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if quote == '"' && escaped {
				escaped = false
				continue
			}
			if quote == '"' && c == '\\' {
				escaped = true
				continue
			}
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '#' {
			for i < len(s) && s[i] != '\n' {
				i++
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
		} else if c == '[' || c == '{' {
			depth++
		} else if c == ']' || c == '}' {
			depth--
		}
	}
	return depth > 0 || quote != 0
}

func headerName(s string, array bool) ([]string, error) {
	close := "]"
	start := 1
	if array {
		close = "]]"
		start = 2
	}
	end := strings.Index(s[start:], close)
	if end < 0 {
		return nil, fmt.Errorf("unterminated table header")
	}
	end += start
	if strings.TrimSpace(s[end+len(close):]) != "" && !strings.HasPrefix(strings.TrimSpace(s[end+len(close):]), "#") {
		return nil, fmt.Errorf("invalid table header")
	}
	return parseKey(strings.TrimSpace(s[start:end]))
}

func parseKey(s string) ([]string, error) {
	parts := []string{}
	for {
		s = strings.TrimLeft(s, " \t")
		if s == "" {
			break
		}
		if s[0] == '"' || s[0] == '\'' {
			if strings.HasPrefix(s, "\"\"\"") || strings.HasPrefix(s, "'''") {
				return nil, fmt.Errorf("multiline string is not a key")
			}
			v, n, err := decodeString(s)
			if err != nil {
				return nil, err
			}
			parts = append(parts, v.Value)
			s = s[n:]
		} else {
			n := strings.IndexByte(s, '.')
			if n < 0 {
				n = len(s)
			}
			p := strings.TrimSpace(s[:n])
			if p == "" {
				return nil, fmt.Errorf("invalid key")
			}
			for _, c := range p {
				if !(c == '_' || c == '-' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
					return nil, fmt.Errorf("invalid key")
				}
			}
			parts = append(parts, p)
			s = s[n:]
		}
		s = strings.TrimSpace(s)
		if s == "" {
			break
		}
		if s[0] != '.' {
			return nil, fmt.Errorf("invalid key")
		}
		s = s[1:]
		if s == "" {
			return nil, fmt.Errorf("invalid key")
		}
	}
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid key")
	}
	return parts, nil
}

func assignmentIndex(line string) int {
	var quote byte
	escaped := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		if quote != 0 {
			if quote == '"' && escaped {
				escaped = false
				continue
			}
			if quote == '"' && c == '\\' {
				escaped = true
				continue
			}
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '=' {
			return i
		}
	}
	return -1
}

func tableAt(root map[string]any, path []string) (map[string]any, error) {
	cur := root
	for _, p := range path {
		if v, ok := cur[p]; ok {
			m, ok := v.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("table redefinition")
			}
			cur = m
		} else {
			m := map[string]any{}
			cur[p] = m
			cur = m
		}
	}
	return cur, nil
}

func tableParent(root map[string]any, path []string) (map[string]any, string, error) {
	if len(path) == 0 {
		return nil, "", fmt.Errorf("invalid table path")
	}
	cur, err := tableAt(root, path[:len(path)-1])
	return cur, path[len(path)-1], err
}

func isIntegerToken(raw string) bool {
	if raw == "" {
		return false
	}
	start := 0
	if raw[0] == '+' || raw[0] == '-' {
		start = 1
	}
	if start == len(raw) {
		return false
	}
	if start+2 <= len(raw) && raw[start] == '0' {
		switch raw[start+1] {
		case 'x', 'o', 'b':
			return validDigits(raw[start+2:], baseForPrefix(raw[start+1]))
		}
	}
	return validDecimalDigits(raw[start:])
}

func baseForPrefix(prefix byte) int {
	switch prefix {
	case 'x':
		return 16
	case 'o':
		return 8
	default:
		return 2
	}
}

func validateInteger(raw string) error {
	_, err := normalizeInteger(raw)
	return err
}

func normalizeInteger(raw string) (string, error) {
	cleaned := strings.ReplaceAll(raw, "_", "")
	start := 0
	if cleaned[0] == '+' || cleaned[0] == '-' {
		start++
	}
	base := 10
	digits := cleaned[start:]
	parseValue := cleaned
	if len(digits) > 2 && digits[0] == '0' {
		if start > 0 {
			return "", fmt.Errorf("invalid integer")
		}
		base = baseForPrefix(digits[1])
		parseValue = digits[2:]
	}
	n, err := strconv.ParseInt(parseValue, base, 64)
	if err != nil {
		return "", fmt.Errorf("invalid integer")
	}
	return strconv.FormatInt(n, 10), nil
}

func validDecimalDigits(s string) bool {
	if len(s) == 0 || s[0] == '_' || s[len(s)-1] == '_' {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	return validDigits(s, 10)
}

func validDigits(s string, base int) bool {
	if len(s) == 0 || s[0] == '_' || s[len(s)-1] == '_' {
		return false
	}
	previousUnderscore := false
	for i := 0; i < len(s); i++ {
		if s[i] == '_' {
			if previousUnderscore || i == 0 || i+1 == len(s) {
				return false
			}
			previousUnderscore = true
			continue
		}
		value := digitValue(s[i])
		if value < 0 || value >= base {
			return false
		}
		previousUnderscore = false
	}
	return true
}

func digitValue(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'a' && c <= 'f':
		return int(c-'a') + 10
	case c >= 'A' && c <= 'F':
		return int(c-'A') + 10
	default:
		return -1
	}
}

func isFloatToken(raw string) bool {
	if raw == "" {
		return false
	}
	start := 0
	if raw[0] == '+' || raw[0] == '-' {
		start++
	}
	if start == len(raw) || strings.ContainsAny(raw[start:], "xXoObB") {
		return false
	}
	body := raw[start:]
	dot := strings.IndexByte(body, '.')
	exp := strings.IndexAny(body, "eE")
	if dot < 0 && exp < 0 {
		return false
	}
	if dot >= 0 && exp >= 0 && exp < dot {
		return false
	}
	mantissa := body
	exponent := ""
	if exp >= 0 {
		mantissa, exponent = body[:exp], body[exp+1:]
		if exponent == "" {
			return false
		}
		if exponent[0] == '+' || exponent[0] == '-' {
			exponent = exponent[1:]
		}
		if !validDecimalDigitsNoLeadingRule(exponent) {
			return false
		}
	}
	if dot >= 0 {
		left, right := mantissa[:dot], mantissa[dot+1:]
		if !validDecimalDigits(left) || !validDigits(right, 10) {
			return false
		}
		return true
	}
	return validDecimalDigitsNoLeadingRule(mantissa)
}

func validDecimalDigitsNoLeadingRule(s string) bool {
	return validDigits(s, 10)
}
