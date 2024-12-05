/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var fourthCmd = &cobra.Command{
	Use:   "4th",
	Short: "Day 4 of Advent of Code 2024",
	Long: `Day 4 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/four.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayFourCommonPart(string(filecontent), dayFourFirstPart)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			fourthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = dayFourCommonPart(string(filecontent), dayFourSecondPart)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			fourthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(fourthCmd)
}

func findXmas(input [][]rune, x, y, dx, dy, counter int) int {
	xmas := "XMAS"

	x += dx
	y += dy

	if !utils.InRange(input, x, y) {
		return 0
	}

	if input[y][x] == rune(xmas[counter]) {
		if counter == 3 {
			return 1
		}
		return findXmas(input, x, y, dx, dy, counter+1)
	}

	return 0
}

func dayFourFirstPart(input [][]rune, x, y int) (parResult int) {
	parResult = 0

	if input[y][x] == 'X' {
		parResult += findXmas(input, x, y, -1, -1, 1)
		parResult += findXmas(input, x, y, 0, -1, 1)
		parResult += findXmas(input, x, y, 1, -1, 1)
		parResult += findXmas(input, x, y, -1, 0, 1)
		parResult += findXmas(input, x, y, 1, 0, 1)
		parResult += findXmas(input, x, y, -1, 1, 1)
		parResult += findXmas(input, x, y, 0, 1, 1)
		parResult += findXmas(input, x, y, 1, 1, 1)
	}

	return
}

func findMasAsX(input [][]rune, x, y int) bool {
	twice := 0

	if !utils.InRange(input, x-1, y-1) || !utils.InRange(input, x+1, y+1) {
		return false
	}

	if (input[y-1][x-1] == 'M' && input[y+1][x+1] == 'S') ||
		(input[y-1][x-1] == 'S' && input[y+1][x+1] == 'M') {
		twice += 1
	}

	if (input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S') ||
		(input[y-1][x+1] == 'S' && input[y+1][x-1] == 'M') {
		twice += 1
	}

	return twice == 2
}

func dayFourSecondPart(input [][]rune, x, y int) (parResult int) {

	if input[y][x] == 'A' && findMasAsX(input, x, y) {
		return 1
	}

	return 0
}

func dayFourCommonPart(filecontent string, part func([][]rune, int, int) int) (result int) {
	input := utils.MakeRuneMatrixStr(filecontent)

	result = 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			result += part(input, x, y)

		}
	}

	return
}

func fourthSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 4: Ceres Search ---
"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:

	..X...
	.SAMX.
	.A..A.
	XMAS.S
	.X....

The actual word search will be full of letters instead. For example:

	MMMSXXMASM
	MSAMXMSMSA
	AMXSXMAAMM
	MSAMASMSMX
	XMASAMXAMM
	XXAMMXXAMA
	SMSMSASXSS
	SAXAMASAAA
	MAMMMXMMMM
	MXMXAXMASX

In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

	....XXMAS.
	.SAMXMS...
	...S..A...
	..A.A.MS.X
	XMASAMX.MM
	X.....XA.A
	S.S.S.S.SS
	.A.A.A.A.A
	..M.M.M.MM
	.X.X.XMASX

Take a look at the little Elf's word search. How many times does XMAS appear?`)
	} else {
		fmt.Println(`--- Part Two ---
The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

	M.S
	.A.
	M.S

Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

	.M.S......
	..A..MSMS.
	.M.S.MAA..
	..A.ASMSM.
	.M.S.M....
	..........
	S.S.S.S.S.
	.A.A.A.A..
	M.M.M.M.M.
	..........

In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?`)
	}
	fmt.Println()
}

//
