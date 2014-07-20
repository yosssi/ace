package ace

import "strings"

const unicodeSpace = 32

const indentTop = 0

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
	return len(l.tokens) > 1 && l.tokens[0] == equal
}

// isHelperMethodOf returns true if the line is a specified helper method.
func (l *line) isHelperMethodOf(name string) bool {
	return l.isHelperMethod() && l.tokens[1] == name
}

// isPlainText returns true if the line is a plain text.
func (l *line) isPlainText() bool {
	return len(l.tokens) > 0 && l.tokens[0] == pipe
}

// isComment returns true if the line is a comment.
func (l *line) isComment() bool {
	return len(l.tokens) > 0 && l.tokens[0] == slash
}

// isHTMLComment returns true if the line is an HTML comment.
func (l *line) isHTMLComment() bool {
	return len(l.tokens) > 0 && l.tokens[0] == slash+slash
}

// isConditionalComment returns true if the line is a conditional comment.
func (l *line) isConditionalComment() bool {
	return len(l.tokens) > 0 && l.tokens[0] == slash+slash+slash
}

// newLine creates and returns a line.
func newLine(no int, str string) *line {
	return &line{
		no:     no,
		str:    str,
		indent: indent(str),
		tokens: strings.Split(strings.TrimLeft(str, space), space),
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
