package producttype

import (
	"shopping-cart/api/typespec"
)

type Product struct {
	Name  string  `form:"name" json:"name" validate:"required"`   //item name
	Price float32 `form:"price" json:"price" validate:"required"` //item price
	Total float32 `form:"total" json:"total" validate:"required"` //item amount
}

type AddProductRequest struct {
	Product
}

type AddProductResponse struct {
}

type GetProductListRequest struct {
	ID int `form:"id" json:"id"` //item ID
	typespec.PagerRequest
}

type GetProductListResponse struct {
	Offset int                  `json:"offset"` 
	Length int                  `json:"length"` 
	Total  int64                `json:"total"`  
	List   []GetProductResponse `json:"list"`
}

type GetProductRequest struct {
	ID int64 `form:"id" json:"id"` // ID
}

type GetProductResponse struct {
	ID int64 `json:"id"` // ID
	Product
}

type UpdateProductRequest struct {
	ID    int64   `form:"id" json:"id"` // ID
	Price float32 `form:"price" json:"price" validate:"required"` //price
}

type UpdateProductResponse struct {
}

type DeleteProductRequest struct {
	ID int64 `form:"id" json:"id"` // ID
}

type DeleteProductResponse struct {
}
