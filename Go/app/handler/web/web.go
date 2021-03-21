package web

import (
	"github.com/fate-lovely/phi"
	"github.com/sample-fasthttp-rest-server/app/handler/web/home"
)

// Handler returns an http.Handler
func Handler() *phi.Mux {
	r := phi.NewRouter()

	r.Get("/text", home.Text)
	r.Get("/email", home.Email)
	return r
}
