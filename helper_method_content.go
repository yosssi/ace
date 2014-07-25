package ace

import (
	"bytes"
	"io"
)

// helperMethodContent represents a helper method content.
type helperMethodContent struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodContent) WriteTo(w io.Writer) (int64, error) {
	var bf bytes.Buffer

	// Write the children's HTML.
	if i, err := e.writeChildren(&bf); err != nil {
		return i, err
	}

	// Write the buffer.
	i, err := w.Write(bf.Bytes())

	return int64(i), err
}

// newHelperMethodContent creates and returns a helper method content.
func newHelperMethodContent(ln *line, rslt *result, parent element, opts *Options) *helperMethodContent {
	return &helperMethodContent{
		elementBase: newElementBase(ln, rslt, parent, opts),
	}
}
