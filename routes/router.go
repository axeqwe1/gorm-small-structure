package routes

import (
	_ "coffee-shop/docs"

	"coffee-shop/internal/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // เปลี่ยนจาก swaggerFiles
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products", controller.GetProductsAllController)
	r.POST("/products", controller.CreateProductController)
	r.PUT("/products/:id", controller.UpdateProductController)
	r.DELETE("/products/:id", controller.DeleteProductController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
