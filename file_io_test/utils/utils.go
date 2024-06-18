package utils

import (
	"file_io_test/constants"
	"math"
	"time"
)

var currentProcessConc = 1
var currentProcessConcLastUpdateTS = time.Now()

func init() { // Implement slow-start
	go func() {
		for currentProcessConc < constants.ProcessConc {
			timer := time.NewTimer(constants.ProcessConcSlowStartDoublingTime)
			<-timer.C
			currentProcessConc *= 2
			currentProcessConcLastUpdateTS = time.Now()
		}
		currentProcessConc = constants.ProcessConc
	}()
}

// GetCurrentProcessConc returns current concurrency; implementing slow-start with smoothening
func GetCurrentProcessConc() int {
	// No need for smoothening if already reached max-val
	if currentProcessConc == constants.ProcessConc {
		return currentProcessConc
	}
	// Implement smoothening; val = x*(1-lambda) + (2x)*lambda
	// This ensures we have a steady linear ramp-up from the current value(x) to the next value(2*x)
	nextProcessConc := math.Min(float64(constants.ProcessConc), float64(2*currentProcessConc))
	lambda := float64(time.Now().UnixMilli()-currentProcessConcLastUpdateTS.UnixMilli()) / float64(constants.ProcessConcSlowStartDoublingTime.Milliseconds())
	smoothenedProcessConc := int(float64(currentProcessConc)*(1-lambda) + nextProcessConc*lambda)
	return smoothenedProcessConc
}
