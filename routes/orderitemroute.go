package routes

import (
"github.com/gin-gonic/gin"
controllers "restaurant-management-system/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderitems",controllers.GetOrderItems())
	incomingRoutes.GET("/orderitem/:id",controllers.GetOrderItem())
	incomingRoutes.GET("/orderitem/:order_id",controllers.GetOrderItemByOrderID())
	incomingRoutes.POST("/orderitem/store",controllers.StoreOrderItem())
	incomingRoutes.DELETE("/orderitem/delete/:id",controllers.DeleteOrderItem())
	incomingRoutes.PATCH("/orderitem/update/:id",controllers.UpdateOrderItem())
}

git remote add my_awesome_new_remote_repo git@git.assembla.com:portfolio/space.space_name.git