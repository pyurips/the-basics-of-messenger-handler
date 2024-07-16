package main

import (
	"the_basics_of_messenger_handler/handlers"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func main() {
	utilities.DotEnvHandler()
	r := initializeRoutes()
	r.Run(":8080")
}

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

func initializeRoutes() *gin.Engine {
	r := gin.Default()

	routes := []Route{
		{
			Path:    "/v1/message",
			Method:  "POST",
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
