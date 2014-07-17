package ace

import "strings"

// Prefixes of a line
const (
	prefixHelper = "="
)

// Helpers
const (
	helperInclude = "include"
)

// parseSource parses the source and returns the result.
func parseSource(src *source) *result {
	base := parseBytes(src.base.data)

	inner := parseBytes(src.inner.data)

	includes := make(map[string][]element)

	for _, f := range src.includes {
		includes[f.path] = parseBytes(f.data)
	}

	return newResult(base, inner, includes)
}

// parseBytes parses the byte data and returns the elements.
func parseBytes(data []byte) []element {
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
			e := newElement(ln)
			elements = append(elements, e)
		}
	}

	return elements
}
