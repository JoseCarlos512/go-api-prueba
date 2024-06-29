package main

import (
	"go-api-prueba/app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/rotate-matrix", handlers.RotateMatrixHandler)
	app.Post("/qr-factorization", handlers.QRFactorizationHandler)

	log.Fatal(app.Listen(":3100"))
}
