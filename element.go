package ace

import (
	"fmt"
	"io"
)

// Helper method names
const (
	helperMethodNameConditionalComment = "conditionalComment"
	helperMethodNameContent            = "content"
	helperMethodNameCSS                = "css"
	helperMethodNameDoctype            = "doctype"
	helperMethodNameYield              = "yield"
	helperMethodNameInclude            = "include"
	helperMethodNameJavascript         = "javascript"
)

// element is an interface for storing an element.
type element interface {
	io.WriterTo
	AppendChild(child element)
	ContainPlainText() bool
	Base() *elementBase
	CanHaveChildren() bool
	InsertBr() bool
}

// newElement creates and returns an element.
func newElement(ln *line, src *source, parent element, opts *Options) (element, error) {
	var e element
	var err error

	switch {
	case parent != nil && parent.ContainPlainText():
		e = newPlainTextInner(ln, src, parent, parent.InsertBr(), opts)
	case ln.isEmpty():
		e = newEmptyElement(ln, src, parent, opts)
	case ln.isComment():
		e = newComment(ln, src, parent, opts)
	case ln.isHTMLComment():
		e = newHTMLComment(ln, src, parent, opts)
	case ln.isHelperMethod():
		switch {
		case ln.isHelperMethodOf(helperMethodNameConditionalComment):
			e, err = newHelperMethodConditionalComment(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameContent):
			e, err = newHelperMethodContent(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameCSS):
			e = newHelperMethodCSS(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameDoctype):
			e, err = newHelperMethodDoctype(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameInclude):
			e, err = newHelperMethodInclude(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameJavascript):
			e = newHelperMethodJavascript(ln, src, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameYield):
			e = newHelperMethodYield(ln, src, parent, opts)
		default:
			err = fmt.Errorf("the helper method name is invalid [line: %d]", ln.no)
		}
	case ln.isPlainText():
		e = newPlainText(ln, src, parent, opts)
	case ln.isAction():
		e = newAction(ln, src, parent, opts)
	default:
		e, err = newHTMLTag(ln, src, parent, opts)
	}

	return e, err
}
