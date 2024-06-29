package utils

import (
	"encoding/json"
	"fmt"
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

// SendQRDecomposition envía las matrices Q y R a la API de Node.js y recibe las estadísticas
func SendQRDecomposition(Q, R json.RawMessage) (*models.QRStatsResponse, error) {

	var QMatrix, RMatrix [][]float64

	// Deserializar Q
	if err := json.Unmarshal(Q, &QMatrix); err != nil {
		return nil, err
	}

	// Deserializar R
	if err := json.Unmarshal(R, &RMatrix); err != nil {
		return nil, err
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(models.QRDecompositionOutput{Q: QMatrix, R: RMatrix}).
		Post("http://localhost:3000/api/qr")

	if err != nil {
		return nil, err
	}

	var stats models.QRStatsResponse

	fmt.Printf("Estadisticas: %+v\n", stats)
	err = json.Unmarshal(resp.Body(), &stats)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Estadisticas: %+v\n", stats)

	return &stats, nil
}
