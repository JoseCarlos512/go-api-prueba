package utils

import (
	"encoding/json"
	"go-api-prueba/app/models"

	"github.com/go-resty/resty/v2"
)

// SendRotatedMatrix envía la matriz rotada a la API de Node.js y recibe las estadísticas
func SendRotatedMatrix(rotatedMatrix [][]float64) (*models.StatsResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(models.RotatedMatrixOutput{RotatedMatrix: rotatedMatrix}).
		Post("http://localhost:3000/api/matrix")

	if err != nil {
		return nil, err
	}

	var stats models.StatsResponse
	err = json.Unmarshal(resp.Body(), &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
