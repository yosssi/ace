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
	parseBytes(src.base.data)

	parseBytes(src.inner.data)

	for _, f := range src.includes {
		parseBytes(f.data)
	}

	return nil
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
		}
	}

	return elements
}
