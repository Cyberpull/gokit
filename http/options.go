package http

type Options struct {
	BaseURL string
}

func (o *Options) urlOf(uri string) string {
	if o.BaseURL == "" || isUrl(uri) {
		return uri
	}

	return join(o.BaseURL, uri)
}

// ===================

func defaultOptions(opts ...*Options) *Options {
	if len(opts) > 0 {
		return opts[0]
	}

	return &Options{}
}
