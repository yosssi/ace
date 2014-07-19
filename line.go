package ace

import "strings"

const unicodeSpace = 32

const indentTop = 0

const (
	strDoctypeHelperMethod = "doctype"
	strYieldHelperMethod   = "yield"
)

// line represents a line of codes.
type line struct {
	no     int
	str    string
	indent int
	tokens []string
}

// isEmpty returns true if the line is empty.
func (l *line) isEmpty() bool {
	return strings.TrimSpace(l.str) == ""
}

// isTopIndent returns true if the line's indent is the top level.
func (l *line) isTopIndent() bool {
	return l.indent == indentTop
}

// isHelperMethod returns true if the line is a helper method.
func (l *line) isHelperMethod() bool {
	return len(l.tokens) > 0 && l.tokens[0] == EQUAL
}

// isDoctypeHelperMethod returns true if the line is
// a doctype helper method.
func (l *line) isDoctypeHelperMethod() bool {
	return l.isHelperMethod() && len(l.tokens) > 1 && l.tokens[1] == strDoctypeHelperMethod
}

// isYieldHelperMethod returns true if the line is
// a yield helper method.
func (l *line) isYieldHelperMethod() bool {
	return l.isHelperMethod() && len(l.tokens) > 1 && l.tokens[1] == strYieldHelperMethod
}

// newLine creates and returns a line.
func newLine(no int, str string) *line {
	return &line{
		no:     no,
		str:    str,
		indent: indent(str),
		tokens: strings.Split(strings.TrimLeft(str, SPACE), SPACE),
	}
}

// indent returns the line's indent.
func indent(str string) int {
	var i int

	for _, b := range str {
		if b != unicodeSpace {
			break
		}
		i++
	}

	return i / 2
}
