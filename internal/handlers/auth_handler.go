package handlers

import (
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	s "e-commerce/internal/shared"
	"e-commerce/internal/utils"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthServ *services.AuthService
}

func (a *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	resp, token, err := a.AuthServ.Login(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{Message: "Login successful", Data: resp, Token: token})
}