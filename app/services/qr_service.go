package services

import (
	"encoding/json"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// RotateMatrix rota la matriz 90 grados en sentido horario
func RotateMatrix(matrix [][]float64) [][]float64 {
	rowCount := len(matrix)
	colCount := len(matrix[0])
	rotated := make([][]float64, colCount)
	for i := range rotated {
		rotated[i] = make([]float64, rowCount)
	}
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			rotated[j][rowCount-1-i] = matrix[i][j]
		}
	}
	return rotated
}

// QRFactorization realiza la factorización QR de la matriz
func QRFactorization(matrix [][]float64) (Q, R json.RawMessage, err error) {
	m := mat.NewDense(len(matrix), len(matrix[0]), nil)
	for i := range matrix {
		m.SetRow(i, matrix[i])
	}

	var qr mat.QR
	qr.Factorize(m)

	rows, cols := m.Dims()

	Qmat := mat.NewDense(rows, rows, nil)
	Rmat := mat.NewDense(rows, cols, nil)
	qr.QTo(Qmat)
	qr.RTo(Rmat)

	QBytes, err := matrixToJSON(Qmat)
	if err != nil {
		return nil, nil, err
	}

	RBytes, err := matrixToJSON(Rmat)
	if err != nil {
		return nil, nil, err
	}

	return QBytes, RBytes, nil
}

// matrixToJSON convierte una matriz a un JSON
func matrixToJSON(m mat.Matrix) (json.RawMessage, error) {
	rows, cols := m.Dims()
	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			data[i][j] = m.At(i, j)
		}
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling matrix to JSON: %v\n", err)
		return nil, err
	}
	return jsonData, nil
}
