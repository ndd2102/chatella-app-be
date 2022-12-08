package static

import (
	"net/http"
	"web-service/pkg/server"

	"github.com/go-chi/chi"
)

var fePath = server.Config.GetString("FE_PATH")
var fileServer = http.FileServer(http.Dir(fePath))

var StaticSubRoute = chi.NewRouter()

// Init package with sub-router for mails service
func init() {
	StaticSubRoute.Group(func(_ chi.Router) {
		StaticSubRoute.Handle("/*", fileServer)
	})
}
