package ace

// element is an interface for storing an element.
type element interface{}

// newElement creates and returns an element.
func newElement(ln *line) element {
	switch {
	case ln.isDoctypeHelperMethod():
		return newDoctypeHelperMethod(ln)
	}
	return nil
}
