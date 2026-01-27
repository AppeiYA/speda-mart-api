package handlers

import (
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/middlewares"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	s "e-commerce/internal/shared"
	"e-commerce/internal/utils"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserServ *services.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload models.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid request body", Errors: errs})
		return
	}

	err:= h.UserServ.CreateUser(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{Message: "User registered successfully"})
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	ctx:= r.Context()

	user, ok := middlewares.GetUserFromContext(ctx)
	if !ok {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User unauthorized"})
		return
	}

	getUser, err := h.UserServ.GetUserProfile(ctx, user.Email)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}
		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "Internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "User Profile fetched Successfully",
		Data: getUser,
	})
}
