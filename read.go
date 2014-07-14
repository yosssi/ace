package ace

import (
	"io/ioutil"
	"strings"
)

// Line feed characters and spaces
const (
	CR   = "\r"
	LF   = "\n"
	CRLF = "\r\n"

	SPACE = " "
)

// readFiles reads files and returns source for the parsing process.
func readFiles(bathPath, innerPath string, opts *Options) (*source, error) {
	// Read the base file.
	base, err := readFile(bathPath, opts)
	if err != nil {
		return nil, err
	}

	// Read the inner file.
	inner, err := readFile(innerPath, opts)
	if err != nil {
		return nil, err
	}

	var includes []*file

	// Find include files from the base file.
	if err := findIncludes(base.data, opts, &includes); err != nil {
		return nil, err
	}

	// Find include files from the inner file.
	if err := findIncludes(inner.data, opts, &includes); err != nil {
		return nil, err
	}

	return newSource(base, inner, includes), nil
}

// readFile reads a file and returns a file struct.
func readFile(path string, opts *Options) (*file, error) {
	data, err := ioutil.ReadFile(path + "." + opts.Extension)
	if err != nil {
		return nil, err
	}

	return newFile(path, data), nil
}

// findIncludes finds and adds include files.
func findIncludes(data []byte, opts *Options, includes *[]*file) error {
	for _, includePath := range findIncludePaths(data) {
		if !hasFile(*includes, includePath) {
			f, err := readFile(includePath, opts)
			if err != nil {
				return err
			}

			*includes = append(*includes, f)

			if err := findIncludes(f.data, opts, includes); err != nil {
				return err
			}
		}
	}

	return nil
}

// findIncludePaths finds and returns include paths.
func findIncludePaths(data []byte) []string {
	var includePaths []string

	for _, line := range strings.Split(formatLF(string(data)), LF) {
		if includePath := findIncludePath(line); includePath != "" {
			includePaths = append(includePaths, includePath)
		}
	}

	return includePaths
}

// findIncludePath finds and returns an include path.
func findIncludePath(line string) string {
	tokens := strings.Split(strings.TrimSpace(line), SPACE)

	if len(tokens) < 3 || tokens[0] != prefixHelper || tokens[1] != helperInclude {
		return ""
	}

	return tokens[2]
}

// formatLF replaces the line feed codes with LF and returns the result.
func formatLF(s string) string {
	return strings.Replace(strings.Replace(s, CRLF, LF, -1), CR, LF, -1)
}

// hasFile return if files has a file which has the path specified by the parameter.
func hasFile(files []*file, path string) bool {
	for _, f := range files {
		if f.path == path {
			return true
		}
	}

	return false
}
