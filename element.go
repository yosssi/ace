package ace

import "io"

// Helper method names
const (
	helperMethodNameCSS        = "css"
	helperMethodNameDoctype    = "doctype"
	helperMethodNameYield      = "yield"
	helperMethodNameInclude    = "include"
	helperMethodNameJavascript = "javascript"
)

// element is an interface for storing an element.
type element interface {
	io.WriterTo
	appendChild(child element)
}

// newElement creates and returns an element.
func newElement(ln *line, rslt *result, parent element) (element, error) {
	var e element
	var err error

	switch {
	case isPlainTextInner(parent):
		e = newPlainTextInner(ln, rslt, parent)
	case ln.isComment():
		e = newComment(ln, rslt, parent)
	case ln.isHTMLComment():
		e = newHTMLComment(ln, rslt, parent)
	case ln.isConditionalComment():
		e = newConditionalComment(ln, rslt, parent)
	case ln.isHelperMethod():
		switch {
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
		}
	case ln.isPlainText():
		e = newPlainText(ln, rslt, parent)
	default:
		e, err = newHTMLTag(ln, rslt, parent)
	}

	return e, err
}

// isPlainTextInner returns true if the element is a plain text inner.
func isPlainTextInner(parent element) bool {
	var ok bool

	switch parent.(type) {
	case *helperMethodCSS:
		ok = true
	case *helperMethodJavascript:
		ok = true
	case *plainText:
		ok = true
	case *htmlTag:
		ok = parent.(*htmlTag).containPlainText
	}

	return ok
}
