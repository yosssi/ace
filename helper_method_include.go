package ace

import "io"

// helperMethodInclude represents a helper method include.
type helperMethodInclude struct {
	elementBase
}

// WriteTo writes data to w.
func (h *helperMethodInclude) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// newHelperMethodInclude creates and returns a helper method include.
func newHelperMethodInclude(ln *line, rslt *result) *helperMethodInclude {
	return &helperMethodInclude{
		elementBase: newElementBase(ln, rslt),
	}
}
