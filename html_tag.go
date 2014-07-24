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

// Attribute names
const (
	attributeNameID    = "id"
	attributeNameClass = "class"
)

// htmlTag represents an HTML tag.
type htmlTag struct {
	elementBase
	tagName          string
	id               string
	classes          []string
	containPlainText bool
	attributes       map[string]string
	textValue        string
}

// WriteTo writes data to w.
func (e *htmlTag) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// ContainPlainText returns the HTML tag's containPlainText field.
func (e *htmlTag) ContainPlainText() bool {
	return e.containPlainText
}

// setAttributes parses the tokens and set attributes to the element.
func (e *htmlTag) setAttributes() error {
	var parsedTokens []string
	var unclosedTokenValues []string
	var unclosed bool
	var closeMark string

	for _, token := range e.ln.tokens {
		if unclosed {
			unclosedTokenValues = append(unclosedTokenValues, token)

			if closed(token, closeMark) {
				parsedTokens = append(parsedTokens, strings.Join(unclosedTokenValues, space))
				unclosedTokenValues = make([]string, 0)
				unclosed = false
			}
		} else {
			if unclosed, closeMark = unclosedToken(token, e.opts); unclosed {
				unclosedTokenValues = append(unclosedTokenValues, token)
			} else {
				parsedTokens = append(parsedTokens, token)
			}
		}
	}

	if unclosed {
		parsedTokens = append(parsedTokens, unclosedTokenValues...)
	}

	var i int
	var token string

	// Set attributes to the element.
	for i, token = range parsedTokens {
		kv := strings.Split(token, equal)

		if len(kv) < 2 {
			break
		}

		k := kv[0]
		v := kv[1]

		// Remove the prefix and suffix of the double quotes.
		if len(v) > 1 && strings.HasPrefix(v, doubleQuote) && strings.HasSuffix(v, doubleQuote) {
			v = v[1 : len(v)-1]
		}

		switch k {
		case attributeNameID:
			if e.id != "" {
				return fmt.Errorf("multiple IDs are specified [line: %d]", e.ln.no)
			}

			e.id = v
		case attributeNameClass:
			e.classes = append(e.classes, v)
		default:
			e.attributes[k] = v
		}
	}

	// Set a text value to the element.
	e.textValue = strings.Join(parsedTokens[i:], space)

	return nil
}

// newHTMLTag creates and returns an HTML tag.
func newHTMLTag(ln *line, rslt *result, parent element, opts *Options) (*htmlTag, error) {
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
		elementBase:      newElementBase(ln, rslt, parent, opts),
		tagName:          tagName,
		id:               id,
		classes:          classes,
		containPlainText: containPlainText,
	}

	if err := e.setAttributes(); err != nil {
		return nil, err
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

// unclosedToken returns true if the token is unclosed.
func unclosedToken(s string, opts *Options) (bool, string) {
	if len(strings.Split(s, doubleQuote)) == 2 {
		return true, doubleQuote
	}

	if len(strings.Split(s, opts.DelimLeft)) == 2 {
		return true, opts.DelimRight
	}

	return false, ""
}

// closedToken returns true if the token is closed.
func closed(s string, closeMark string) bool {
	return strings.HasSuffix(s, closeMark)
}
