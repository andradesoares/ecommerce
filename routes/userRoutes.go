package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/signin", controllers.Signin())
	incomingRoutes.GET("users/productView", controllers.ProductView())
	incomingRoutes.GET("users/search", controllers.SearchProduct())
	incomingRoutes.POST("users/addProduct", controllers.AddProduct())
}
