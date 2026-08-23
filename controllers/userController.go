package controllers

import(
	"context"
	"fmt"
	"log"
	"time"
	"restio/database"
	"restio/models"
	helper "restio/helpers"

	"github.com/gin-gonic/gin"

)


func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	}
}

func GetUser() gin.HandlerFunc  {
	return func(ctx *gin.Context) {

	}
}


func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


func HashPassword(password string) string {

}


func VerifyPassword(userPassword string, providedPassword string) (bool, string)  {
	
}