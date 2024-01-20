package app

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
)

// RunServer runs the server.
func RunServer() {
	e := echo.New()
	Router(e)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
