package server

type OptFunc func(opts *Opts)

type Opts struct {
	maxConn int
	tls     bool
	id      string
}

func defaultOptions() Opts {
	return Opts{
		maxConn: 10,
		tls:     false,
		id:      "default server",
	}
}

func WithTLS(opts *Opts) {
	opts.tls = true
}

func WithMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func WithServerID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

type Server struct {
	opts Opts // Make opts private
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOptions()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		opts: o, // Assign to opts field instead of Opts
	}
}

func (s *Server) GetConfig() *Opts {
	return &s.opts
}
