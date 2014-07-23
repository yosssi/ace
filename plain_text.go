package ace

import "io"

// plainText represents a plain text.
type plainText struct {
	elementBase
}

// WriteTo writes data to w.
func (e *plainText) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// ContainPlainText returns true.
func (e *plainText) ContainPlainText() bool {
	return true
}

// newPlainText creates and returns a plain text.
func newPlainText(ln *line, rslt *result, parent element) *plainText {
	return &plainText{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
