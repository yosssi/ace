package ace

// elementBase holds common fields for the elements.
type elementBase struct {
	ln *line
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line) elementBase {
	return elementBase{
		ln: ln,
	}
}
