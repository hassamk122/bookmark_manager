package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	customerr "github.com/hassamk122/bookmark_manager/internals/custom_err"
	"github.com/hassamk122/bookmark_manager/internals/dtos"
	"github.com/hassamk122/bookmark_manager/internals/utils"
	"github.com/hassamk122/bookmark_manager/internals/validation"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var userReq dtos.CreateUserRequest
		if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
			log.Println(userReq)
			utils.RespondWithError(res, http.StatusBadGateway, customerr.ErrInvalidRequestPayload.Error())
			return
		}

		if err := validation.Validate(&userReq); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, err.Error())
			return
		}

		err := h.UserService.Register(ctx, userReq.Username, userReq.Email, userReq.Password)
		if errors.Is(err, customerr.ErrEmailTaken) {
			utils.RespondWithError(res, http.StatusConflict, customerr.ErrEmailTaken.Error())
			return
		}

		if err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, customerr.ErrSomethingWentWrong.Error())
			return
		}

		utils.RespondWithSuccess(res, http.StatusCreated, "User created Successfully", nil)
	}
}
