package ace

// yieldHelperMethod represents a yield helper method.
type yieldHelperMethod struct {
	elementBase
}

// newYieldHelperMethod creates and returns a yield helper method.
func newYieldHelperMethod(ln *line) *yieldHelperMethod {
	return &yieldHelperMethod{
		elementBase: newElementBase(ln),
	}
}
