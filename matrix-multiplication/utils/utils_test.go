package utils

import (
	"matmul/constants"
	"matmul/helpers"
	"testing"
)

var (
	matrixA = [constants.MatrixSize][constants.MatrixSize]int{}
	matrixB = [constants.MatrixSize][constants.MatrixSize]int{}
)

func BenchmarkMatMul(b *testing.B) {
	helpers.GenerateRandomMatrix(&matrixA)
	helpers.GenerateRandomMatrix(&matrixB)
	MatMul(matrixA, matrixB)
}

func BenchmarkMatMulMutex(b *testing.B) {
	helpers.GenerateRandomMatrix(&matrixA)
	helpers.GenerateRandomMatrix(&matrixB)
	MatMulMutex(matrixA, matrixB)
}

func BenchmarkMatMulRWMutex(b *testing.B) {
	helpers.GenerateRandomMatrix(&matrixA)
	helpers.GenerateRandomMatrix(&matrixB)
	MatMulRWMutex(matrixA, matrixB)
}
