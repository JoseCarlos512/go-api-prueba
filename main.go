package main

import (
	"go-api-prueba/app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/qr-factorization", handlers.QRFactorizationHandler)

	log.Fatal(app.Listen(":3100"))
}
