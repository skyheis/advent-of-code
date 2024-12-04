package utils

import (
	"strings"
	"unicode"
)

func RemoveIndex(slice []int, index int) (new []int) {
	for i, v := range slice {
		if i == index {
			continue
		}
		new = append(new, v)
	}
	return
}

func CountDigits(value string) (i int) {
	if len(value) == 0 {
		return
	}
	for unicode.IsDigit(rune(value[i])) {
		i++
	}
	return
}

func MakeRuneMatrixByte(filecontent []byte) [][]rune {
	inputFile := strings.Split(string(filecontent), "\n")
	inputRune := make([][]rune, len(inputFile))

	for i := range inputRune {
		inputRune[i] = []rune(inputFile[i])
	}

	return inputRune
}

func MakeRuneMatrixStr(filecontent string) [][]rune {
	inputFile := strings.Split(filecontent, "\n")
	inputRune := make([][]rune, len(inputFile))

	for i := range inputRune {
		inputRune[i] = []rune(inputFile[i])
	}

	return inputRune
}

func InRange(input [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(input) && x < len(input[y])
}
