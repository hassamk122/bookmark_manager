package routes

import (
	"net/http"

	"github.com/hassamk122/bookmark_manager/internals/handlers"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	userMux := http.NewServeMux()

	mux.Handle("/users/", http.StripPrefix("/users", userMux))

	userMux.Handle("POST /register", handler.CreateUserHandler())
	userMux.Handle("POST /login", handler.LoginUserHandler())
}
