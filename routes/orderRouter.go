package routes

import (
	controller "github.com/Coffeedragon96/go-restaurant-manager-backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.AddOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
