package ace

import "io"

// helperMethodContent represents a helper method content.
type helperMethodContent struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodContent) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHelperMethodContent creates and returns a helper method content.
func newHelperMethodContent(ln *line, rslt *result, parent element) *helperMethodContent {
	return &helperMethodContent{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
