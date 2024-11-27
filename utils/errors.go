package utils

import (
	"fmt"
	"os"
	"time"
)

func CheckInputFileError(e error) {
	if e != nil {
		fmt.Println("Sorry, input file invalid:", e)
		os.Exit(1)
	}
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func CheckPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintResult(part string, result int, duration time.Duration) {
	fmt.Printf("Result for part %s: %d\n", part, result)
	fmt.Printf("Execution time: %s\n\n", duration)
}
