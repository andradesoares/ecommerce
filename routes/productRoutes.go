package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("users/productView", controllers.ProductView())
	incomingRoutes.GET("users/search", controllers.SearchProduct())
	incomingRoutes.POST("users/addProduct", controllers.AddProduct())
}
