package ace

// elementBase holds common fields for the elements.
type elementBase struct {
	ln       *line
	rslt     *result
	parent   element
	children []element
}

func (e *elementBase) appendChild(child element) {
	e.children = append(e.children, child)
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line, rslt *result, parent element) elementBase {
	return elementBase{
		ln:     ln,
		rslt:   rslt,
		parent: parent,
	}
}
