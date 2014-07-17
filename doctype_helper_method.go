package ace

// doctypeHelperMethod represents a doctype helper method.
type doctypeHelperMethod struct {
	elementBase
}

func newDoctypeHelperMethod(ln *line) *doctypeHelperMethod {
	return &doctypeHelperMethod{
		elementBase: newElementBase(ln),
	}
}
