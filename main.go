package main

import (
	"os"
	"restio/database"
	middleware "restio/middleware"
	routes "restio/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())
	
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)

	router.Run(":" + port)
}