package routes

import (
	"net/http"

	"github.com/hassamk122/bookmark_manager/internals/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
	SetupUserRoutes(mux, handler)
}
