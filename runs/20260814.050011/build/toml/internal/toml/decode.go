package toml

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Decode parses a TOML 1.0 document into tables, arrays, and tagged scalar values.
func Decode(document string) (Object, error) {
	if !utf8.ValidString(document) {
		return nil, fmt.Errorf("invalid UTF-8")
	}
	if strings.Contains(strings.ReplaceAll(document, "\r\n", ""), "\r") {
		return nil, fmt.Errorf("invalid line ending")
	}
	document = strings.ReplaceAll(document, "\r\n", "\n")
	if err := validateDocumentControls(document); err != nil {
		return nil, err
	}
	root := Object{}
	current := root
	explicit := map[string]bool{}
	arrays := map[string]bool{}
	defined := map[string]bool{}
	implicit := map[string]bool{}
	dotted := map[string]bool{}
	arrayTables := map[string]bool{}
	closed := map[string]bool{}
	currentPath := []string{}
	lines, err := statements(document)
	if err != nil {
		return nil, err
	}
	for n, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if strings.Contains(line, "\n") && strings.HasPrefix(line, "[") {
			return nil, lineErr(n, fmt.Errorf("newline in table header"))
		}
		if strings.HasPrefix(line, "[[") && strings.HasSuffix(line, "]]") {
			path, err := parseKey(line[2 : len(line)-2])
			if err != nil {
				return nil, lineErr(n, err)
			}
			name := stateName(path)
			if closed[name] || explicit[name] || defined[name] {
				return nil, lineErr(n, fmt.Errorf("table %q is already defined or closed", name))
			}
			for i := 1; i < len(path); i++ {
				prefix := stateName(path[:i])
				if closed[prefix] || (defined[prefix] && !arrays[prefix]) || (arrays[prefix] && !arrayTables[prefix]) {
					return nil, lineErr(n, fmt.Errorf("table is inside a closed or defined value"))
				}
			}
			parent, leaf, err := tableParent(root, path, true, implicit, arrays)
			if err != nil {
				return nil, lineErr(n, err)
			}
			if existing, ok := parent[leaf]; ok {
				if _, isArray := existing.([]any); !isArray {
					return nil, lineErr(n, fmt.Errorf("table conflicts with existing value"))
				}
			}
			arr, _ := parent[leaf].([]any)
			obj := Object{}
			arr = append(arr, obj)
			parent[leaf] = arr
			current = obj
			currentPath = path
			arrays[name] = true
			arrayTables[name] = true
			clearPathState(defined, implicit, closed, dotted, name)
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			path, err := parseKey(line[1 : len(line)-1])
			if err != nil {
				return nil, lineErr(n, err)
			}
			name := stateName(path)
			if closed[name] || defined[name] || arrays[name] {
				return nil, lineErr(n, fmt.Errorf("table %q is already defined or closed", name))
			}
			for i := 1; i < len(path); i++ {
				prefix := stateName(path[:i])
				if closed[prefix] || (defined[prefix] && !arrays[prefix]) || (arrays[prefix] && !arrayTables[prefix]) {
					return nil, lineErr(n, fmt.Errorf("table is inside a closed inline table"))
				}
			}
			obj, err := ensureTable(root, path, explicit, implicit, arrays, dotted)
			if err != nil {
				return nil, lineErr(n, err)
			}
			current = obj
			currentPath = path
			continue
		}
		pos, err := assignmentEquals(line)
		if err != nil {
			return nil, lineErr(n, err)
		}
		key, err := parseKey(line[:pos])
		if err != nil {
			return nil, lineErr(n, err)
		}
		if len(key) == 0 {
			return nil, lineErr(n, fmt.Errorf("empty key"))
		}
		value, err := parseValue(strings.TrimSpace(line[pos+1:]))
		if err != nil {
			return nil, lineErr(n, err)
		}
		fullKey := append(append([]string{}, currentPath...), key...)
		fullName := stateName(fullKey)
		if defined[fullName] || closed[fullName] {
			return nil, lineErr(n, fmt.Errorf("duplicate or closed key %q", fullName))
		}
		for i := 1; i < len(fullKey); i++ {
			prefix := stateName(fullKey[:i])
			if closed[prefix] || (i > len(currentPath) && explicit[prefix]) {
				return nil, lineErr(n, fmt.Errorf("key %q is in a closed inline table", fullName))
			}
		}
		target := current
		for i, part := range key[:len(key)-1] {
			next, ok := target[part].(Object)
			if !ok {
				if target[part] != nil {
					return nil, lineErr(n, fmt.Errorf("key %q is not a table", part))
				}
				next = Object{}
				target[part] = next
				prefix := append([]string{}, currentPath...)
				prefix = append(prefix, key[:i+1]...)
				implicit[stateName(prefix)] = true
				dotted[stateName(prefix)] = true
			}
			target = next
		}
		leaf := key[len(key)-1]
		if _, exists := target[leaf]; exists {
			return nil, lineErr(n, fmt.Errorf("duplicate key %q", leaf))
		}
		target[leaf] = value
		defined[fullName] = true
		if obj, ok := value.(Object); ok {
			markClosed(closed, fullKey, obj)
		}
		if _, ok := value.([]any); ok {
			arrays[fullName] = true
		}
	}
	return root, nil
}

