package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hassamk122/bookmark_manager/internals/dtos"
	"github.com/hassamk122/bookmark_manager/internals/errors"
	"github.com/hassamk122/bookmark_manager/internals/utils"
	"github.com/hassamk122/bookmark_manager/internals/validation"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var userReq dtos.CreateUserRequest
		if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
			utils.RespondWithError(res, http.StatusBadGateway, errors.ErrInvalidRequestPayload)
			return
		}

		if err := validation.Validate(&userReq); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, err.Error())
			return
		}

		// have to create service to register
	}
}
