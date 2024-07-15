package main

import (
	"the_basics_of_messenger_handler/utilities"
)

func main() {
	r := utilities.InitializeRoutes()
	r.Run(":8080")
}
