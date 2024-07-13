package main

import (
	"fmt"
	"net/http"
	"the_basics_of_messenger_handler/utilities"
)

func main() {
	utilities.InitializeRoutes()

	fmt.Println("Running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
