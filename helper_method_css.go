package ace

import "io"

// helperMethodCSS represents a helper method css.
type helperMethodCSS struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodCSS) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// ContainPlainText returns true.
func (e *helperMethodCSS) ContainPlainText() bool {
	return true
}

// helperMethodCSS creates and returns a helper method css.
func newHelperMethodCSS(ln *line, rslt *result, parent element) *helperMethodCSS {
	return &helperMethodCSS{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
