package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCoordinates(input []string) ([2]int, error) {
	for y, str := range input {
		for x, char := range str {
			if char == 'S' {
				return [2]int{x, y}, nil
			}
		}
	}
	return [2]int{0, 0}, errors.New("s is not present")
}

func checkInMap(x int, y int) {
	if y < 0 || x < 0 || x > 139 || y > 139 {
		panic(errors.New("out of border, input error"))
	}
}

func pipeFromTop(chart *[][]rune, x int, y int, steps *int) {
	checkInMap(x, y)
	if (*chart)[y][x] != 's' {
		(*steps)++
		if (*chart)[y][x] == '|' || (*chart)[y][x] == 'b' {
			if (*chart)[y][x] == 'b' {
				replaceRight(chart, x-1, y)
			} else {
				(*chart)[y][x] = 'b'
			}
			pipeFromTop(chart, x, y+1, steps)
		} else if (*chart)[y][x] == 'J' || (*chart)[y][x] == 'c' {
			(*chart)[y][x] = 'c'
			pipeFromRight(chart, x-1, y, steps)
		} else {
			if (*chart)[y][x] == 'f' {
				replaceRight(chart, x-1, y)
				replaceRight(chart, x, y+1)
			} else {
				(*chart)[y][x] = 'f'
			}
			pipeFromLeft(chart, x+1, y, steps)
		}
	}
}

func pipeFromBottom(chart *[][]rune, x int, y int, steps *int) {
	checkInMap(x, y)
	if (*chart)[y][x] != 's' {
		(*steps)++
		if (*chart)[y][x] == '|' || (*chart)[y][x] == 'b' {
			if (*chart)[y][x] == 'b' {
				replaceRight(chart, x+1, y)
			} else {
				(*chart)[y][x] = 'b'
			}
			pipeFromBottom(chart, x, y-1, steps)
		} else if (*chart)[y][x] == 'F' || (*chart)[y][x] == 'e' {
			(*chart)[y][x] = 'e'
			pipeFromLeft(chart, x+1, y, steps)
		} else {
			if (*chart)[y][x] == 'd' {
				replaceRight(chart, x+1, y)
				replaceRight(chart, x, y+1)
			} else {
				(*chart)[y][x] = 'd'
			}
			pipeFromRight(chart, x-1, y, steps)
		}
	}
}

func pipeFromLeft(chart *[][]rune, x int, y int, steps *int) {
	checkInMap(x, y)
	if (*chart)[y][x] != 's' {
		(*steps)++
		if (*chart)[y][x] == '-' || (*chart)[y][x] == 'a' {
			if (*chart)[y][x] == 'a' {
				replaceRight(chart, x, y+1)
			} else {
				(*chart)[y][x] = 'a'
			}
			pipeFromLeft(chart, x+1, y, steps)
		} else if (*chart)[y][x] == 'J' || (*chart)[y][x] == 'c' {
			if (*chart)[y][x] == 'c' {
				replaceRight(chart, x+1, y)
				replaceRight(chart, x, y+1)
			} else {
				(*chart)[y][x] = 'c'
			}
			pipeFromBottom(chart, x, y-1, steps)
		} else {
			(*chart)[y][x] = 'd'
			pipeFromTop(chart, x, y+1, steps)
		}
	}
}

func pipeFromRight(chart *[][]rune, x int, y int, steps *int) {
	checkInMap(x, y)
	if (*chart)[y][x] != 's' {
		(*steps)++
		if (*chart)[y][x] == '-' || (*chart)[y][x] == 'a' {
			if (*chart)[y][x] == 'a' {
				replaceRight(chart, x-1, y-1)
			} else {
				(*chart)[y][x] = 'a'
			}
			pipeFromRight(chart, x-1, y, steps)
		} else if (*chart)[y][x] == 'L' || (*chart)[y][x] == 'f' {
			(*chart)[y][x] = 'f'
			pipeFromBottom(chart, x, y-1, steps)
		} else {
			if (*chart)[y][x] == 'e' {
				replaceRight(chart, x-1, y)
				replaceRight(chart, x, y+1)
			} else {
				(*chart)[y][x] = 'e'
			}
			pipeFromTop(chart, x, y+1, steps)
		}
	}
}

