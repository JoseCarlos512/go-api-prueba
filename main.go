// main.go

package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "go-api-prueba/configs"
    "go-api-prueba/app/handlers"
)

func main() {
    // Cargar configuración
    config := configs.LoadConfig()

    // Inicializar aplicación Fiber
    app := fiber.New()

    // Rutas
    app.Post("/rotate-matrix", handlers.RotateMatrixHandler)

    // Escuchar en el puerto especificado en la configuración
    log.Printf("Servidor iniciado en el puerto %s", config.Port)
    log.Fatal(app.Listen(":" + config.Port))
}
