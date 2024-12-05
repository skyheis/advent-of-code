/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentythree

import (
	"advent-of-code/utils"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var tenthCmd = &cobra.Command{
	Use:   "10th",
	Short: "Day 10 of Advent of Code 2023",
	Long: `Day 10 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/10`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/ten.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayTenCommonPart(filecontent, false)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			tenthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = dayTenCommonPart(filecontent, true)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			tenthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(tenthCmd)
}

func dayTenCommonPart(filecontent []byte, secondPart bool) (result int) {
	inputFile := strings.Split(string(filecontent), "\n")
	inputRune := make([][]rune, len(inputFile))

	for i := range inputRune {
		inputRune[i] = []rune(inputFile[i])
	}

	startPoint, err := getCoordinates(inputFile)
	utils.CheckPanic(err)

	// replace all real pipes [2] and count steps[1]
	result = startRun(&inputRune, startPoint[0], startPoint[1]) / 2

	if !secondPart { // just part one c:
		return
	}

	// check on the right (good for this specific input) [2]
	startRun(&inputRune, startPoint[0], startPoint[1])

	for y, str := range inputRune {
		for x, char := range str {
			if char == 'i' {
				floodFill(&inputRune, x, y)
			}
		}
	}

	result = 0
	for _, str := range inputRune {
		for _, char := range str {
			if char == 'z' {
				result++
			}
		}
	}

	return
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

func tenthSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
