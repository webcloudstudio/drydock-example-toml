package toml

// arrayValue is a TOML array while a document is read. Elements may mix types and may
// themselves be arrays or inline tables. export turns it into a plain JSON array, since
// the tagged {"type", "value"} form is for scalars only.
type arrayValue []any

// array reads a bracketed array at p.pos, which holds '['. Line breaks and comments may
// appear between elements, and one trailing comma is allowed.
func (p *parser) array() (arrayValue, error) {
	p.pos++
	elems := arrayValue{}
	for {
		if err := p.skipArrayGap(); err != nil {
			return nil, err
		}
		if p.done() {
			return nil, p.errorf("unterminated array")
		}
		if p.src[p.pos] == ']' {
			p.pos++
			return elems, nil
		}
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		elems = append(elems, v)
		if err := p.skipArrayGap(); err != nil {
			return nil, err
		}
		if p.done() {
			return nil, p.errorf("unterminated array")
		}
		switch p.src[p.pos] {
		case ',':
			p.pos++
		case ']':
			// The top of the loop consumes the closing bracket.
		default:
			return nil, p.errorf("expected ',' or ']' in array")
		}
	}
}

// skipArrayGap consumes the spaces, line breaks, and comments that may separate array
// elements. A comment runs to the end of its line and is subject to the same control
// character rule as any other comment.
func (p *parser) skipArrayGap() error {
	for !p.done() {
		switch {
		case p.src[p.pos] == ' ' || p.src[p.pos] == '\t' || p.src[p.pos] == '\n':
			p.pos++
		case p.atCRLF():
			p.pos += 2
		case p.src[p.pos] == '#':
			p.pos++
			for !p.done() && !p.atNewline() {
				if isDisallowedControl(p.src[p.pos]) {
					return p.errorf("control character in comment")
				}
				p.pos++
			}
		default:
			return nil
		}
	}
	return nil
}
