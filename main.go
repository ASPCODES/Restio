package main

import (
	"log"
	"os"
	"restio/database"
	middleware "restio/middleware"
	routes "restio/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not set in .env file")
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