func replaceRight(chart *[][]rune, x int, y int) {

	if y < 0 || x < 0 || x > 139 || y > 139 {
		return
	} else if (*chart)[y][x] != 's' && ((*chart)[y][x] < 'a' || (*chart)[y][x] > 'f') {
		(*chart)[y][x] = 'i'
	}
}

func floodFill(chart *[][]rune, x int, y int) {
	if y < 0 || x < 0 || x > 139 || y > 139 {
		return
	} else if (*chart)[y][x] != 's' && (*chart)[y][x] != 'z' && ((*chart)[y][x] < 'a' || (*chart)[y][x] > 'f') {
		(*chart)[y][x] = 'z'
		floodFill(chart, x+1, y)
		floodFill(chart, x-1, y)
		floodFill(chart, x, y+1)
		floodFill(chart, x, y-1)
	}
}

func startRun(chart *[][]rune, x int, y int) int {
	var steps int = 0

	(*chart)[y][x] = 's'
	if x-1 >= 0 && ((*chart)[y][x-1] == '-' || (*chart)[y][x-1] == 'F' || (*chart)[y][x-1] == 'L' || (*chart)[y][x-1] == 'a' || (*chart)[y][x-1] == 'e' || (*chart)[y][x-1] == 'f') {
		steps++
		pipeFromRight(chart, x-1, y, &steps)
	} else if y-1 >= 0 && ((*chart)[y-1][x] == '|' || (*chart)[y-1][x] == 'F' || (*chart)[y-1][x] == '7' || (*chart)[y-1][x] == 'b' || (*chart)[y-1][x] == 'e' || (*chart)[y-1][x] == 'd') {
		steps++
		pipeFromBottom(chart, x, y-1, &steps)
	} else if x+1 <= 139 && ((*chart)[y][x+1] == '-' || (*chart)[y][x+1] == 'J' || (*chart)[y][x+1] == '7' || (*chart)[y][x+1] == 'a' || (*chart)[y][x+1] == 'c' || (*chart)[y][x+1] == 'd') {
		steps++
		pipeFromLeft(chart, x+1, y, &steps)
	} else if y+1 <= 139 && ((*chart)[y+1][x] == '|' || (*chart)[y+1][x] == 'J' || (*chart)[y+1][x] == 'L' || (*chart)[y+1][x] == 'b' || (*chart)[y+1][x] == 'c' || (*chart)[y+1][x] == 'f') {
		steps++
		pipeFromTop(chart, x, y+1, &steps)
	}
	return steps
}

func main() {
	fContent, err := os.ReadFile("10th_day/input")
	check(err)

	start := time.Now()

	inputFile := strings.Split(string(fContent), "\n")
	inputRune := make([][]rune, len(inputFile))
	for i := range inputRune {
		inputRune[i] = []rune(inputFile[i])
	}

	startPoint, err := getCoordinates(inputFile)
	check(err)

	// replace all real pipes [2] and count steps[1]
	resultFirst := startRun(&inputRune, startPoint[0], startPoint[1]) / 2

	endFirst := time.Since(start)

	// check on the right (good for this specific input) [2]
	startRun(&inputRune, startPoint[0], startPoint[1])

	for y, str := range inputRune {
		for x, char := range str {
			if char == 'i' {
				floodFill(&inputRune, x, y)
			}
		}
	}

	resultSecond := 0
	for _, str := range inputRune {
		for _, char := range str {
			if char == 'z' {
				resultSecond++
			}
		}
	}

	endSecond := time.Since(start)

	fmt.Println("First part's result: ", resultFirst)
	fmt.Println("ExecTime first part: ", endFirst)
	fmt.Println("Second part's result:", resultSecond)
	fmt.Println("ExecTime second part:", endSecond)
}
