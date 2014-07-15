package ace

// result represents a result of the parsing process.
type result struct {
	base     []element
	inner    []element
	includes map[string][]element
}
