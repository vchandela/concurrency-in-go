package helpers

import (
	"math/rand"
	"matmul/constants"
)

func GenerateRandomMatrix(matrix *[constants.MatrixSize][constants.MatrixSize]int) {
	for row := 0; row < constants.MatrixSize; row++ {
		for col := 0; col < constants.MatrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5 //[-5, -5]
		}
	}
}
