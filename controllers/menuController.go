package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)



func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}



func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}


func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}