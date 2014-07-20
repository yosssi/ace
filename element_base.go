package ace

// elementBase holds common fields for the elements.
type elementBase struct {
	ln   *line
	rslt *result
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line, rslt *result) elementBase {
	return elementBase{
		ln:   ln,
		rslt: rslt,
	}
}
