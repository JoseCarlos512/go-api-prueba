package services

import (
    "fmt"
)

// RotateMatrix90Degrees rota una matriz 90 grados en sentido antihorario
func RotateMatrix90Degrees(matrix [][]int) [][]int {
    rows := len(matrix)
    cols := len(matrix[0])

    rotated := make([][]int, cols)
    for i := range rotated {
        rotated[i] = make([]int, rows)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            rotated[cols-j-1][i] = matrix[i][j]
        }
    }

    // Imprimir la matriz rotada para verificar
    fmt.Println("Matriz rotada en servicio:")
    printMatrix(rotated)

    return rotated
}

// Función de utilidad para imprimir una matriz
func printMatrix(matrix [][]int) {
    for _, row := range matrix {
        fmt.Println(row)
    }
    fmt.Println()
}
