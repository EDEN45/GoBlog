package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGetAction() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"hello": "OK"})
	}
}
