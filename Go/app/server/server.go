package server

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
)

// A Server defines parameters for running an HTTP server.
type Server struct {
	Host    string
	Handler phi.Handler
}

// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe() error {
	var g errgroup.Group
	g.Go(func() error {
		return fasthttp.ListenAndServe(s.Host, s.Handler.ServeFastHTTP)
	})
	return g.Wait()
}
