package ace

import "io"

// helperMethodJavascript represents a helper method javascript.
type helperMethodJavascript struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodJavascript) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// helperMethodJavascript creates and returns a helper method javascript.
func newHelperMethodJavascript(ln *line, rslt *result, parent element) *helperMethodJavascript {
	return &helperMethodJavascript{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
