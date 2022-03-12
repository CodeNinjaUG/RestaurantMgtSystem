package routes

import (
	"github.com/gin-gonic/gin"
   controllers "restaurant-management-system/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods",controllers.GetFoods())
	incomingRoutes.GET("/foods/:id",controllers.GetFood())
	incomingRoutes.POST("/post/food",controllers.StoreFood())
	incomingRoutes.DELETE("/delete/food/:id",controllers.DeleteFood())
	incomingRoutes.PATCH("/food/:id",controllers.UpdateFood())
}
