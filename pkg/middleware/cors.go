package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// to allows resources such as web fonts to be loaded from domains other than the origin parent domain
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	return cors.New(config)
}