func validateDocumentControls(s string) error {
	for _, r := range s {
		if r == '\t' || r == '\n' || r == '\r' {
			continue
		}
		if unicode.IsSpace(r) && r != ' ' {
			return fmt.Errorf("invalid whitespace character")
		}
		if r < 0x20 || r == 0x7f {
			return fmt.Errorf("prohibited control character")
		}
	}
	return nil
}

func lineErr(n int, err error) error { return fmt.Errorf("line %d: %w", n+1, err) }

func statements(s string) ([]string, error) {
	var out []string
	start, depth := 0, 0
	quote := byte(0)
	triple := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote == 0 && c == '#' {
			for i < len(s) && s[i] != '\n' {
				i++
			}
			if i == len(s) {
				break
			}
			c = s[i]
		}
		if quote != 0 {
			if triple {
				if c == quote && i+2 < len(s) && s[i+1] == quote && s[i+2] == quote && (quote == '\'' || !escapedAt(s, i)) {
					run := 3
					for i+run < len(s) && s[i+run] == quote {
						run++
					}
					i += run - 1
					quote, triple = 0, false
				}
			} else if c == quote && (quote == '\'' || !escapedAt(s, i)) {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			if i+2 < len(s) && s[i+1] == c && s[i+2] == c {
				triple = true
				i += 2
			}
			continue
		}
		if c == '[' || c == '{' {
			depth++
		}
		if c == ']' || c == '}' {
			depth--
			if depth < 0 {
				return nil, fmt.Errorf("unexpected closing delimiter")
			}
		}
		if c == '\n' && depth == 0 {
			out = append(out, stripComment(s[start:i]))
			start = i + 1
		}
	}
	if quote != 0 || depth != 0 {
		return nil, fmt.Errorf("unterminated value")
	}
	out = append(out, stripComment(s[start:]))
	return out, nil
}

func escapedAt(s string, i int) bool {
	count := 0
	for i > 0 && s[i-1] == '\\' {
		count++
		i--
	}
	return count%2 == 1
}

func stripComment(s string) string {
	quote := byte(0)
	triple := false
	escaped := false
	var out strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if triple {
				if c == quote && i+2 < len(s) && s[i+1] == quote && s[i+2] == quote && (quote == '\'' || !escapedAt(s, i)) {
					out.WriteString(s[i : i+3])
					i += 2
					quote, triple = 0, false
					continue
				}
				out.WriteByte(c)
				continue
			}
			if quote == '"' && escaped {
				escaped = false
				out.WriteByte(c)
				continue
			}
			if quote == '"' && c == '\\' {
				escaped = true
				out.WriteByte(c)
				continue
			}
			if c == quote {
				quote = 0
			}
			out.WriteByte(c)
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			if i+2 < len(s) && s[i+1] == c && s[i+2] == c {
				triple = true
				out.WriteString(s[i : i+3])
				i += 2
				continue
			}
		}
		if c == '#' {
			for i < len(s) && s[i] != '\n' {
				i++
			}
			if i < len(s) {
				out.WriteByte('\n')
			}
			continue
		}
		out.WriteByte(c)
	}
	return out.String()
}

