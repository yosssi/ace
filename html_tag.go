package ace

import (
	"fmt"
	"io"
	"strings"
)

// Tag names
const (
	tagNameDiv = "div"
)

// htmlTag represents an HTML tag.
type htmlTag struct {
	elementBase
	tagName          string
	id               string
	classes          []string
	containPlainText bool
	attributes       []string
}

// WriteTo writes data to w.
func (e *htmlTag) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// ContainPlainText returns the HTML tag's containPlainText field.
func (e *htmlTag) ContainPlainText() bool {
	return e.containPlainText
}

// newHTMLTag creates and returns an HTML tag.
func newHTMLTag(ln *line, rslt *result, parent element) (*htmlTag, error) {
	if len(ln.tokens) < 1 {
		return nil, fmt.Errorf("an HTML tag is not specified [line: %d]", ln.no)
	}

	s := ln.tokens[0]

	tagName := extractTagName(s)

	id, err := extractID(s, ln)
	if err != nil {
		return nil, err
	}

	classes := extractClasses(s)

	containPlainText := doesContainPlainText(s)

	e := &htmlTag{
		elementBase:      newElementBase(ln, rslt, parent),
		tagName:          tagName,
		id:               id,
		classes:          classes,
		containPlainText: containPlainText,
	}

	return e, nil
}

// extractTag extracts and returns a tag.
func extractTagName(s string) string {
	tagName := strings.Split(strings.Split(s, sharp)[0], dot)[0]

	if tagName == "" {
		tagName = tagNameDiv
	}

	return tagName
}

// extractID extracts and returns an ID.
func extractID(s string, ln *line) (string, error) {
	tokens := strings.Split(s, sharp)

	l := len(tokens)

	if l < 2 {
		return "", nil
	}

	if l > 2 {
		return "", fmt.Errorf("multiple IDs are specified [line: %d]", ln.no)
	}

	return strings.Split(tokens[1], dot)[0], nil
}

// extractClasses extracts and returns classes.
func extractClasses(s string) []string {
	var classes []string

	for i, token := range strings.Split(s, dot) {
		if i == 0 {
			continue
		}

		class := strings.Split(token, sharp)[0]

		if class == "" {
			continue
		}

		classes = append(classes, class)
	}

	return classes
}

// doesContainPlainText returns true if the element contains
// a plain text.
func doesContainPlainText(s string) bool {
	return strings.HasSuffix(s, dot)
}
