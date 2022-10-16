package application

import (
	"context"

	"shopping-cart/api/typespec/producttype"
	"shopping-cart/service"
)

type Product struct{}

// GetProductList returns product list
func (p *Product) GetProductList(ctx context.Context, req *producttype.GetProductListRequest, resp *producttype.GetProductListResponse) error {
	productSvc := &service.Product{}
	if err := productSvc.GetMany(ctx, req, resp); err != nil {
		return err
	}

	return nil
}

// GetProduct returns data by criteria
func (p *Product) GetProduct(ctx context.Context, req *producttype.GetProductRequest, resp *producttype.GetProductResponse) error {
	productSvc := &service.Product{}
	if err := productSvc.Get(ctx, req, resp); err != nil {
		return err
	}

	return nil
}

// AddProduct adds product to database
func (p *Product) AddProduct(ctx context.Context, req *producttype.AddProductRequest, resp *producttype.AddProductResponse) error {
	productSvc := &service.Product{}
	if err := productSvc.Add(ctx, req, resp); err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates product by criteria
func (p *Product) UpdateProduct(ctx context.Context, req *producttype.UpdateProductRequest, resp *producttype.UpdateProductResponse) error {
	productSvc := &service.Product{}
	if err := productSvc.Update(ctx, req, resp); err != nil {
		return err
	}

	return nil
}

// DeleteProduct deletes product by criteria
func (p *Product) DeleteProduct(ctx context.Context, req *producttype.DeleteProductRequest, resp *producttype.DeleteProductResponse) error {
	productSvc := &service.Product{}
	if err := productSvc.Delete(ctx, req, resp); err != nil {
		return err
	}
	return nil
}
