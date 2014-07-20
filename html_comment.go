package ace

import "io"

// htmlComment represents an HTML comment.
type htmlComment struct {
	elementBase
}

// WriteTo writes data to w.
func (e *htmlComment) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHTMLComment creates and returns an HTML comment.
func newHTMLComment(ln *line, rslt *result, parent element) *htmlComment {
	return &htmlComment{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