func assignmentEquals(s string) (int, error) {
	quote := byte(0)
	depth := 0
	for i, c := range []byte(s) {
		if quote != 0 {
			if c == quote && (i == 0 || s[i-1] != '\\') {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
		}
		if c == '[' || c == '{' {
			depth++
		}
		if c == ']' || c == '}' {
			depth--
		}
		if c == '=' && depth == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("expected key/value pair")
}

func parseKey(s string) ([]string, error) {
	var keys []string
	for {
		s = strings.TrimSpace(s)
		if s == "" {
			return nil, fmt.Errorf("empty key")
		}
		var key string
		if s[0] == '"' || s[0] == '\'' {
			q := s[0]
			end := 1
			for end < len(s) {
				if s[end] == q && (q == '\'' || !escapedAt(s, end)) {
					break
				}
				end++
			}
			if end == len(s) {
				return nil, fmt.Errorf("unterminated key")
			}
			raw := s[:end+1]
			if q == '"' {
				var err error
				raw, err = unquote(raw)
				if err != nil {
					return nil, err
				}
			} else {
				raw = raw[1:end]
				if strings.ContainsAny(raw, "\r\n") {
					return nil, fmt.Errorf("newline in key")
				}
			}
			key, s = raw, s[end+1:]
		} else {
			end := strings.IndexByte(s, '.')
			if end < 0 {
				end = len(s)
			}
			key, s = strings.TrimSpace(s[:end]), s[end:]
			if key == "" {
				return nil, fmt.Errorf("invalid key")
			}
			for _, r := range key {
				if !(r == '_' || r == '-' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9') {
					return nil, fmt.Errorf("invalid bare key")
				}
			}
		}
		keys = append(keys, key)
		s = strings.TrimSpace(s)
		if s == "" {
			return keys, nil
		}
		if s[0] != '.' {
			return nil, fmt.Errorf("invalid dotted key")
		}
		s = s[1:]
	}
}

func tableParent(root Object, path []string, forArray bool, implicit, arrays map[string]bool) (Object, string, error) {
	if len(path) == 0 {
		return nil, "", fmt.Errorf("empty table")
	}
	obj := root
	for i, part := range path[:len(path)-1] {
		name := stateName(path[:i+1])
		if arr, ok := obj[part].([]any); ok {
			if len(arr) == 0 {
				return nil, "", fmt.Errorf("empty table array")
			}
			next, ok := arr[len(arr)-1].(Object)
			if !ok {
				return nil, "", fmt.Errorf("table conflict")
			}
			obj = next
			continue
		}
		if next, ok := obj[part].(Object); ok {
			obj = next
			continue
		}
		if obj[part] != nil {
			return nil, "", fmt.Errorf("table conflict")
		}
		next := Object{}
		obj[part] = next
		implicit[name] = true
		obj = next
	}
	return obj, path[len(path)-1], nil
}

func ensureTable(root Object, path []string, explicit, implicit, arrays, dotted map[string]bool) (Object, error) {
	obj := root
	for i, p := range path {
		if x, ok := obj[p].(Object); ok {
			obj = x
			continue
		}
		if arr, ok := obj[p].([]any); ok && len(arr) > 0 {
			if x, ok := arr[len(arr)-1].(Object); ok {
				obj = x
				continue
			}
		}
		if obj[p] != nil {
			return nil, fmt.Errorf("table conflict")
		}
		x := Object{}
		obj[p] = x
		obj = x
		implicit[stateName(path[:i+1])] = true
	}
	name := stateName(path)
	if dotted[name] {
		return nil, fmt.Errorf("table was already created by a dotted key")
	}
	if explicit[name] && !hasArrayAncestor(root, path) {
		return nil, fmt.Errorf("table redefinition")
	}
	explicit[name] = true
	delete(implicit, name)
	return obj, nil
}

func hasArrayAncestor(root Object, path []string) bool {
	obj := root
	for _, part := range path[:len(path)-1] {
		if arr, ok := obj[part].([]any); ok {
			return len(arr) > 0
		}
		x, ok := obj[part].(Object)
		if !ok {
			return false
		}
		obj = x
	}
	return false
}

func parseValue(s string) (any, error) {
	s = strings.TrimSpace(s)
	if s == "true" || s == "false" {
		return Value{Bool, s}, nil
	}
	if len(s) >= 2 && (s[0] == '"' || s[0] == '\'') {
		return parseString(s)
	}
	if strings.HasPrefix(s, "[") {
		return parseArray(s)
	}
	if strings.HasPrefix(s, "{") {
		return parseInline(s)
	}
	if kind, ok := parseDatetime(s); ok {
		return Value{kind, s}, nil
	}
	if kind, lexeme, ok := parseNumber(s); ok {
		return Value{kind, lexeme}, nil
	}
	return nil, fmt.Errorf("invalid value %q", s)
}

var (
	datePattern          = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	timePattern          = regexp.MustCompile(`^(\d{2}):(\d{2}):(\d{2})(\.\d+)?$`)
	datetimePattern      = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})[Tt ](\d{2}:\d{2}:\d{2}(?:\.\d+)?)(Z|z|[+-]\d{2}:\d{2})$`)
	localDatetimePattern = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})[Tt ](\d{2}:\d{2}:\d{2}(?:\.\d+)?)$`)
	floatPattern         = regexp.MustCompile(`^[+-]?(0|[1-9](?:[0-9]|_[0-9])*)(?:\.[0-9](?:[0-9]|_[0-9])*)?(?:[eE][+-]?[0-9](?:[0-9]|_[0-9])*)?$`)
)

