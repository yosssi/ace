package ace

import (
	"fmt"
	"io"
)

// Helper method names
const (
	helperMethodNameContent    = "content"
	helperMethodNameCSS        = "css"
	helperMethodNameDoctype    = "doctype"
	helperMethodNameYield      = "yield"
	helperMethodNameInclude    = "include"
	helperMethodNameJavascript = "javascript"
)

// element is an interface for storing an element.
type element interface {
	io.WriterTo
	AppendChild(child element)
	ContainPlainText() bool
	Base() *elementBase
	CanHaveChildren() bool
}

// newElement creates and returns an element.
func newElement(ln *line, rslt *result, parent element) (element, error) {
	var e element
	var err error

	switch {
	case parent != nil && parent.ContainPlainText():
		e = newPlainTextInner(ln, rslt, parent)
	case ln.isEmpty():
		e = newEmptyElement(ln, rslt, parent)
	case ln.isComment():
		e = newComment(ln, rslt, parent)
	case ln.isHTMLComment():
		e = newHTMLComment(ln, rslt, parent)
	case ln.isConditionalComment():
		e = newConditionalComment(ln, rslt, parent)
	case ln.isHelperMethod():
		switch {
		case ln.isHelperMethodOf(helperMethodNameContent):
			e = newHelperMethodContent(ln, rslt, parent)
		case ln.isHelperMethodOf(helperMethodNameCSS):
			e = newHelperMethodCSS(ln, rslt, parent)
		case ln.isHelperMethodOf(helperMethodNameDoctype):
			e, err = newHelperMethodDoctype(ln, rslt, parent)
		case ln.isHelperMethodOf(helperMethodNameInclude):
			e = newHelperMethodInclude(ln, rslt, parent)
		case ln.isHelperMethodOf(helperMethodNameJavascript):
			e = newHelperMethodJavascript(ln, rslt, parent)
		case ln.isHelperMethodOf(helperMethodNameYield):
			e = newHelperMethodYield(ln, rslt, parent)
		default:
			err = fmt.Errorf("the helper method name is invalid [line: %d]", ln.no)
		}
	case ln.isPlainText():
		e = newPlainText(ln, rslt, parent)
	default:
		e, err = newHTMLTag(ln, rslt, parent)
	}

	return e, err
}
