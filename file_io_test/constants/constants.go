package constants

import "time"

const (
	MAX_POLL_RECORDS                 = 1024
	NUM_WORKERS                      = 10
	ProcessConc                      = 20
	ProcessConcSlowStartDoublingTime = 5 * time.Minute
)
