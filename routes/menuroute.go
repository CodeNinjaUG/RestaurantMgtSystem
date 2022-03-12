package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management-system/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus",controllers.GetMenus())
	incomingRoutes.GET("/menus/:id",controllers.GetMenu())
	incomingRoutes.POST("/menu/store",controllers.StoreMenu())
	incomingRoutes.DELETE("/menu/delete/:id",controllers.DeleteMenu())
	incomingRoutes.PATCH("/menu/update/:id",controllers.UpdateMenu())
}