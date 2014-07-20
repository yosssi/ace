package ace

import "strings"

// parseSource parses the source and returns the result.
func parseSource(src *source, opts *Options) (*result, error) {
	var rslt *result

	base, err := parseBytes(src.base.data, rslt, opts)
	if err != nil {
		return nil, err
	}

	inner, err := parseBytes(src.inner.data, rslt, opts)
	if err != nil {
		return nil, err
	}

	includes := make(map[string][]element)

	for _, f := range src.includes {
		includes[f.path], err = parseBytes(f.data, rslt, opts)
		if err != nil {
			return nil, err
		}
	}

	rslt = newResult(base, inner, includes)

	return rslt, nil
}

// parseBytes parses the byte data and returns the elements.
func parseBytes(data []byte, rslt *result, opts *Options) ([]element, error) {
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
		ln := newLine(i, lines[i])
		i++

		// Ignore the empty line.
		if ln.isEmpty() {
			continue
		}

		if ln.isTopIndent() {
			e, err := newElement(ln, rslt, nil)
			if err != nil {
				return nil, err
			}

			elements = append(elements, e)
		}
	}

	return elements, nil
}
