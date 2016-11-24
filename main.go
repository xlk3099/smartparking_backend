package main

import (
	"github.com/xlk3099/smartparking_backend/web"
)

func main() {

	// Create a new server using that MongoDB session
	server := web.NewServer()

	// Begin listening for HTTP requests
	server.Run(":8080")
}
