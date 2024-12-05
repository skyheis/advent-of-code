/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentythree

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var thirdCmd = &cobra.Command{
	Use:   "3rd",
	Short: "Day 3 of Advent of Code 2023",
	Long: `Day 3 of Advent of Code 2023
Here the problem statement: https://adventofcode.com/2023/day/3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/three.input"

		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		// file, error := os.Open(inputPath)
		// utils.CheckInputFileError(error)
		// defer file.Close()

		start := time.Now()
		file_mat := strings.Split(string(filecontent), "\n")
		result := dayThreeFirstPart(file_mat)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(true)
		}
		utils.PrintResult("one", result, end)

		// file.Seek(0, 0)

		start = time.Now()
		file_mat = strings.Split(string(filecontent), "\n")
		result = dayThreeSecondPart(file_mat)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(thirdCmd)
}

func dayThreeReplaceChar(file_mat *[]string, x int, y int) {
	s := (*file_mat)[y]
	r := []rune(s)
	r[x] = '.'
	s = string(r)
	(*file_mat)[y] = s
}

func dayThreeFirstPartCheckSymbol(file_mat *[]string, symbolFlag *bool, x int, y int) {
	if *symbolFlag || x < 0 || y < 0 || x >= len((*file_mat)[y]) || y >= len(*file_mat) {
		return
	} else if (*file_mat)[y][x] == '.' {
		return
	} else if unicode.IsDigit(rune((*file_mat)[y][x])) {
		dayThreeReplaceChar(file_mat, x, y)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x+1, y)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x-1, y-1)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x, y-1)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x+1, y-1)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x-1, y)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x-1, y+1)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x, y+1)
		dayThreeFirstPartCheckSymbol(file_mat, symbolFlag, x+1, y+1)
	} else {
		*symbolFlag = true
	}
}

func dayThreeFirstPart(file_mat []string) int {
	result := 0

	var width int
	height := len(file_mat)

	if height > 0 {
		width = len(file_mat[0])
	} else {
		panic("File is empty")
	}

	for y := 0; y < height-1; y++ {
		for x := 0; x < width; x++ {
			s := 0
			for x+s < width && unicode.IsDigit(rune(file_mat[y][x+s])) {
				s++
			}

			if s > 0 {
				num, err := strconv.Atoi(file_mat[y][x : x+s])
				utils.CheckPanic(err)
				symbolFlag := false
				dayThreeFirstPartCheckSymbol(&file_mat, &symbolFlag, x, y)
				if symbolFlag {
					result += num
				}
				x += s
			}
		}
	}

	return result
}

func dayThreeSecondPartCheckSymbol(file_mat *[]string, symbolFlag *bool, num *int, x int, y int) {
	if *symbolFlag || x < 0 || y < 0 || x >= len((*file_mat)[y]) || y >= len(*file_mat) {
		return
	} else if (*file_mat)[y][x] == '.' {
		return
	} else if unicode.IsDigit(rune((*file_mat)[y][x])) {

		r := 0
		for x+r < len((*file_mat)[y]) && unicode.IsDigit(rune((*file_mat)[y][x+r])) {
			r++
		}
		l := 0
		for x-l >= 0 && unicode.IsDigit(rune((*file_mat)[y][x-l])) {
			l++
		}
		atoi, err := strconv.Atoi((*file_mat)[y][x-l+1 : x+r])
		for i := x - l + 1; i < x+r; i++ {
			dayThreeReplaceChar(file_mat, i, y)
		}
		utils.CheckPanic(err)
		if *num == 0 {
			*num = atoi
		} else {
			*num *= atoi
			*symbolFlag = true
		}

	} else if (*file_mat)[y][x] == '*' {
		dayThreeReplaceChar(file_mat, x, y)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x+1, y)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x-1, y-1)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x, y-1)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x+1, y-1)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x-1, y)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x-1, y+1)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x, y+1)
		dayThreeSecondPartCheckSymbol(file_mat, symbolFlag, num, x+1, y+1)
	}
}

func dayThreeSecondPart(file_mat []string) int {
	result := 0

	var width int
	height := len(file_mat)

	if height > 0 {
		width = len(file_mat[0])
	} else {
		panic("File is empty")
	}

	for y := 0; y < height-1; y++ {
		for x := 0; x < width; x++ {
			if file_mat[y][x] == '*' {
				num := 0
				symbolFlag := false
				dayThreeSecondPartCheckSymbol(&file_mat, &symbolFlag, &num, x, y)
				if symbolFlag {
					result += num
				}
			}
		}
	}

	return result
}

func thirdSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?`)
	} else {
		fmt.Println(`--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?`)
	}
	fmt.Println()
}
