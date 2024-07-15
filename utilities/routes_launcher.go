package utilities

import (
	"github.com/gin-gonic/gin"

	"the_basics_of_messenger_handler/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

func InitializeRoutes() *gin.Engine {
	r := gin.Default()

	routes := []Route{
		{
			Path:    "/users",
			Method:  "GET",
			Handler: handlers.SendMessage,
		},
	}

	for _, route := range routes {
		if route.Method == "GET" {
			r.GET(route.Path, route.Handler)
		}
		if route.Method == "POST" {
			r.POST(route.Path, route.Handler)
		}
	}

	return r
}
