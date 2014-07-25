package ace

import "io"

// helperMethodYield represents a helper method yield.
type helperMethodYield struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodYield) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHelperMethodYield creates and returns a helper method yield.
func newHelperMethodYield(ln *line, src *source, parent element, opts *Options) *helperMethodYield {
	return &helperMethodYield{
		elementBase: newElementBase(ln, src, parent, opts),
	}
}
