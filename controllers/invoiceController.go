package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restio/database"
	"restio/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id 			string
	Payment_method      string
	Order_id			string
	Payment_status		*string
	Payment_due			interface{}
	Table_number		interface{}
	Payment_due_date	time.Time
	Order_details		interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100* time.Second)
		defer cancel()

		result, err := invoiceCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing invoice items"})
			return 
		}

		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while decoding invoice items"})
			return 
		}

		if allInvoices == nil {
			allInvoices = []bson.M{}
		}

		c.JSON(http.StatusOK, allInvoices)
	}
}


func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}


func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}


func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}


