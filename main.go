package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func main() {
	baseUrl, exists := os.LookupEnv("BASE_URL")
	if !exists {
		baseUrl = "/"
	}

	folder, exists := os.LookupEnv("FOLDER")
	if !exists {
		folder = "/public"
	}

	port := 8080
	portString, exists := os.LookupEnv("PORT")
	if !exists {
		port, _ = strconv.Atoi(portString)
	}

	app := fiber.New()
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(cache.New())
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app.Static(baseUrl, folder)
	app.Static("*", folder+"/index.html")

	log.Fatal(app.Listen(":" + strconv.Itoa(port)))
}
