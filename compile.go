package ace

import (
	"bytes"
	"html/template"
)

// compileResult compiles the parsed result to the template.Template.
func compileResult(name string, rslt *result, opts *Options) (*template.Template, error) {
	// Create a buffer.
	b := bytes.NewBuffer(nil)

	// Write data to the buffer.
	for _, e := range rslt.base {
		if _, err := e.WriteTo(b); err != nil {
			return nil, err
		}
	}

	// Create a template.
	t := template.New(name)

	// Parse a string to the template.
	return t.Parse(b.String())
}
