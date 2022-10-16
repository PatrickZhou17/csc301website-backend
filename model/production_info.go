package model

import (
	"context"
	"gorm.io/gorm"

	"shopping-cart/common/orm"
	"shopping-cart/pkg/lormg"
	"strings"
)

type Product struct {
	orm.Model

	Name  string  `gorm:"column:name" json:"name"`   //name
	Price float32 `gorm:"column:price" json:"price"` //price
	Total float32 `gorm:"column:total" json:"total"` //amount
}

// TableName sets insert table name for this struct type
func (p *Product) TableName() string {
	return "product_info"
}

// GetProductList retrieves a list of product from database
func (p *Product) GetProductList(ctx context.Context, page, pageSize int) ([]Product, int64, error) {
	var (
		productList []Product
		count int64
		db = lormg.FromContext(ctx)
	)

	result := db.Where(p).Scopes(orm.Paginate(page, pageSize, true)).Find(&productList)
	if err := result.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	orm.DB.Model(&Product{}).Count(&count)

	return productList, count, nil
}

// GetProduct retrieves a single record of product from database
func (p *Product) GetProduct(ctx context.Context) (product Product, err error) {
	var db = lormg.FromContext(ctx)

	curErr := db.Where("id = ?", p.ID).First(&product).Error
	if curErr != nil && curErr != gorm.ErrRecordNotFound {
		err = curErr
	}

	return
}

// AddProduct persists product to database
func (p *Product) AddProduct(ctx context.Context) error {
	var db = lormg.FromContext(ctx)

	if err := db.Create(p).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := db.Model(&Product{}).Where("name = ?", p.Name).Updates(p).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

// UpdateProduct changes product by id
func (p *Product) UpdateProduct(ctx context.Context) error {
	var db = lormg.FromContext(ctx)

	if err := db.Model(&Product{}).Where("id = ?", p.ID).Updates(p).Error; err != nil {
		return err
	}

	return nil
}

// DeleteProduct product by id
func (p *Product) DeleteProduct(ctx context.Context) error {
	var db = lormg.FromContext(ctx)

	if err := db.Model(&Product{}).Unscoped().Where("id = ?", p.ID).Delete(p).Error; err != nil {
		return err
	}

	return nil
}
