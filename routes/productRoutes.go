package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("products/productView", controllers.ProductView())
	incomingRoutes.GET("products/search", controllers.SearchProduct())
	incomingRoutes.POST("products/addProduct", controllers.AddProduct())
}
