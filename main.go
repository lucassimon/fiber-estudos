package main

import (
	todos "fiber-estudos/src/todos"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println(os.Getenv("POSTGRES_URL"))

}

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}

func main() {
	app := fiber.New()

	// Default config
	app.Use(favicon.New())

	// Default middleware config
	app.Use(requestid.New())

	// Custom Timer middleware
	app.Use(Timer())

	// Middleware
	app.Use(recover.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	if os.Getenv("APP_ENV") == "dev" {
		// Or extend your config for customization
		app.Use(logger.New(logger.Config{
			Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "America/Sao_Paulo",
			Output:     os.Stdout,
		}))
	}

	// Default compression config
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "api",
		})
		return c.Next()
	}) // /api

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "v1",
		})
		return c.Next()
	}) // /api/v1

	// v2 := api.Group("/v2", func(c *fiber.Ctx) error {
	// 	c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"message": "v2",
	// 	})
	// 	return c.Next()
	// }) // /api/v2

	todos.TodosRouterV1(v1)
	// todos.TodosRouterV2(v2)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Resource not found.",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