func parseDatetime(s string) (Kind, bool) {
	if match := datetimePattern.FindStringSubmatch(s); match != nil {
		if match[3] != "Z" && match[3] != "z" {
			hour, _ := strconv.Atoi(match[3][1:3])
			minute, _ := strconv.Atoi(match[3][4:6])
			if hour > 23 || minute > 59 {
				return "", false
			}
		}
		fraction := fractionDigits(match[2])
		if len(fraction) > 9 {
			fraction = fraction[:9]
		}
		candidate := match[1] + "T" + trimFraction(match[2], fraction) + strings.ToUpper(match[3])
		if _, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", candidate); err == nil {
			return Datetime, true
		}
	}
	if match := localDatetimePattern.FindStringSubmatch(s); match != nil {
		fraction := fractionDigits(match[2])
		if len(fraction) > 9 {
			fraction = fraction[:9]
		}
		candidate := match[1] + "T" + trimFraction(match[2], fraction)
		if _, err := time.Parse("2006-01-02T15:04:05.999999999", candidate); err == nil {
			return DatetimeLocal, true
		}
	}
	if datePattern.MatchString(s) {
		if _, err := time.Parse("2006-01-02", s); err == nil {
			return DateLocal, true
		}
	}
	if match := timePattern.FindStringSubmatch(s); match != nil {
		fraction := fractionDigits(s)
		if len(fraction) > 9 {
			fraction = fraction[:9]
		}
		candidate := match[1] + ":" + match[2] + ":" + match[3]
		if fraction != "" {
			candidate += "." + fraction
		}
		if _, err := time.Parse("15:04:05.999999999", candidate); err == nil {
			return TimeLocal, true
		}
	}
	return "", false
}

func fractionDigits(s string) string {
	if dot := strings.IndexByte(s, '.'); dot >= 0 {
		return s[dot+1:]
	}
	return ""
}

func trimFraction(s, fraction string) string {
	if dot := strings.IndexByte(s, '.'); dot >= 0 {
		return s[:dot] + "." + fraction
	}
	return s
}

func parseNumber(s string) (Kind, string, bool) {
	if s == "inf" || s == "+inf" || s == "-inf" {
		return Float, strings.TrimPrefix(s, "+"), true
	}
	if s == "nan" || s == "+nan" || s == "-nan" {
		return Float, "nan", true
	}
	if validInteger(s) {
		clean := strings.ReplaceAll(s, "_", "")
		if n, err := parseSignedInteger(clean); err == nil {
			return Integer, strconv.FormatInt(n, 10), true
		}
	}
	if validFloat(s) {
		clean := strings.ReplaceAll(s, "_", "")
		if f, err := strconv.ParseFloat(clean, 64); err == nil && !math.IsInf(f, 0) && !math.IsNaN(f) {
			return Float, s, true
		}
	}
	return "", "", false
}

func normalizeInteger(s string) string {
	negative := strings.HasPrefix(s, "-")
	s = strings.TrimLeft(s, "+-")
	base := 10
	if strings.HasPrefix(s, "0x") {
		base = 16
		s = s[2:]
	} else if strings.HasPrefix(s, "0o") {
		base = 8
		s = s[2:]
	} else if strings.HasPrefix(s, "0b") {
		base = 2
		s = s[2:]
	}
	n, _ := strconv.ParseInt(s, base, 64)
	if negative {
		n = -n
	}
	return strconv.FormatInt(n, 10)
}

