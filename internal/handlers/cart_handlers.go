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

type CartsHandler struct {
	CartService *services.CartsService
}

func (h *CartsHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := middlewares.GetUserFromContext(ctx)
	if !ok {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "Unauthorized user"})
		return
	}

	var payload models.AddToCart
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	err := h.CartService.AddToCart(ctx, user.UserId, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}
	s.ReqResponse(w, http.StatusOK, s.Payload{Message: "Item added to cart"})
}