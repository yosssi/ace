package ace

// Defaults
const (
	defaultExtension = "ace"
)

// Options represents options for the template engine.
type Options struct {
	// Extension represents an extension of files.
	Extension string
}

// initializeOptions initializes the options
func initializeOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Extension == "" {
		opts.Extension = defaultExtension
	}

	return opts
}
