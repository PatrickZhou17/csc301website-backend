package product

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart/common/server"
	appUtil "shopping-cart/common/util/app"

	"shopping-cart/api/typespec/producttype"
	productApp "shopping-cart/application/product"
	"shopping-cart/common/log"
)

// AddProduct
// @Summary add item
// @Router /product/add [POST]
func AddProduct(c *gin.Context) {
	var (
		req    producttype.AddProductRequest
		resp   producttype.AddProductResponse
		ginLog = log.GetFromGin(c)
	)

	err := appUtil.BindReqAndValid(c, &req)
	if err != nil {
		appUtil.Error(c, http.StatusOK, err)
		return
	}
	ginLog.Debug("GetProductList")

	app := &productApp.Product{}
	ctx := server.NewContext(context.Background(), c)
	if err := app.AddProduct(ctx, &req, &resp); err != nil {
		ginLog.Errorf("call app.AddProduct trigger err: %v", err)
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	appUtil.Success(c, resp)
}

// GetProductList
// @Summary get item from the list
// @Router /product/list [GET]
func GetProductList(c *gin.Context) {
	var (
		req    producttype.GetProductListRequest
		resp   producttype.GetProductListResponse
		ginLog = log.GetFromGin(c)
	)

	err := appUtil.BindReqAndValid(c, &req)
	if err != nil {
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	ginLog.Debug("GetProductList")

	app := productApp.Product{}
	ctx := server.NewContext(context.Background(), c)
	if err := app.GetProductList(ctx, &req, &resp); err != nil {
		ginLog.Errorf("call app.GetProductList trigger err: %v", err)
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	appUtil.Success(c, resp)
}

// GetProduct
// @Summary get the name of product
// @Router /product/info [GET]
func GetProduct(c *gin.Context) {
	var (
		req    producttype.GetProductRequest
		resp   producttype.GetProductResponse
		ginLog = log.GetFromGin(c)
	)

	err := appUtil.BindReqAndValid(c, &req)
	if err != nil {
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	ginLog.Debug("GetProducts")

	app := productApp.Product{}
	ctx := server.NewContext(context.Background(), c)
	if err := app.GetProduct(ctx, &req, &resp); err != nil {
		ginLog.Errorf("call app.GetProduct trigger err: %v", err)
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	appUtil.Success(c, resp)
}

// UpdateProduct
// @Summary update item
// @Router /product/update [POST]
func UpdateProduct(c *gin.Context) {
	var (
		req    producttype.UpdateProductRequest
		resp   producttype.UpdateProductResponse
		ginLog = log.GetFromGin(c)
	)

	err := appUtil.BindReqAndValid(c, &req)
	if err != nil {
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	ginLog.Debug("UpdateProduct")

	app := productApp.Product{}
	ctx := server.NewContext(context.Background(), c)
	if err := app.UpdateProduct(ctx, &req, &resp); err != nil {
		ginLog.Errorf("call app.UpdateProduct trigger err: %v", err)
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	appUtil.Success(c, resp)
}

// DeleteProduct
// @Summary delete item
// @Router /product/delete [DELETE]
func DeleteProduct(c *gin.Context) {
	var (
		req    producttype.DeleteProductRequest
		resp   producttype.DeleteProductResponse
		ginLog = log.GetFromGin(c)
	)

	err := appUtil.BindReqAndValid(c, &req)
	if err != nil {
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	ginLog.Debug("DeleteProduct")

	app := productApp.Product{}
	ctx := server.NewContext(context.Background(), c)
	if err := app.DeleteProduct(ctx, &req, &resp); err != nil {
		ginLog.Errorf("call app.DeleteProduct trigger err: %v", err)
		appUtil.Error(c, http.StatusOK, err)
		return
	}

	appUtil.Success(c, resp)
}
