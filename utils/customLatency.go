package utils

import (
	"fmt"
	"time"
)

func CustomLatency(start time.Time, end time.Time) string {
	latency := end.Sub(start)
	latencyMillis := float64(latency.Nanoseconds()) / 1e6 
	return fmt.Sprintf("%.2fms", latencyMillis) 
}     