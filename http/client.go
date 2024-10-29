package http

type Client[T any] interface {
	Get(uri string, opts ...*RequestOptions) (data T, err error)
	Post(uri string, opts ...*RequestOptions) (data T, err error)
	Patch(uri string, opts ...*RequestOptions) (data T, err error)
	Put(uri string, opts ...*RequestOptions) (data T, err error)
	Delete(uri string, opts ...*RequestOptions) (data T, err error)
	Request(method, uri string, opts ...*RequestOptions) (data T, err error)
}

type httpClient[T any] struct {
	opts *Options
}

func (c httpClient[T]) Get(uri string, opts ...*RequestOptions) (data T, err error) {
	return Get[T](c.opts.urlOf(uri), opts...)
}

func (c httpClient[T]) Post(uri string, opts ...*RequestOptions) (data T, err error) {
	return Post[T](c.opts.urlOf(uri), opts...)
}

func (c httpClient[T]) Patch(uri string, opts ...*RequestOptions) (data T, err error) {
	return Patch[T](c.opts.urlOf(uri), opts...)
}

func (c httpClient[T]) Put(uri string, opts ...*RequestOptions) (data T, err error) {
	return Put[T](c.opts.urlOf(uri), opts...)
}

func (c httpClient[T]) Delete(uri string, opts ...*RequestOptions) (data T, err error) {
	return Delete[T](c.opts.urlOf(uri), opts...)
}

func (c httpClient[T]) Request(method, uri string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](method, c.opts.urlOf(uri), opts...)
}

// ==========================

func NewClient[T any](opts ...*Options) Client[T] {
	return &httpClient[T]{
		opts: defaultOptions(opts...),
	}
}
