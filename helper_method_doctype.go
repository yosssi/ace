package ace

import "io"

// helperMethodDoctype represents a helper method doctype.
type helperMethodDoctype struct {
	elementBase
}

// WriteTo writes data to w.
func (h *helperMethodDoctype) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// newHelperMethodDoctype creates and returns a helper method doctype.
func newHelperMethodDoctype(ln *line, rslt *result) *helperMethodDoctype {
	return &helperMethodDoctype{
		elementBase: newElementBase(ln, rslt),
	}
}
