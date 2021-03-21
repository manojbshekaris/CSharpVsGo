package main

import (
	"github.com/fate-lovely/phi"
	"github.com/google/wire"
	"github.com/sample-fasthttp-rest-server/app/config"
	"github.com/sample-fasthttp-rest-server/app/handler/web"
	"github.com/sample-fasthttp-rest-server/app/server"
)

// wire set for loading the server.
var serverSet = wire.NewSet(
	provideRouter,
	provideServer,
)

// provideServer is a Wire provider function that returns an
// http server that is configured from the environment.
func provideServer(handler phi.Handler, config config.Config) *server.Server {
	return &server.Server{
		Host:    config.Server.Host,
		Handler: handler,
	}
}

// provideRouter is a Wire provider function that returns a
// router that is serves the provided handlers.
func provideRouter() phi.Handler {
	r := phi.NewRouter()
	r.Mount("/", web.Handler())
	return r
}
