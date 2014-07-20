package ace

import "io"

// helperMethodYield represents a helper method yield.
type helperMethodYield struct {
	elementBase
}

// WriteTo writes data to w.
func (h *helperMethodYield) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// newHelperMethodYield creates and returns a helper method yield.
func newHelperMethodYield(ln *line, rslt *result) *helperMethodYield {
	return &helperMethodYield{
		elementBase: newElementBase(ln, rslt),
	}
}
