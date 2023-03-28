package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/addToCart", controllers.AddToCart())
	incomingRoutes.GET("/removeItem", controllers.RemoveItem())
	incomingRoutes.GET("/cartCheckout", controllers.CartCheckout())
	incomingRoutes.GET("/instantBuy", controllers.InstantBuy())
}
