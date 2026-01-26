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

type ProductHandler struct {
	ProductServ *services.ProductService
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	if len(payload.Variants) == 0 {
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "a product must have at least one variant"})
		return
    }

	_, exists := middlewares.GetUserFromContext(ctx)
	if !exists {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User not authorized"})
		return
	}

	product, err := h.ProductServ.AddProduct(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}
		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: err.Error()})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Product added successfully",
		Data: product,
	})
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

func (h *ProductHandler) AddProductToCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload models.AddProductToCategory
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	_, exists := middlewares.GetUserFromContext(ctx)
	if !exists {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User not authorized"})
		return
	}

	err := h.ProductServ.AddProductToCategory(ctx, payload.ProductId, payload.CategoryId)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}
		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: err.Error()})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{Message: "Product added to category"})
}
func (h *ProductHandler) RemoveProductFromCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload models.RemoveProductFromCategory
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	_, exists := middlewares.GetUserFromContext(ctx)
	if !exists {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User not authorized"})
		return
	}

	err := h.ProductServ.RemoveProductFromCategory(ctx, payload.ProductId, payload.CategoryId)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}
		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: err.Error()})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{Message: "Product removed from category"})
}

func (h *ProductHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload models.CreateProductCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	_, exists := middlewares.GetUserFromContext(ctx)
	if !exists {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User not authorized"})
		return
	}

	category, err := h.ProductServ.CreateCategory(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}
		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: err.Error()})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Category created successfully",
		Data: category,
	})

}

func (h *ProductHandler) GetProductsInCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := utils.ExtractParams(r, "category_id")
	limit, offset := utils.ParsePagination(r)

	products, err := h.ProductServ.GetProductsInCategory(ctx, params["category_id"], limit, offset)
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
func (h *ProductHandler) GetSubCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := utils.ExtractParams(r, "category_id")

	products, err := h.ProductServ.GetSubCategories(ctx, params["category_id"])
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Categories fetched successfully",
		Data: products,
	})
}

func (h *ProductHandler) DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := utils.ExtractParams(r, "category_id")

	if _, ok := middlewares.GetUserFromContext(ctx); !ok {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message:"User unauthorized"})
		return
	}

	err := h.ProductServ.DeleteProductCategory(ctx, params["category_id"])
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Category deleted successfully",
	})
}

func (h *ProductHandler) UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload models.UpdateProductCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "no body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusUnprocessableEntity, s.Payload{Message: "invalid body content", Errors: errs})
		return
	}

	if _, ok := middlewares.GetUserFromContext(ctx); !ok {
		s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message:"User unauthorized"})
		return
	}

	updatedCategory, err := h.ProductServ.UpdateProductCategory(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok{
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusOK, s.Payload{
		Message: "Category updated successfully",
		Data: updatedCategory,
	})
}