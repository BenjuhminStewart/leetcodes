package helpers

import (
	"time"
)

func GetAverageTime(times []time.Duration) time.Duration {
	sum := 0.0
	for _, v := range times {
		sum += float64(v.Microseconds())
	}
	return time.Duration(sum/float64(len(times))) * 1000
}
