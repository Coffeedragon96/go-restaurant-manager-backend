package routes

import (
	controller "github.com/Coffeedragon96/go-restaurant-manager-backend/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.AddTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}
