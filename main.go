package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := StudentApiRouter()

	fmt.Println("Server listening on port 3000 ....")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		logError("Server error", err)
	}

}
