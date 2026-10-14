package toml

// origin records how a table came to exist. A later [header] or dotted key consults
// it to decide whether it may add to or define the table.
type origin int

const (
	// originImplicit marks a super-table created only as a step on a longer header
	// path, such as a for [a.b]. A later [a] header may still define it.
	originImplicit origin = iota
	// originHeader marks a table defined by its own [name] header, or an element
	// created by a [[name]] header. No later header may define it again.
	originHeader
	// originDotted marks a table created by a dotted key. Later dotted keys may extend
	// it, but no header may define it.
	originDotted
	// originInline marks an inline table. It is closed: neither a dotted key nor a
	// header may add a key to it once its braces have closed.
	originInline
)

// table holds the entries of one TOML table while a document is read. Each entry is a
// *table, a tableArray, or the tagged map that a value returns. The three dynamic types
// differ, so a table is never mistaken for a value.
type table struct {
	entries map[string]any
	origin  origin
}

// tableArray is the value of a [[name]] key. Each element is one table, defined by one
// header, in document order.
type tableArray []*table

func newTable(o origin) *table {
	return &table{entries: map[string]any{}, origin: o}
}

// header reads a [name] or [[name]] line through its closing brackets. It returns the
// key parts and whether the header is an array header. The caller ends the line.
func (p *parser) header() ([]string, bool, error) {
	p.pos++
	array := !p.done() && p.src[p.pos] == '['
	if array {
		p.pos++
	}
	p.skipSpaces()
	parts, err := p.key()
	if err != nil {
		return nil, false, err
	}
	p.skipSpaces()
	if err := p.expectByte(']'); err != nil {
		return nil, false, err
	}
	if array {
		if err := p.expectByte(']'); err != nil {
			return nil, false, err
		}
	}
	return parts, array, nil
}

func (p *parser) expectByte(c byte) error {
	if p.done() || p.src[p.pos] != c {
		return p.errorf("expected %q in table header", c)
	}
	p.pos++
	return nil
}

// enterHeader finds or creates the table that a header names below root, and returns it
// as the new scope for key/value pairs. Every rule that makes a header invalid is
// checked here, so an accepted document never redefines a table.
func (p *parser) enterHeader(root *table, parts []string, array bool) (*table, error) {
	cur := root
	for _, name := range parts[:len(parts)-1] {
		next, err := p.headerTable(cur, name)
		if err != nil {
			return nil, err
		}
		cur = next
	}
	name := parts[len(parts)-1]
	existing, ok := cur.entries[name]
	if !ok {
		t := newTable(originHeader)
		if array {
			cur.entries[name] = tableArray{t}
		} else {
			cur.entries[name] = t
		}
		return t, nil
	}
	switch v := existing.(type) {
	case *table:
		if array || v.origin != originImplicit {
			return nil, p.errorf("table %q is already defined", name)
		}
		v.origin = originHeader
		return v, nil
	case tableArray:
		if !array {
			return nil, p.errorf("%q is an array of tables and cannot be a table", name)
		}
		t := newTable(originHeader)
		cur.entries[name] = append(v, t)
		return t, nil
	}
	return nil, p.errorf("key %q holds a value and cannot be a table", name)
}

// headerTable steps from cur into the named child on the way to a header's last part.
// A missing child becomes an implicit super-table. An array of tables is entered at its
// most recent element, which is where the next sub-table belongs.
func (p *parser) headerTable(cur *table, name string) (*table, error) {
	next, ok := cur.entries[name]
	if !ok {
		t := newTable(originImplicit)
		cur.entries[name] = t
		return t, nil
	}
	switch v := next.(type) {
	case *table:
		if v.origin == originInline {
			return nil, p.errorf("inline table %q is closed and cannot be extended", name)
		}
		return v, nil
	case tableArray:
		return v[len(v)-1], nil
	}
	return nil, p.errorf("key %q holds a value and cannot be used as a table", name)
}

// export converts the table tree to plain maps and slices for the caller. Leaves are
// shared, not copied, because each one is already its final tagged form.
func (t *table) export() map[string]any {
	out := make(map[string]any, len(t.entries))
	for k, v := range t.entries {
		out[k] = exportEntry(v)
	}
	return out
}

func exportEntry(v any) any {
	switch v := v.(type) {
	case *table:
		return v.export()
	case tableArray:
		out := make([]any, len(v))
		for i, el := range v {
			out[i] = el.export()
		}
		return out
	case arrayValue:
		out := make([]any, len(v))
		for i, el := range v {
			out[i] = exportEntry(el)
		}
		return out
	}
	return v
}
