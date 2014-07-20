package ace

import "io"

// Helper method names
const (
	helperMethodNameDoctype = "doctype"
	helperMethodNameYield   = "yield"
	helperMethodNameInclude = "include"
)

// element is an interface for storing an element.
type element interface {
	io.WriterTo
}

// newElement creates and returns an element.
func newElement(ln *line, rslt *result) element {
	switch {
	case ln.isHelperMethod():
		switch {
		case ln.isHelperMethodOf(helperMethodNameDoctype):
			return newHelperMethodDoctype(ln, rslt)
		case ln.isHelperMethodOf(helperMethodNameYield):
			return newHelperMethodYield(ln, rslt)
		case ln.isHelperMethodOf(helperMethodNameInclude):
			return newHelperMethodInclude(ln, rslt)
		}
	}
	return nil
}
