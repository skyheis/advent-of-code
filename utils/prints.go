package utils

import (
	"fmt"
	"time"
)

func PrintResult(part string, result int, duration time.Duration) {
	fmt.Printf("Result for part %s: %d\n", part, result)
	fmt.Printf("Execution time: %s\n\n", duration)
}

func PrintResult64(part string, result int64, duration time.Duration) {
	fmt.Printf("Result for part %s: %d\n", part, result)
	fmt.Printf("Execution time: %s\n\n", duration)
}
