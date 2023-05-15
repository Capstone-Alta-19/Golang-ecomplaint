package main

import (
	"fmt"
	"os"
	"project_structure/config"
	"project_structure/middleware"
	"project_structure/route"

	"github.com/labstack/echo"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	port := os.Getenv("PORT")
	const DEFAULT_PORT = "8080"

	if port == "" {
		port = DEFAULT_PORT
	}
	appPort := fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(appPort))
}
