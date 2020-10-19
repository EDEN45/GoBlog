package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestAction() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, map[string]string{"ERROR": "404"})
	}
}
