package ace

import "html/template"

// ParseFiles parses template files and returns an HTML template.
func ParseFiles(basePath, innerPath string, opts *Options) (*template.Template, error) {
	// Initialize the options.
	opts = initializeOptions(opts)

	// Read files.
	src, err := readFiles(basePath, innerPath, opts)
	if err != nil {
		return nil, err
	}

	// Parse the source.
	rslt, err := parseSource(src, opts)
	if err != nil {
		return nil, err
	}

	// Compile the parsed result.
	return compileResult(basePath+colon+innerPath, rslt, opts)
}
