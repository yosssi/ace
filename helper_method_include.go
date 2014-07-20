package ace

import "io"

// helperMethodInclude represents a helper method include.
type helperMethodInclude struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodInclude) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHelperMethodInclude creates and returns a helper method include.
func newHelperMethodInclude(ln *line, rslt *result, parent element) *helperMethodInclude {
	return &helperMethodInclude{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
