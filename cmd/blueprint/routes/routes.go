package routes

import (
	"github.com/EDEN45/GoBlog/cmd/blueprint/controllers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method string
	Path   string
	Action gin.HandlerFunc
}

func Run() {
	router := gin.Default()
	for _, val := range getRoutes() {
		switch val.Method {
		case "GET":
			router.GET(val.Path, val.Action)
		case "POST":
			router.POST(val.Path, val.Action)
		}
	}
	router.Run()
}

func getRoutes() []Route {
	return []Route{
		{Method: "GET", Path: "/", Action: controllers.TestGetAction()},
		{Method: "GET", Path: "/oper", Action: controllers.OperGetAction()},
	}
}
