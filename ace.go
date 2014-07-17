package ace

import (
	"fmt"
	"html/template"
)

// ParseFiles parses template files and returns an HTML template.
func ParseFiles(bathPath, innerPath string, opts *Options) (*template.Template, error) {
	// Initialize the options.
	opts = initializeOptions(opts)

	src, err := readFiles(bathPath, innerPath, opts)
	if err != nil {
		return nil, err
	}

	rslt := parseSource(src)

	fmt.Printf("%+v\n", rslt.base[0])

	return nil, nil
}
