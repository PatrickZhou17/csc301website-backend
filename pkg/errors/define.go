package errors

import (
	"shopping-cart/common/errors"
)

var (
	ErrorServerInternal   = errors.ServerError
	ErrorProductNameExist = errors.NewError(10001, "item already exists")
)
