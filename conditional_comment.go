package ace

import "io"

// conditionalComment represents an conditional comment.
type conditionalComment struct {
	elementBase
}

// WriteTo writes data to w.
func (e *conditionalComment) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newConditionalComment creates and returns an HTML comment.
func newConditionalComment(ln *line, rslt *result, parent element) *conditionalComment {
	return &conditionalComment{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
