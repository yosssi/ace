package ace

// Defaults
const (
	defaultExtension  = "ace"
	defaultDelimLeft  = "{{"
	defaultDelimRight = "}}"
)

// Options represents options for the template engine.
type Options struct {
	// Extension represents an extension of files.
	Extension string
	// DelimLeft represents a left delimiter for the html template.
	DelimLeft string
	// DelimRight represents a right delimiter for the html template.
	DelimRight string
}

// initializeOptions initializes the options
func initializeOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Extension == "" {
		opts.Extension = defaultExtension
	}

	if opts.DelimLeft == "" {
		opts.DelimLeft = defaultDelimLeft
	}

	if opts.DelimRight == "" {
		opts.DelimRight = defaultDelimRight
	}

	return opts
}
