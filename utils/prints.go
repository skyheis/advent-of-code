package utils

import (
	"fmt"
	"os"
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

func WriteFileFromRuneMatrix(labMap [][]rune, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range labMap {
		for _, spot := range line {
			_, err := file.WriteString(string(spot))
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func PrintRuneMatrix(mat [][]rune) {
	for _, line := range mat {
		for _, spot := range line {
			fmt.Printf("%c", spot)
		}
		fmt.Println()
	}
}
