package ace

import "strings"

// parseSource parses the source and returns the result.
func parseSource(src *source, opts *Options) *result {
	var rslt *result

	base := parseBytes(src.base.data, rslt, opts)

	inner := parseBytes(src.inner.data, rslt, opts)

	includes := make(map[string][]element)

	for _, f := range src.includes {
		includes[f.path] = parseBytes(f.data, rslt, opts)
	}

	rslt = newResult(base, inner, includes)

	return rslt
}

// parseBytes parses the byte data and returns the elements.
func parseBytes(data []byte, rslt *result, opts *Options) []element {
	var elements []element

	lines := strings.Split(formatLF(string(data)), LF)

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
			e := newElement(ln, rslt)
			elements = append(elements, e)
		}
	}

	return elements
}
