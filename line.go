package ace

import "strings"

const unicodeSpace = 32

const indentTop = 0

// line represents a line of codes.
type line struct {
	no     int
	str    string
	indent int
}

// empty returns true if the line is empty.
func (l *line) empty() bool {
	return strings.TrimSpace(l.str) == ""
}

// topIndent returns true if the line's indent is the top level.
func (l *line) topIndent() bool {
	return l.indent == indentTop
}

// newLine creates and returns a line.
func newLine(no int, str string) *line {
	return &line{
		no:     no,
		str:    str,
		indent: indent(str),
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
