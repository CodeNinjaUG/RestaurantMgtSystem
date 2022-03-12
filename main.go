package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"restaurant-management-system/database"
	"restaurant-management-system/routes"
	"restaurant-management-system/middlewares"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")
func main()  {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.logger())
	router.Use(middlewares.Authentication())
	route.UserRoutes(router)
    router.Use(middlewares.Authentication)

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
  	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
