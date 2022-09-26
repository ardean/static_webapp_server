package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	folder := os.Getenv("FOLDER")
	if len(folder) == 0 {
		folder = "./public"
	}

	port := 8080

	portString := os.Getenv("PORT")
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	app := fiber.New()

	app.Static("/", folder)
	app.Static("*", folder+"/index.html")

	log.Fatal(app.Listen(":" + strconv.Itoa(port)))
}
