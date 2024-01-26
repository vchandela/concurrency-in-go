package utils

import (
	"matmul/constants"
	"sync"
)

var (
	resultRWMutex = [constants.MatrixSize][constants.MatrixSize]int{}
	wgRW          sync.WaitGroup
	lockRW        = sync.RWMutex{}
)

func workoutRowRWMutex(wgRW *sync.WaitGroup, row int, matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	defer wgRW.Done()

	lockRW.RLock()
	for col := 0; col < constants.MatrixSize; col++ {
		for k := 0; k < constants.MatrixSize; k++ {
			resultRWMutex[row][col] += (matrixA[row][k] * matrixB[k][col])
		}
	}
	lockRW.RUnlock()
}

func MatMulRWMutex(matrixA, matrixB [constants.MatrixSize][constants.MatrixSize]int) {
	wgRW.Add(constants.MatrixSize)
	for row := 0; row < constants.MatrixSize; row++ {
		go workoutRowRWMutex(&wgRW, row, matrixA, matrixB)
	}
	wgRW.Wait()

	// for row := 0; row < constants.MatrixSize; row++ {
	// 	fmt.Println(resultRWMutex[row])
	// }
}
