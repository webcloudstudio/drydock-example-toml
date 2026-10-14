// Package toml decodes TOML v1.0.0 documents using only the standard library.
package toml

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// ErrSyntax reports a document that TOML v1.0.0 does not allow. Callers inspect
// it with errors.Is; the wrapped message names the line that failed.
var ErrSyntax = errors.New("invalid TOML")

// Decode reads a TOML document from r and returns its tagged-value model.
// Each leaf maps to {"type", "value"} with the value as a string; each table maps to
// an object, and each array of tables maps to an array of objects.
// Values are limited to strings, booleans, integers, floats, and date-times so far.
func Decode(r io.Reader) (map[string]any, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read document: %w", err)
	}
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("%w: document is not valid UTF-8", ErrSyntax)
	}
	p := &parser{src: string(data)}
	doc, err := p.document()
	if err != nil {
		return nil, err
	}
	return doc.export(), nil
}

// parser scans src by byte offset. Multi-byte runes are copied through unchanged,
// so the ASCII delimiters the scanner checks are never confused with them.
type parser struct {
	src string
	pos int
}

// document reads every line into the root table. cur is the table that key/value pairs
// land in; it changes at each header.
func (p *parser) document() (*table, error) {
	root := newTable(originHeader)
	cur := root
	for {
		p.skipSpaces()
		if p.done() {
			return root, nil
		}
		switch c := p.src[p.pos]; {
		case c == '#' || c == '\n' || c == '\r':
			if err := p.endOfLine(); err != nil {
				return nil, err
			}
			continue
		case c == '[':
			parts, array, err := p.header()
			if err != nil {
				return nil, err
			}
			if cur, err = p.enterHeader(root, parts, array); err != nil {
				return nil, err
			}
		default:
			parts, value, err := p.keyValue()
			if err != nil {
				return nil, err
			}
			if err := p.insert(cur, parts, value); err != nil {
				return nil, err
			}
		}
		if err := p.endOfLine(); err != nil {
			return nil, err
		}
	}
}

func (p *parser) keyValue() ([]string, any, error) {
	parts, err := p.key()
	if err != nil {
		return nil, nil, err
	}
	name := strings.Join(parts, ".")
	p.skipSpaces()
	if p.done() || p.src[p.pos] != '=' {
		return nil, nil, p.errorf("expected '=' after key %q", name)
	}
	p.pos++
	p.skipSpaces()
	value, err := p.value()
	if err != nil {
		return nil, nil, err
	}
	return parts, value, nil
}

// value reads one value. Strings start with a quote and arrays and inline tables with a
// bracket or brace; everything else is a bare boolean, integer, or float. The result is
// a tagged model, an arrayValue, or a closed *table. The longer string delimiters are
// tested first so that `"""` is never read as `""`.
func (p *parser) value() (any, error) {
	if p.done() {
		return nil, p.errorf("expected a value")
	}
	var s string
	var err error
	switch {
	case strings.HasPrefix(p.src[p.pos:], `"""`):
		s, err = p.multiline('"', true)
	case p.src[p.pos] == '"':
		s, err = p.basicString()
	case strings.HasPrefix(p.src[p.pos:], "'''"):
		s, err = p.multiline('\'', false)
	case p.src[p.pos] == '\'':
		s, err = p.literalString()
	case p.src[p.pos] == '[':
		return p.array()
	case p.src[p.pos] == '{':
		return p.inlineTable()
	default:
		return p.scalar()
	}
	if err != nil {
		return nil, err
	}
	return map[string]any{"type": "string", "value": s}, nil
}

// endOfLine consumes trailing spaces, an optional comment, and the line break.
// End of input also ends a line.
func (p *parser) endOfLine() error {
	p.skipSpaces()
	if !p.done() && p.src[p.pos] == '#' {
		p.pos++
		for !p.done() && !p.atNewline() {
			if isDisallowedControl(p.src[p.pos]) {
				return p.errorf("control character in comment")
			}
			p.pos++
		}
	}
	if p.done() {
		return nil
	}
	if p.src[p.pos] == '\n' {
		p.pos++
		return nil
	}
	if p.atCRLF() {
		p.pos += 2
		return nil
	}
	return p.errorf("unexpected text after value")
}

func (p *parser) skipSpaces() {
	for !p.done() && (p.src[p.pos] == ' ' || p.src[p.pos] == '\t') {
		p.pos++
	}
}

func (p *parser) done() bool {
	return p.pos >= len(p.src)
}

func (p *parser) atCRLF() bool {
	return strings.HasPrefix(p.src[p.pos:], "\r\n")
}

func (p *parser) atNewline() bool {
	return !p.done() && (p.src[p.pos] == '\n' || p.atCRLF())
}

func (p *parser) errorf(format string, args ...any) error {
	line := 1 + strings.Count(p.src[:p.pos], "\n")
	return fmt.Errorf("%w: line %d: %s", ErrSyntax, line, fmt.Sprintf(format, args...))
}

func isBareKeyByte(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_'
}

// isDisallowedControl reports whether c is a control character TOML forbids in
// strings and comments. Tab is the only control character it allows. Line breaks
// are excluded here because each caller decides whether a line break is legal.
func isDisallowedControl(c byte) bool {
	return (c < 0x20 && c != '\t' && c != '\n') || c == 0x7f
}
