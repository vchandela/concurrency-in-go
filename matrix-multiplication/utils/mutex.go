package utils

import (
	"matmul/constants"
	"sync"
)

var (
	resultMutex = [constants.MatrixSize][constants.MatrixSize]int{}
	wg          sync.WaitGroup
	lock        = sync.Mutex{}
)

func workoutRowMutex(wg *sync.WaitGroup, row int, matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	defer wg.Done()

	lock.Lock()
	for col := 0; col < constants.MatrixSize; col++ {
		for k := 0; k < constants.MatrixSize; k++ {
			resultMutex[row][col] += (matrixA[row][k] * matrixB[k][col])
		}
	}
	lock.Unlock()
}

func MatMulMutex(matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	wg.Add(constants.MatrixSize)
	for row := 0; row < constants.MatrixSize; row++ {
		go workoutRowMutex(&wg, row, matrixA, matrixB)
	}
	wg.Wait()

	// for row := 0; row < constants.MatrixSize; row++ {
	// 	fmt.Println(resultMutex[row])
	// }
}
