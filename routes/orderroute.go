package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management-system/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders",controllers.GetOrders())
	incomingRoutes.GET("/order/:id",controllers.GetOrder())
	incomingRoutes.POST("/order/store",controllers.StoreOrder())
	incomingRoutes.DELETE("/order/delete/:id",controllers.DeleteOrder())
	incomingRoutes.PATCH("/order/update/:id",controllers.UpdateOrder())
}