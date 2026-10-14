package toml

import (
	"strings"
)

// maxScalar is the largest Unicode code point, U+10FFFF.
const maxScalar = 0x10FFFF

// simpleEscapes are the one-character escapes TOML v1.0.0 lists for basic strings.
var simpleEscapes = map[byte]string{
	'b':  "\b",
	't':  "\t",
	'n':  "\n",
	'f':  "\f",
	'r':  "\r",
	'"':  "\"",
	'\\': "\\",
}

// basicString reads a single-line basic string at p.pos, which holds the opening quote.
func (p *parser) basicString() (string, error) {
	p.pos++
	var b strings.Builder
	for {
		if p.done() || p.atNewline() {
			return "", p.errorf("unterminated basic string")
		}
		c := p.src[p.pos]
		switch {
		case c == '"':
			p.pos++
			return b.String(), nil
		case c == '\\':
			s, err := p.escape()
			if err != nil {
				return "", err
			}
			b.WriteString(s)
		case isDisallowedControl(c):
			return "", p.errorf("control character in basic string")
		default:
			b.WriteByte(c)
			p.pos++
		}
	}
}

// literalString reads a single-line literal string at p.pos, which holds the opening
// quote. Nothing inside is escaped.
func (p *parser) literalString() (string, error) {
	p.pos++
	start := p.pos
	for !p.done() && p.src[p.pos] != '\'' {
		if p.atNewline() || isDisallowedControl(p.src[p.pos]) {
			return "", p.errorf("invalid character in literal string")
		}
		p.pos++
	}
	if p.done() {
		return "", p.errorf("unterminated literal string")
	}
	s := p.src[start:p.pos]
	p.pos++
	return s, nil
}

// multiline reads a multi-line string whose delimiter is quote repeated three times.
// Escapes apply only when escapes is true, which is the basic form. A line break right
// after the opening delimiter is trimmed. Up to two extra quotes may sit directly before
// the closing delimiter.
func (p *parser) multiline(quote byte, escapes bool) (string, error) {
	p.pos += 3
	p.trimOpeningNewline()
	var b strings.Builder
	for {
		if p.done() {
			return "", p.errorf("unterminated multi-line string")
		}
		c := p.src[p.pos]
		switch {
		case c == quote:
			run := p.quoteRun(quote)
			if run < 3 {
				b.WriteString(p.src[p.pos : p.pos+run])
				p.pos += run
				continue
			}
			if run > 5 {
				return "", p.errorf("too many quotes closing multi-line string")
			}
			b.WriteString(p.src[p.pos : p.pos+run-3])
			p.pos += run
			return b.String(), nil
		case escapes && c == '\\' && p.lineEndingBackslash():
			p.skipLineEnding()
		case escapes && c == '\\':
			s, err := p.escape()
			if err != nil {
				return "", err
			}
			b.WriteString(s)
		case p.atCRLF():
			b.WriteString("\r\n")
			p.pos += 2
		case c == '\n':
			b.WriteByte(c)
			p.pos++
		case isDisallowedControl(c):
			return "", p.errorf("control character in multi-line string")
		default:
			b.WriteByte(c)
			p.pos++
		}
	}
}

// escape decodes the escape sequence whose backslash is at p.pos. Only the escapes
// TOML v1.0.0 lists are accepted; any other escape is a syntax error.
func (p *parser) escape() (string, error) {
	p.pos++
	if p.done() {
		return "", p.errorf("unterminated escape")
	}
	c := p.src[p.pos]
	if s, ok := simpleEscapes[c]; ok {
		p.pos++
		return s, nil
	}
	switch c {
	case 'u':
		return p.unicodeEscape(4)
	case 'U':
		return p.unicodeEscape(8)
	}
	return "", p.errorf("unknown escape \\%c", c)
}

// unicodeEscape decodes the digits after \u (digits is 4) or \U (digits is 8). The
// code point must be a Unicode scalar value: at most U+10FFFF and not a surrogate.
func (p *parser) unicodeEscape(digits int) (string, error) {
	start := p.pos + 1
	end := start + digits
	if end > len(p.src) {
		return "", p.errorf("unicode escape needs %d hex digits", digits)
	}
	var code uint32
	for i := start; i < end; i++ {
		v, ok := hexValue(p.src[i])
		if !ok {
			return "", p.errorf("unicode escape needs %d hex digits", digits)
		}
		code = code<<4 | v
	}
	if code > maxScalar || (code >= 0xD800 && code <= 0xDFFF) {
		return "", p.errorf("unicode escape %X is not a Unicode scalar value", code)
	}
	p.pos = end
	return string(rune(code)), nil
}

// lineEndingBackslash reports whether the backslash at p.pos ends a line. Only spaces
// and tabs may stand between it and the line break.
func (p *parser) lineEndingBackslash() bool {
	i := p.pos + 1
	for i < len(p.src) && (p.src[i] == ' ' || p.src[i] == '\t') {
		i++
	}
	return i < len(p.src) && (p.src[i] == '\n' || strings.HasPrefix(p.src[i:], "\r\n"))
}

// skipLineEnding consumes a line-ending backslash and every whitespace character
// after it, up to the next content.
func (p *parser) skipLineEnding() {
	p.pos++
	for !p.done() {
		switch {
		case p.src[p.pos] == ' ' || p.src[p.pos] == '\t' || p.src[p.pos] == '\n':
			p.pos++
		case p.atCRLF():
			p.pos += 2
		default:
			return
		}
	}
}

func (p *parser) trimOpeningNewline() {
	if p.atCRLF() {
		p.pos += 2
	} else if p.atNewline() {
		p.pos++
	}
}

// quoteRun counts the consecutive quote characters starting at p.pos.
func (p *parser) quoteRun(quote byte) int {
	n := 0
	for p.pos+n < len(p.src) && p.src[p.pos+n] == quote {
		n++
	}
	return n
}

func hexValue(c byte) (uint32, bool) {
	switch {
	case c >= '0' && c <= '9':
		return uint32(c - '0'), true
	case c >= 'a' && c <= 'f':
		return uint32(c-'a') + 10, true
	case c >= 'A' && c <= 'F':
		return uint32(c-'A') + 10, true
	}
	return 0, false
}
