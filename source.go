package ace

// source represents source for the parsing process.
type source struct {
	base     *file
	inner    *file
	includes []*file
}

// newSource creates and returns source.
func newSource(base, inner *file, includes []*file) *source {
	return &source{
		base:     base,
		inner:    inner,
		includes: includes,
	}
}
