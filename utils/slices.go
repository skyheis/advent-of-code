package utils

import (
	"slices"
	"strconv"
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

func CopyRuneMatrix(m [][]rune) [][]rune {
	d := make([][]rune, len(m))

	for i, l := range m {
		d[i] = make([]rune, len(l))
		copy(d[i], l)
	}

	return d
}

func InRange(input [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(input) && x < len(input[y])
}

func SliceAtoi(s []string) (i []int) {
	for _, str := range s {
		atoi, err := strconv.Atoi(str)
		CheckPanic(err)
		i = append(i, atoi)
	}

	return
}

func SlicesMoveFront(slice []int, where, who int) (new []int) {
	move := slice[who]
	slice = slices.Delete(slice, who, who+1)

	duplicate := make([]int, len(slice))
	copy(duplicate, slice)

	new = append(duplicate[:where], move)
	new = append(new, slice[where:]...)

	return
}

func CountInMatrix(m [][]rune, r rune) (c int) {
	for _, line := range m {
		for _, val := range line {
			if val == r {
				c++
			}
		}
	}

	return
}

func CountOccurrences(slice []int, target int) (count int) {
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return
}

func InMatrix(mat [][]rune, spot Coordinates) bool {
	return spot.Y >= 0 && spot.Y < len(mat) && spot.X >= 0 && spot.X < len(mat[0])
}
