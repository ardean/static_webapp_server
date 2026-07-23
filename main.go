package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func envOrDefault(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func isCacheEnabled() bool {
	value := strings.ToLower(strings.TrimSpace(envOrDefault("CACHE_ENABLED", "true")))
	return value != "false" && value != "0" && value != "off" && value != "no"
}

func main() {
	baseUrl := envOrDefault("BASE_URL", "/")
	folder := envOrDefault("FOLDER", "/public")
	cacheEnabled := isCacheEnabled()

	port := 8080
	if portString, exists := os.LookupEnv("PORT"); exists {
		port, _ = strconv.Atoi(portString)
	}

	app := fiber.New()
	app.Use(compress.New())
	app.Use(etag.New())

	if cacheEnabled {
		app.Use(cache.New())
	} else {
		app.Use(func(c *fiber.Ctx) error {
			c.Set("Cache-Control", "no-store, no-cache, must-revalidate")
			return c.Next()
		})
	}

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app.Static(baseUrl, folder)
	app.Static("*", folder+"/index.html")

	log.Printf("static webapp server listening on :%d (cache enabled: %t)", port, cacheEnabled)
	log.Fatal(app.Listen(":" + strconv.Itoa(port)))
}
