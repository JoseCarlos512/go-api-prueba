// app/handlers/matrix_handler.go

package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go-api-prueba/app/services"
)

// RotateMatrixHandler maneja la rotación de la matriz 90 grados
func RotateMatrixHandler(c *fiber.Ctx) error {
    var requestBody struct {
        Matrix [][]int `json:"matrix"`
    }

    // Parsear la matriz desde el cuerpo JSON de la solicitud
    if err := c.BodyParser(&requestBody); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    // Aplicar la rotación de la matriz
    rotated := services.RotateMatrix90Degrees(requestBody.Matrix)

    // Devolver la matriz rotada como JSON
    return c.JSON(rotated)
}
