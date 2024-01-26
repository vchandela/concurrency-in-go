package main

import (
	"fmt"
	"matmul/constants"
	"matmul/helpers"
	"matmul/utils"
	"time"
)

var (
	matrixA = [constants.MatrixSize][constants.MatrixSize]int{}
	matrixB = [constants.MatrixSize][constants.MatrixSize]int{}
)

func main() {
	helpers.GenerateRandomMatrix(&matrixA)
	helpers.GenerateRandomMatrix(&matrixB)

	start := time.Now()
	utils.MatMul(matrixA, matrixB)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for classic algo: ", elapsed)

	start = time.Now()
	utils.MatMulMutex(matrixA, matrixB)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time using Mutex: ", elapsed)

	start = time.Now()
	utils.MatMulRWMutex(matrixA, matrixB)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time using RWMutex: ", elapsed)
}