func parseSignedInteger(s string) (int64, error) {
	sign := ""
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign, s = s[:1], s[1:]
	}
	base := 10
	if len(s) > 2 && s[0] == '0' {
		switch s[1] {
		case 'x':
			base, s = 16, s[2:]
		case 'o':
			base, s = 8, s[2:]
		case 'b':
			base, s = 2, s[2:]
		}
	}
	if sign == "-" {
		n, err := strconv.ParseUint(s, base, 64)
		if err != nil || n > 1<<63 {
			return 0, strconv.ErrRange
		}
		if n == 1<<63 {
			return -1 << 63, nil
		}
		return -int64(n), nil
	}
	n, err := strconv.ParseUint(s, base, 64)
	if err != nil || n > 1<<63-1 {
		return 0, strconv.ErrRange
	}
	return int64(n), nil
}

func validInteger(s string) bool {
	if (strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-")) && len(s) > 2 {
		unsigned := s[1:]
		if strings.HasPrefix(unsigned, "0x") || strings.HasPrefix(unsigned, "0o") || strings.HasPrefix(unsigned, "0b") {
			return false
		}
	}
	unsigned := strings.TrimPrefix(strings.TrimPrefix(s, "+"), "-")
	if unsigned == "" || !validUnderscores(unsigned) {
		return false
	}
	digits, base := unsigned, 10
	if len(unsigned) > 2 && unsigned[0] == '0' {
		switch unsigned[1] {
		case 'x':
			base, digits = 16, unsigned[2:]
		case 'o':
			base, digits = 8, unsigned[2:]
		case 'b':
			base, digits = 2, unsigned[2:]
		}
	}
	if base == 10 && strings.ContainsAny(digits, ".eE") {
		return false
	}
	if digits == "" || digits[0] == '_' || digits[len(digits)-1] == '_' || (base == 10 && len(digits) > 1 && digits[0] == '0') {
		return false
	}
	for _, r := range digits {
		if r == '_' {
			continue
		}
		valid := (base == 2 && (r == '0' || r == '1')) || (base == 8 && r >= '0' && r <= '7') || (base == 10 && r >= '0' && r <= '9') || (base == 16 && ((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')))
		if !valid {
			return false
		}
	}
	for i := 1; i < len(digits)-1; i++ {
		if digits[i] == '_' && !validBaseDigit(digits[i-1], base) || digits[i] == '_' && !validBaseDigit(digits[i+1], base) {
			return false
		}
	}
	return true
}

func validBaseDigit(r byte, base int) bool {
	return (base == 2 && (r == '0' || r == '1')) || (base == 8 && r >= '0' && r <= '7') || (base == 10 && r >= '0' && r <= '9') || (base == 16 && ((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')))
}

func validUnderscores(s string) bool {
	return !strings.Contains(s, "__") && !strings.HasPrefix(s, "_") && !strings.HasSuffix(s, "_")
}

func validFloat(s string) bool {
	if s == "" || !validUnderscores(s) {
		return false
	}
	unsigned := strings.TrimPrefix(strings.TrimPrefix(s, "+"), "-")
	if strings.HasPrefix(unsigned, "0x") || strings.HasPrefix(unsigned, "0o") || strings.HasPrefix(unsigned, "0b") || !strings.ContainsAny(unsigned, ".eE") {
		return false
	}
	return floatPattern.MatchString(s)
}

/*
The remaining numeric characters are checked above; this helper is retained
for compatibility with the parser's earlier lexical utilities.
*/
func validNumber(s string) bool {
	for _, r := range s {
		if r == '_' {
			continue
		}
		if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F' || strings.ContainsRune("+-._xobXOBeE", r)) {
			return false
		}
	}
	return true
}
func parseString(s string) (Value, error) {
	q := s[0]
	if len(s) >= 6 && (s[:3] == "\"\"\"" || s[:3] == "'''") {
		end := multilineEnd(s, q)
		if end < 0 || strings.TrimSpace(s[end+3:]) != "" {
			return Value{}, fmt.Errorf("unterminated multiline string")
		}
		body := s[3:end]
		body = strings.TrimPrefix(body, "\n")
		if q == '"' {
			decoded, err := decodeBasic(body, true)
			if err != nil {
				return Value{}, err
			}
			body = decoded
		} else if err := validateStringChars(body, true); err != nil {
			return Value{}, err
		}
		return Value{String, body}, nil
	}
	end := -1
	for i := 1; i < len(s); i++ {
		if s[i] == q && (q == '\'' || !escapedAt(s, i)) {
			end = i
			break
		}
	}
	if end < 0 || strings.TrimSpace(s[end+1:]) != "" {
		return Value{}, fmt.Errorf("unterminated string")
	}
	if q == '\'' {
		body := s[1:end]
		if err := validateStringChars(body, false); err != nil {
			return Value{}, err
		}
		return Value{String, body}, nil
	}
	raw, err := decodeBasic(s[1:end], false)
	return Value{String, raw}, err
}

// multilineEnd returns the start of the closing delimiter. In a longer run
// of quotes, the final three quotes close the string and preceding quotes are content.
func multilineEnd(s string, q byte) int {
	for i := 3; i < len(s); {
		if s[i] != q {
			i++
			continue
		}
		start := i
		for i < len(s) && s[i] == q {
			i++
		}
		if i-start >= 3 {
			candidate := i - 3
			if q == '\'' || !escapedAt(s, candidate) {
				if i-start < 6 || escapedAt(s, start) {
					return candidate
				}
				return -1
			}
		}
	}
	return -1
}
func unquote(s string) (string, error) {
	v, err := strconv.Unquote(s)
	if err != nil {
		return "", fmt.Errorf("invalid string: %w", err)
	}
	return v, nil
}

func decodeBasic(s string, multiline bool) (string, error) {
	var out strings.Builder
	for i := 0; i < len(s); {
		if s[i] != '\\' {
			r, size := utf8.DecodeRuneInString(s[i:])
			if r == utf8.RuneError && size == 1 {
				return "", fmt.Errorf("invalid UTF-8 in string")
			}
			if prohibitedStringRune(r, multiline) {
				return "", fmt.Errorf("prohibited control character in string")
			}
			out.WriteRune(r)
			i += size
			continue
		}
		i++
		if i == len(s) {
			return "", fmt.Errorf("invalid escape")
		}
		if multiline && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
			start := i
			for i < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
				i++
			}
			if !strings.ContainsAny(s[start:i], "\n\r") {
				return "", fmt.Errorf("invalid escape")
			}
			continue
		}
		escapes := map[byte]rune{'b': '\b', 't': '\t', 'n': '\n', 'f': '\f', 'r': '\r', '"': '"', '\\': '\\'}
		if r, ok := escapes[s[i]]; ok {
			out.WriteRune(r)
			i++
			continue
		}
		width := 0
		if s[i] == 'u' {
			width = 4
		}
		if s[i] == 'U' {
			width = 8
		}
		if width == 0 || i+1+width > len(s) {
			return "", fmt.Errorf("invalid escape")
		}
		n, err := strconv.ParseUint(s[i+1:i+1+width], 16, 32)
		if err != nil || !utf8.ValidRune(rune(n)) || n >= 0xD800 && n <= 0xDFFF {
			return "", fmt.Errorf("invalid Unicode escape")
		}
		out.WriteRune(rune(n))
		i += width + 1
	}
	return out.String(), nil
}

func prohibitedStringRune(r rune, multiline bool) bool {
	if r == '\n' || r == '\r' {
		return !multiline
	}
	return r < 0x20 && r != '\t'
}

func validateStringChars(s string, multiline bool) error {
	if !utf8.ValidString(s) {
		return fmt.Errorf("invalid UTF-8 in string")
	}
	for _, r := range s {
		if prohibitedStringRune(r, multiline) {
			return fmt.Errorf("prohibited control character in string")
		}
	}
	return nil
}
func splitItems(s string) ([]string, error) {
	var out []string
	start, depth := 0, 0
	quote := byte(0)
	triple := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if triple {
				if c == quote && i+2 < len(s) && s[i+1] == quote && s[i+2] == quote && (quote == '\'' || !escapedAt(s, i)) {
					quote, triple = 0, false
					i += 2
				}
				continue
			}
			if c == quote && (quote == '\'' || !escapedAt(s, i)) {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			if i+2 < len(s) && s[i+1] == c && s[i+2] == c {
				triple = true
				i += 2
			}
		}
		if c == '[' || c == '{' {
			depth++
		}
		if c == ']' || c == '}' {
			depth--
			if depth < 0 {
				return nil, fmt.Errorf("unexpected closing delimiter")
			}
		}
		if c == ',' && depth == 0 {
			item := strings.TrimSpace(s[start:i])
			if item == "" {
				if len(out) == 0 || i == len(s)-1 {
					return nil, fmt.Errorf("invalid array element")
				}
				return nil, fmt.Errorf("invalid array element")
			}
			out = append(out, item)
			start = i + 1
		}
	}
	if quote != 0 || depth != 0 {
		return nil, fmt.Errorf("unterminated array element")
	}
	last := strings.TrimSpace(s[start:])
	if last != "" {
		out = append(out, last)
	} else if len(out) > 0 {
		// An empty final item is the one permitted trailing comma.
		return out, nil
	}
	return out, nil
}
func parseArray(s string) ([]any, error) {
	s = stripComment(s)
	s = strings.TrimSpace(s)
	if len(s) < 2 || s[len(s)-1] != ']' {
		return nil, fmt.Errorf("unterminated array")
	}
	items, err := splitItems(s[1 : len(s)-1])
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(items))
	for _, item := range items {
		v, e := parseValue(item)
		if e != nil {
			return nil, e
		}
		out = append(out, v)
	}
	return out, nil
}
func parseInline(s string) (Object, error) {
	if inlineTableHasNewline(s) {
		return nil, fmt.Errorf("newline in inline table")
	}
	s = stripComment(s)
	if len(s) < 2 || s[len(s)-1] != '}' {
		return nil, fmt.Errorf("unterminated inline table")
	}
	o := Object{}
	closed := map[string]bool{}
	items, err := splitInlineItems(s[1 : len(s)-1])
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		p, e := assignmentEquals(item)
		if e != nil {
			return nil, e
		}
		k, e := parseKey(item[:p])
		if e != nil || len(k) == 0 {
			return nil, fmt.Errorf("invalid inline key")
		}
		target := o
		for _, part := range k[:len(k)-1] {
			if closed[stateName(k[:len(k)-1])] {
				return nil, fmt.Errorf("key is inside a closed inline table")
			}
			next, ok := target[part].(Object)
			if !ok {
				if target[part] != nil {
					return nil, fmt.Errorf("duplicate key")
				}
				next = Object{}
				target[part] = next
			}
			target = next
		}
		if _, ok := target[k[len(k)-1]]; ok {
			return nil, fmt.Errorf("duplicate key")
		}
		v, e := parseValue(strings.TrimSpace(item[p+1:]))
		if e != nil {
			return nil, e
		}
		leaf := k[len(k)-1]
		target[leaf] = v
		if child, ok := v.(Object); ok {
			markClosed(closed, k, child)
		}
	}
	return o, nil
}

