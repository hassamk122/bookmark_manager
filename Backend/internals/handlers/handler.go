package handlers

import "github.com/hassamk122/bookmark_manager/internals/services"

type Handler struct {
	UserService services.UserService
}

func NewHandler(userService services.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}
