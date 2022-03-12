package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management-system/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id",controllers.GetUser())
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.LogIn())
	incomingRoutes.PATCH("/user/update/:user_id",controllers.UpdateUser())
}