func splitInlineItems(s string) ([]string, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	items, err := splitItems(s)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(strings.TrimSpace(s), ",") {
		return nil, fmt.Errorf("trailing comma in inline table")
	}
	for _, item := range items {
		if strings.TrimSpace(item) == "" {
			return nil, fmt.Errorf("empty inline table entry")
		}
	}
	return items, nil
}

func inlineTableHasNewline(s string) bool {
	quote := byte(0)
	triple := false
	arrayDepth := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if triple {
				if c == quote && i+2 < len(s) && s[i+1] == quote && s[i+2] == quote && (quote == '\'' || !escapedAt(s, i)) {
					quote, triple = 0, false
					i += 2
				}
				continue
			}
			if c == quote && (quote == '\'' || !escapedAt(s, i)) {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			if i+2 < len(s) && s[i+1] == c && s[i+2] == c {
				triple = true
				i += 2
			}
			continue
		}
		if c == '[' {
			arrayDepth++
		} else if c == ']' && arrayDepth > 0 {
			arrayDepth--
		}
		if c == '\n' || c == '\r' {
			if arrayDepth == 0 {
				return true
			}
		}
	}
	return false
}

func markClosed(closed map[string]bool, path []string, obj Object) {
	name := stateName(path)
	closed[name] = true
	for key, value := range obj {
		if child, ok := value.(Object); ok {
			markClosed(closed, append(append([]string{}, path...), key), child)
		}
	}
}

func clearPathState(defined, implicit, closed, dotted map[string]bool, prefix string) {
	for _, state := range []map[string]bool{defined, implicit, closed, dotted} {
		for key := range state {
			if key == prefix || strings.HasPrefix(key, prefix+"\x00") {
				delete(state, key)
			}
		}
	}
}

func stateName(path []string) string {
	return strings.Join(path, "\x00")
}
