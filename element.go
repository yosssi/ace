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
func newElement(ln *line, rslt *result, parent element, opts *Options) (element, error) {
	var e element
	var err error

	switch {
	case parent != nil && parent.ContainPlainText():
		e = newPlainTextInner(ln, rslt, parent, parent.InsertBr(), opts)
	case ln.isEmpty():
		e = newEmptyElement(ln, rslt, parent, opts)
	case ln.isComment():
		e = newComment(ln, rslt, parent, opts)
	case ln.isHTMLComment():
		e = newHTMLComment(ln, rslt, parent, opts)
	case ln.isHelperMethod():
		switch {
		case ln.isHelperMethodOf(helperMethodNameConditionalComment):
			e, err = newHelperMethodConditionalComment(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameContent):
			e = newHelperMethodContent(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameCSS):
			e = newHelperMethodCSS(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameDoctype):
			e, err = newHelperMethodDoctype(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameInclude):
			e = newHelperMethodInclude(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameJavascript):
			e = newHelperMethodJavascript(ln, rslt, parent, opts)
		case ln.isHelperMethodOf(helperMethodNameYield):
			e = newHelperMethodYield(ln, rslt, parent, opts)
		default:
			err = fmt.Errorf("the helper method name is invalid [line: %d]", ln.no)
		}
	case ln.isPlainText():
		e = newPlainText(ln, rslt, parent, opts)
	default:
		e, err = newHTMLTag(ln, rslt, parent, opts)
	}

	return e, err
}
