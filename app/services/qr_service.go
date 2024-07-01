package services

import (
	"encoding/json"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Ratate matriz 90 grades
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

// Realiza la factorización QR a la matriz
func QRFactorization(matrix [][]float64) (Q, R json.RawMessage, err error) {

	rows := len(matrix)
	cols := len(matrix[0])

	var m *mat.Dense
	if cols > rows {
		// Crear una nueva matriz cuadrada con filas adicionales de ceros
		m = mat.NewDense(cols, cols, nil)
		for i := range matrix {
			// Asegurar que cada fila tenga la longitud adecuada
			row := make([]float64, cols)
			copy(row, matrix[i])
			m.SetRow(i, row) // Establecer cada fila de la nueva matriz con los valores de la matriz original
		}
	} else {
		// Si hay más filas o las mismas que columnas, crear una matriz densa normal
		m = mat.NewDense(rows, cols, nil)
		for i := range matrix {
			m.SetRow(i, matrix[i]) // Establecer cada fila de la nueva matriz con los valores de la matriz original
		}
	}

	// Realizar la factorización QR
	var qr mat.QR
	qr.Factorize(m)

	var Qmat, Rmat *mat.Dense
	// Obtener las matrices Q y R
	if m.RawMatrix().Rows == m.RawMatrix().Cols {
		// matriz cuadrada
		Qmat = mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Cols, nil)
		Rmat = mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Cols, nil)
	} else if m.RawMatrix().Rows > m.RawMatrix().Cols {
		// matriz rectangular (caso no comprendido)
		Qmat = mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Rows, nil)
		Rmat = mat.NewDense(m.RawMatrix().Cols, m.RawMatrix().Cols, nil)
	} else {
		// matriz rectangular
		Qmat = mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Rows, nil)
		Rmat = mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Cols, nil)
	}

	qr.QTo(Qmat)
	qr.RTo(Rmat)

	// Recortar ceros
	if cols > rows {
		Qmat = Qmat.Slice(0, rows, 0, rows).(*mat.Dense)
		Rmat = Rmat.Slice(0, rows, 0, cols).(*mat.Dense)
	}

	// Convertir a JSON
	QBytes, err := matrixToJSON(Qmat)
	if err != nil {
		return nil, nil, err
	}

	// Convertir a JSON
	RBytes, err := matrixToJSON(Rmat)
	if err != nil {
		return nil, nil, err
	}

	return QBytes, RBytes, nil
}

// Conviertir una matriz a un JSON
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
		fmt.Printf("Error al parsesar a JSON: %v\n", err)
		return nil, err
	}
	return jsonData, nil
}
