package models

// MatrixInput representa la estructura de la entrada
type MatrixInput struct {
	Matrix [][]float64 `json:"matrix"`
}

// RotatedMatrixOutput representa la estructura de la matriz rotada
type RotatedMatrixOutput struct {
	RotatedMatrix [][]float64 `json:"rotatedMatrix"`
}

// StatsResponse representa la respuesta de la API de Node.js
type StatsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Mean       float64 `json:"mean"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}
