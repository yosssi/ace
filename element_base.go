package ace

// elementBase holds common fields for the elements.
type elementBase struct {
	ln       *line
	rslt     *result
	parent   element
	children []element
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

// newElementBase creates and returns an element base.
func newElementBase(ln *line, rslt *result, parent element) elementBase {
	return elementBase{
		ln:     ln,
		rslt:   rslt,
		parent: parent,
	}
}
