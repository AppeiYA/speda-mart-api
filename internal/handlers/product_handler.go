package handlers

import (
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	s "e-commerce/internal/shared"
	"e-commerce/internal/utils"
	"net/http"
)

type ProductHandler struct {
	ProductServ *services.ProductService
}

func(h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	// get query values 
	name, color, origin, minPrice, maxPrice, err := utils.ExtractProductFilters(r)
	if err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "invalid query parameters"})
		return
	}
	limit, offset := utils.ParsePagination(r)

	// put in payload object 
	payload := &models.GetProductsRequest{
		Name: name,
		Origin: origin,
		Color: color,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Limit: limit,
		Offset: offset,
	}

	// call services
	products, err := h.ProductServ.GetProducts(ctx, payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}
	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Products fetched successfully",
		Data: products,
	})
}

func(h *ProductHandler) GetProduct(w http.ResponseWriter, r*http.Request) {
	ctx := r.Context()

	params := utils.ExtractParams(r, "product_id")

	// call product service
	product, err := h.ProductServ.GetProduct(ctx, params["product_id"])
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}
	s.ReqResponse(w, 
		http.StatusOK, 
		s.Payload{
			Message: "Product fetched successfully", 
			Data: product,
		},
	)
}