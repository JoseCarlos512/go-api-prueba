package models

// Estructura matriz de entrada
type MatrixInput struct {
	Matrix [][]float64 `json:"matrix"`
}

// Estructura matriz rotada
type RotatedMatrixOutput struct {
	RotatedMatrix [][]float64 `json:"rotatedMatrix"`
}

// Estructura factorización QR
type QRDecompositionOutput struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// Estructura QRStatsResponse representa la respuesta de la API de Node
type QRStatsResponse struct {
	Q StatsResponse `json:"Q"`
	R StatsResponse `json:"R"`
}

// Estructura Estadisticas
type StatsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Mean       float64 `json:"mean"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}
