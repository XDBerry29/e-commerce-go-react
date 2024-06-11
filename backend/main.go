package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading the env...")
	}

	PORT := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + PORT))
	fmt.Print("hello")
}
