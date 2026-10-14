package toml

// key reads a dotted key: one or more key parts joined by '.'. Spaces and tabs around
// each part are ignored. It returns the parts in order.
func (p *parser) key() ([]string, error) {
	var parts []string
	for {
		part, err := p.keyPart()
		if err != nil {
			return nil, err
		}
		parts = append(parts, part)
		p.skipSpaces()
		if p.done() || p.src[p.pos] != '.' {
			return parts, nil
		}
		p.pos++
		p.skipSpaces()
	}
}

// keyPart reads one bare key or one single-line quoted key.
func (p *parser) keyPart() (string, error) {
	if p.done() {
		return "", p.errorf("expected a key")
	}
	switch p.src[p.pos] {
	case '"':
		return p.basicString()
	case '\'':
		return p.literalString()
	}
	start := p.pos
	for !p.done() && isBareKeyByte(p.src[p.pos]) {
		p.pos++
	}
	if p.pos == start {
		return "", p.errorf("expected a key")
	}
	return p.src[start:p.pos], nil
}

// insert stores value at the dotted path parts below cur, creating each intermediate
// table the path implies. A key may be defined once, and a key that holds a value
// cannot later be used as a table.
func (p *parser) insert(cur *table, parts []string, value any) error {
	for _, part := range parts[:len(parts)-1] {
		next, err := p.dottedTable(cur, part)
		if err != nil {
			return err
		}
		cur = next
	}
	leaf := parts[len(parts)-1]
	if _, dup := cur.entries[leaf]; dup {
		return p.errorf("duplicate key %q", leaf)
	}
	cur.entries[leaf] = value
	return nil
}

// dottedTable steps from cur into the named table on the way to a dotted key's leaf. A
// dotted key may create a table or extend one that an earlier dotted key created. It may
// not extend a table a header defined or implied, since the header already fixed its
// contents.
func (p *parser) dottedTable(cur *table, name string) (*table, error) {
	next, ok := cur.entries[name]
	if !ok {
		t := newTable(originDotted)
		cur.entries[name] = t
		return t, nil
	}
	if t, ok := next.(*table); ok && t.origin == originDotted {
		return t, nil
	}
	return nil, p.errorf("key %q cannot be extended by a dotted key", name)
}
