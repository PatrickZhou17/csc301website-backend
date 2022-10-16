package service

import (
	"context"
	"shopping-cart/api/typespec/producttype"
	"shopping-cart/common/util"
	"shopping-cart/model"
)

type Product struct{}

// GetMany returns product list
func (p *Product) GetMany(ctx context.Context, req *producttype.GetProductListRequest, resp *producttype.GetProductListResponse) error {
	var (
		productList []model.Product
		product     model.Product
	)

	if err := util.Bind(&product, req); err != nil {
		return err
	}

	productList, count, err := product.GetProductList(ctx, req.Offset, req.Length)
	if err != nil {
		return err
	}

	err = util.Bind(&resp.List, productList)
	if err != nil {
		return err
	}

	resp.Offset = req.Offset
	resp.Length = req.Length
	resp.Total = count

	return nil
}

// Get returns a single product data
func (p *Product) Get(ctx context.Context, req *producttype.GetProductRequest, resp *producttype.GetProductResponse) error {
	var product model.Product

	if err := util.Bind(&product, req); err != nil {
		return err
	}

	product, err := product.GetProduct(ctx)
	if err != nil {
		return err
	}

	return util.Bind(resp, product)
}

// Add adds a record of product
func (p *Product) Add(ctx context.Context, req *producttype.AddProductRequest, resp *producttype.AddProductResponse) error {
	var product model.Product

	if err := util.Bind(&product, req); err != nil {
		return err
	}

	return product.AddProduct(ctx)
}

// Update updates a record of product
func (p *Product) Update(ctx context.Context, req *producttype.UpdateProductRequest, resp *producttype.UpdateProductResponse) error {
	var product model.Product

	if err := util.Bind(&product, req); err != nil {
		return err
	}

	return product.UpdateProduct(ctx)
}

// Delete a record of product
func (p *Product) Delete(ctx context.Context, req *producttype.DeleteProductRequest, resp *producttype.DeleteProductResponse) error {
	var product model.Product

	if err := util.Bind(&product, req); err != nil {
		return err
	}

	return product.DeleteProduct(ctx)
}
