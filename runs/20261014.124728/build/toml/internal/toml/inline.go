package toml

// inlineTable reads an inline table at p.pos, which holds '{'. The table stays on one
// line, may be empty, and has no trailing comma. Its keys are read like any other dotted
// keys, and the table is marked closed, so no later key or header can extend it.
func (p *parser) inlineTable() (*table, error) {
	p.pos++
	t := newTable(originInline)
	p.skipSpaces()
	if !p.done() && p.src[p.pos] == '}' {
		p.pos++
		return t, nil
	}
	for {
		parts, value, err := p.keyValue()
		if err != nil {
			return nil, err
		}
		if err := p.insert(t, parts, value); err != nil {
			return nil, err
		}
		p.skipSpaces()
		if p.done() {
			return nil, p.errorf("unterminated inline table")
		}
		switch p.src[p.pos] {
		case '}':
			p.pos++
			return t, nil
		case ',':
			p.pos++
			p.skipSpaces()
		default:
			return nil, p.errorf("expected ',' or '}' in inline table")
		}
	}
}
