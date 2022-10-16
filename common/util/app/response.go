package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"shopping-cart/common/errors"
)

const (
	defaultErrCode = 50000
	defaultErrMsg  = "server internal error"
)

type Resp struct {
	Status    int         `json:"status"`
	Content   interface{} `json:"content"`
	ErrorMsg  string      `json:"errorMsg"`
	Timestamp interface{} `json:"timestamp"`
}

func Success(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, http.StatusOK, data)
	return
}

func Error(c *gin.Context, httpCode int, err error) {
	if ce, ok := err.(errors.CustomError); ok {
		errCode := ce.Code()
		errMsg := ce.Error()
		Response(c, httpCode, errCode, errMsg)
		return
	}

	Response(c, httpCode, defaultErrCode, defaultErrMsg)
	return
}

func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	var errMsg string
	if errCode != http.StatusOK {
		errMsg = data.(string)
		data = nil
	}
	cur := time.Now().Unix()
	c.JSON(httpCode, Resp{
		Status:    errCode,
		Content:   data,
		ErrorMsg:  errMsg,
		Timestamp: cur,
	})
	return
}
