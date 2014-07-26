package ace

import "bytes"

// elementBase holds common fields for the elements.
type elementBase struct {
	ln       *line
	rslt     *result
	src      *source
	parent   element
	children []element
	opts     *Options
}

// AppendChild appends the child element to the element.
func (e *elementBase) AppendChild(child element) {
	e.children = append(e.children, child)
}

// ContainPlainText returns false.
// This method should be overrided by a struct which contains
// the element base struct.
func (e *elementBase) ContainPlainText() bool {
	return false
}

// Base returns the element base.
func (e *elementBase) Base() *elementBase {
	return e
}

// CanHaveChildren returns true.
// This method should be overrided by a struct which contains
// the element base struct.
func (e *elementBase) CanHaveChildren() bool {
	return true
}

// InsertBr returns false.
// This method should be overrided by a struct which contains
// the element base struct.
func (e *elementBase) InsertBr() bool {
	return false
}

// writeChildren writes the children's HTML.
func (e *elementBase) writeChildren(bf *bytes.Buffer) (int64, error) {
	for _, child := range e.children {
		if i, err := child.WriteTo(bf); err != nil {
			return int64(i), err
		}
	}
	return 0, nil
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line, rslt *result, src *source, parent element, opts *Options) elementBase {
	return elementBase{
		ln:     ln,
		rslt:   rslt,
		src:    src,
		parent: parent,
		opts:   opts,
	}
}
