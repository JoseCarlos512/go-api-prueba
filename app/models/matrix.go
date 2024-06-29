package models

// MatrixInput representa la estructura de la entrada
type MatrixInput struct {
	Matrix [][]float64 `json:"matrix"`
}

// RotatedMatrixOutput representa la estructura de la matriz rotada
type RotatedMatrixOutput struct {
	RotatedMatrix [][]float64 `json:"rotatedMatrix"`
}

// QRDecompositionOutput representa la estructura de la factorización QR
type QRDecompositionOutput struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

type QRStatsResponse struct {
	Q StatsResponse `json:"Q"`
	R StatsResponse `json:"R"`
}

// StatsResponse representa la respuesta de la API de Node.js
type StatsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Mean       float64 `json:"mean"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}
