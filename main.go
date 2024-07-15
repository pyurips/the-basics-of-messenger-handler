package main

import (
	"the_basics_of_messenger_handler/configs"
)

func main() {
	r := configs.InitializeRoutes()
	r.Run(":8080")
}
