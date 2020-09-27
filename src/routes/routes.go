package routes

import (
	"github.com/EDEN45/GoBlog/src/controllers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method string
	Path   string
	Action gin.HandlerFunc
}

func GetRoutes() []Route {
	return []Route{
		{Method: "GET", Path: "/", Action: controllers.TestGetAction()},
		{Method: "GET", Path: "/oper", Action: controllers.OperGetAction()},
	}
}
