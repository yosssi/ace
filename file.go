package ace

// file represents a file.
type file struct {
	path string
	data []byte
}

// NewFile creates and returns a file.
func NewFile(path string, data []byte) *file {
	return &file{
		path: path,
		data: data,
	}
}
