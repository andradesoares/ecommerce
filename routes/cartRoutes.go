package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("cart/addToCart", controllers.AddToCart())
	incomingRoutes.GET("cart/removeItem", controllers.RemoveItem())
	incomingRoutes.GET("cart/cartCheckout", controllers.CartCheckout())
	incomingRoutes.GET("cart/instantBuy", controllers.InstantBuy())
}
