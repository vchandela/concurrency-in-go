package utils

import (
	"matmul/constants"
)

var (
	result = [constants.MatrixSize][constants.MatrixSize]int{}
)

func workoutRow(row int, matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	for col := 0; col < constants.MatrixSize; col++ {
		for k := 0; k < constants.MatrixSize; k++ {
			result[row][col] += (matrixA[row][k] * matrixB[k][col])
		}
	}
}

func MatMul(matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	for row := 0; row < constants.MatrixSize; row++ {
		workoutRow(row, matrixA, matrixB)
	}

	// for row := 0; row < constants.MatrixSize; row++ {
	// 	fmt.Println(result[row])
	// }
}
