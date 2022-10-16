package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"shopping-cart/common/errors"
	"shopping-cart/common/log"
)

func BindReqAndValid(c *gin.Context, form interface{}) error {
	ginLog := log.GetFromGin(c)
	if err := c.ShouldBind(form); err != nil {
		ginLog.Errorf("bind req err: %v", err)
		return errors.NewValidateError(err)
	}

	if err := c.ShouldBindHeader(form); err != nil {
		ginLog.Errorf("bind header err: %v", err)
		return err
	}

	xvalid := Validation{}
	ok, err := xvalid.Check(form, xvalidFun...)
	if !ok {
		ginLog.Errorf("valid req err: %v", err)
		return errors.NewValidateError(err)
	}

	err = Validate(form)
	if err != nil {
		ginLog.Errorf("valid req err: %v", err)
		return errors.NewValidateError(err)
	}

	return nil
}

func NoSpaceValidate(fl validator.FieldLevel) bool {
	result := false
	result = !strings.HasPrefix(fl.Field().String(), " ")
	if !result {
		return result
	}
	result = !strings.HasSuffix(fl.Field().String(), " ")
	if !result {
		return result
	}
	result = !strings.HasPrefix(fl.Field().String(), "\t")
	if !result {
		return result
	}
	result = !strings.HasSuffix(fl.Field().String(), "\t")
	if !result {
		return result
	}
	return result
}

func Validate(obj interface{}) error {
	valid := validator.New()
	if err := valid.RegisterValidation("nospace", NoSpaceValidate); err != nil {
		return err
	}
	err := valid.Struct(obj)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		errStrList := make([]string, 0, len(errs))
		for _, e := range errs {
			et := fmt.Sprintf("valid req err '%s' failed must be %s %s", e.StructField(), e.ActualTag(), e.Param())
			errStrList = append(errStrList, et)
		}
		errStr := strings.Join(errStrList, "\n")
		rerr := fmt.Errorf("%s", errStr)
		return rerr
	}
	return nil
}
