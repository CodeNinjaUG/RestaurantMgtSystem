package routes


import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management-system/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables",controllers.GetTables())
	incomingRoutes.GET("/tables/:id",controllers.GetTable())
	incomingRoutes.POST("/table/post",controllers.StoreTable())
	incomingRoutes.DELETE("/delete/table/:id",controllers.DeleteTable())
	incomingRoutes.PATCH("/update/table/:id",controllers.UpdateTable())
}
