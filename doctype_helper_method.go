package ace

// doctypeHelperMethod represents a doctype helper method.
type doctypeHelperMethod struct {
	elementBase
}

// newDoctypeHelperMethod creates and returns a doctype helper method.
func newDoctypeHelperMethod(ln *line) *doctypeHelperMethod {
	return &doctypeHelperMethod{
		elementBase: newElementBase(ln),
	}
}
