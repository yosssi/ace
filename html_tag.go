package ace

import "io"

// htmlTag represents an HTML tag.
type htmlTag struct {
	elementBase
	containPlainText bool
}

// WriteTo writes data to w.
func (e *htmlTag) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHTMLTag creates and returns an HTML tag.
func newHTMLTag(ln *line, rslt *result, parent element) *htmlTag {
	return &htmlTag{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
