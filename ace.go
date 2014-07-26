package ace

import (
	"html/template"
	"sync"
)

var cache = make(map[string]template.Template)
var cacheMutex = new(sync.RWMutex)

// ParseFiles parses template files and returns an HTML template.
func ParseFiles(basePath, innerPath string, opts *Options) (*template.Template, error) {
	// Initialize the options.
	opts = initializeOptions(opts)

	name := basePath + colon + innerPath

	if opts.Cache {
		if tpl, ok := getCache(name); ok {
			return &tpl, nil
		}
	}

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
	tpl, err := compileResult(name, rslt, opts)
	if err != nil {
		return nil, err
	}

	if opts.Cache {
		setCache(name, *tpl)
	}

	return tpl, nil
}

// getCache returns the cached template.
func getCache(name string) (template.Template, bool) {
	cacheMutex.RLock()
	tpl, ok := cache[name]
	cacheMutex.RUnlock()
	return tpl, ok
}

// setCache sets the template to the cache.
func setCache(name string, tpl template.Template) {
	cacheMutex.Lock()
	cache[name] = tpl
	cacheMutex.Unlock()
}
