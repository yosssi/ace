package ace

import "strings"

// parseSource parses the source and returns the result.
func parseSource(src *source, opts *Options) (*result, error) {
	base, err := parseBytes(src.base.data, src, opts)
	if err != nil {
		return nil, err
	}

	inner, err := parseBytes(src.inner.data, src, opts)
	if err != nil {
		return nil, err
	}

	includes := make(map[string][]element)

	for _, f := range src.includes {
		includes[f.path], err = parseBytes(f.data, src, opts)
		if err != nil {
			return nil, err
		}
	}

	return newResult(base, inner, includes), nil
}

// parseBytes parses the byte data and returns the elements.
func parseBytes(data []byte, src *source, opts *Options) ([]element, error) {
	var elements []element

	lines := strings.Split(formatLF(string(data)), lf)

	i := 0
	l := len(lines)

	// Ignore the last empty line.
	if l > 0 && lines[l-1] == "" {
		l--
	}

	for i < l {
		// Fetch a line.
		ln := newLine(i+1, lines[i], opts)
		i++

		// Ignore the empty line.
		if ln.isEmpty() {
			continue
		}

		if ln.isTopIndent() {
			e, err := newElement(ln, src, nil, opts)
			if err != nil {
				return nil, err
			}

			// Append child elements to the element.
			if err := appendChildren(e, lines, &i, l, src, opts); err != nil {
				return nil, err
			}

			elements = append(elements, e)
		}
	}

	return elements, nil
}

// appendChildren parses the lines and appends the children to the element.
func appendChildren(parent element, lines []string, i *int, l int, src *source, opts *Options) error {
	for *i < l {
		// Fetch a line.
		ln := newLine(*i+1, lines[*i], opts)

		// Check if the line is a child of the parent.
		ok, err := ln.childOf(parent)

		if err != nil {
			return err
		}

		if !ok {
			return nil
		}

		child, err := newElement(ln, src, parent, opts)
		if err != nil {
			return err
		}

		parent.AppendChild(child)

		*i++

		if child.CanHaveChildren() {
			if err := appendChildren(child, lines, i, l, src, opts); err != nil {
				return err
			}
		}
	}

	return nil
}
