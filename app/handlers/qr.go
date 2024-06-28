package handlers

import (
	"go-api-prueba/app/models"
	"go-api-prueba/app/services"
	"go-api-prueba/app/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// QRFactorizationHandler maneja la factorización QR
func QRFactorizationHandler(c *fiber.Ctx) error {
	var input models.MatrixInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Realizar rotación de la matriz
	rotatedMatrix := services.RotateMatrix(input.Matrix)

	// Enviar matriz rotada a la API de Node.js
	stats, err := utils.SendRotatedMatrix(rotatedMatrix)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Realizar factorización QR
	Q, R, err := services.QRFactorization(input.Matrix)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	result := fiber.Map{
		"Q":     Q,
		"R":     R,
		"stats": stats,
	}

	return c.JSON(result)
}
