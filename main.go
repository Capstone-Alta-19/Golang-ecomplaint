package main

import (
	"capstone/config"
	"capstone/middleware"
	"capstone/route"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
