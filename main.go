package main

import (
	"the_basics_of_messenger_handler/configs"
	"the_basics_of_messenger_handler/utilities"
)

func main() {
	utilities.DotEnvHandler()
	r := configs.InitializeRoutes()
	r.Run(":8080")
}
