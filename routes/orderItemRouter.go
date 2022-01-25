package routes

import (
	controller "github.com/Coffeedragon96/go-restaurant-manager-backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.AddOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
