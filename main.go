package main

import (
	"log"
	"os"
	"shopping-cart/api/rest/probe"
	"shopping-cart/api/rest/product"
	"shopping-cart/common/config"
	"shopping-cart/common/orm"
	"shopping-cart/pkg/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	dbCfg := config.DatabaseConfig{
		ConnectUrls: config.DBConnUrls{
			Master: "uu23w23lguh6ef2q:mfoybymlxryqtpmw@tcp(r98du2bxwqkq3shg.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/	tz4sjutxf9pddu6t?charset=utf8mb4&parseTime=True&loc=Local",
		},
		MaxIdle:  10,
		MaxOpen:  100,
		LifeTime: 180,
	}
	orm.Setup(dbCfg, true)
	router.Use(middleware.Cors(), gin.Recovery())

	router.GET("/product/list", product.GetProductList)
	router.GET("/product/info", product.GetProduct)
	router.POST("/product/add", product.AddProduct)
	router.POST("/product/update", product.UpdateProduct)
	router.DELETE("/product/delete", product.DeleteProduct)

	router.GET("/health", probe.HealthCheck)
	router.GET("/ready", probe.ReadyCheck)
	router.Run(":" + port)
}
