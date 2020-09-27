package kernel

import (
	"github.com/EDEN45/GoBlog/src/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	for _, val := range routes.GetRoutes() {
		switch val.Method {
		case "GET":
			router.GET(val.Path, val.Action)
		case "POST":
			router.POST(val.Path, val.Action)
		}
	}
	router.Run()
}
