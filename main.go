package main

import (
	"os"
	"the_basics_of_messenger_handler/emulator"
	"the_basics_of_messenger_handler/handlers"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func main() {
	utilities.DotEnvHandler()
	r := initializeRoutes()
	go r.Run(":8080")
	if os.Getenv("EMULATOR") == "true" {
		go emulator.InitializeEmulator()
	}
	select {}
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
		{
			Path:    "/v1/receive",
			Method:  "POST",
			Handler: handlers.ReceiveMessage,
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
