package main

import "fmt"

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

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withServerID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

type Server struct {
	Opts
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOptions()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	defaultServer := NewServer()
	fmt.Printf("default server: %+v\n", defaultServer)

	s := NewServer(withTLS, withMaxConn(45), withServerID("local-server"))
	fmt.Printf("server: %+v\n", s)
}
