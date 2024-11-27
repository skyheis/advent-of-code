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

func pipeFromTop(chart []string, x int, y int, steps *int) {
	// fmt.Printf("char %c ", chart[y][x])
	// fmt.Println("pos:", x, y)
	checkInMap(x, y)
	if chart[y][x] != 'S' {
		(*steps)++
		if chart[y][x] == '|' {
			pipeFromTop(chart, x, y+1, steps)
		} else if chart[y][x] == 'J' {
			pipeFromRight(chart, x-1, y, steps)
		} else if chart[y][x] == 'L' {
			pipeFromLeft(chart, x+1, y, steps)
		}
	}
}

func pipeFromBottom(chart []string, x int, y int, steps *int) {
	// fmt.Printf("char %c ", chart[y][x])
	// fmt.Println("pos:", x, y)
	checkInMap(x, y)
	if chart[y][x] != 'S' {
		(*steps)++
		if chart[y][x] == '|' {
			pipeFromBottom(chart, x, y-1, steps)
		} else if chart[y][x] == 'F' {
			pipeFromLeft(chart, x+1, y, steps)
		} else {
			pipeFromRight(chart, x-1, y, steps)
		}
	}
}

func pipeFromLeft(chart []string, x int, y int, steps *int) {
	// fmt.Printf("char %c ", chart[y][x])
	// fmt.Println("pos:", x, y)
	checkInMap(x, y)
	if chart[y][x] != 'S' {
		(*steps)++
		if chart[y][x] == '-' {
			pipeFromLeft(chart, x+1, y, steps)
		} else if chart[y][x] == 'J' {
			pipeFromBottom(chart, x, y-1, steps)
		} else {
			pipeFromTop(chart, x, y+1, steps)
		}
	}
}

func pipeFromRight(chart []string, x int, y int, steps *int) {
	// fmt.Printf("char %c ", chart[y][x])
	// fmt.Println("pos:", x, y)
	checkInMap(x, y)
	if chart[y][x] != 'S' {
		(*steps)++
		if chart[y][x] == '-' {
			pipeFromRight(chart, x-1, y, steps)
		} else if chart[y][x] == 'L' {
			pipeFromBottom(chart, x, y-1, steps)
		} else {
			pipeFromTop(chart, x, y+1, steps)
		}
	}
}

func startRun(chart []string, x int, y int) int {
	fmt.Println("position:", x, y)
	var steps int = 0
	if x-1 >= 0 && (chart[y][x-1] == '-' || chart[y][x-1] == 'F' || chart[y][x-1] == 'L') {
		steps++
		pipeFromRight(chart, x-1, y, &steps)
	} else if y-1 >= 0 && (chart[y-1][x] == '|' || chart[y-1][x] == 'F' || chart[y-1][x] == '7') {
		steps++
		pipeFromBottom(chart, x, y-1, &steps)
	} else if x+1 <= 139 && (chart[y][x+1] == '-' || chart[y][x+1] == 'J' || chart[y][x+1] == '7') {
		steps++
		pipeFromLeft(chart, x+1, y, &steps)
	} else if y+1 <= 139 && (chart[y+1][x] == '|' || chart[y+1][x] == 'J' || chart[y+1][x] == 'L') {
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

	startPoint, err := getCoordinates(inputFile)
	check(err)

	result := startRun(inputFile, startPoint[0], startPoint[1]) / 2

	end := time.Since(start)

	fmt.Println("The result is: ", result, "for")
	fmt.Println("ExecTime is:", end)
}
