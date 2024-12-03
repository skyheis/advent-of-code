/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var firstCmd = &cobra.Command{
	Use:   "1st",
	Short: "Day 1 of Advent of Code 2023",
	Long: `Day 1 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/one.input"
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result := dayOneFirstPart(bufio.NewScanner(file))
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			firstSubject(true)
		}
		utils.PrintResult("one", result, end)

		file.Seek(0, 0)

		start = time.Now()
		result = dayOneSecondPart(bufio.NewScanner(file))
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			firstSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(firstCmd)
}

func digitOrNumber(line string, numbs [9]string) int {

	for i := 0; i < len(numbs); i++ {
		if strings.HasPrefix(line, numbs[i]) {
			return i + 1
		}
	}
	if unicode.IsDigit(rune(line[0])) {
		return int(line[0]) - 48
	}
	return 0
}

func dayOneSecondPart(scanner *bufio.Scanner) int {
	result := 0
	numbs := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		value := 0
		flag := false
		line := scanner.Text()

		for i := 0; !flag && i < len(line); i++ {
			digit := digitOrNumber(line[i:], numbs)
			if digit > 0 {
				value = digit * 10
				flag = true
			}
		}
		for i := len(line) - 1; flag && i >= 0; i-- {
			digit := digitOrNumber(line[i:], numbs)
			if digit > 0 {
				value += digit
				flag = false
			}
		}
		result += value
	}

	return result
}

func dayOneFirstPart(scanner *bufio.Scanner) int {
	result := 0

	for scanner.Scan() {
		value := 0
		flag := false
		line := scanner.Text()

		for i := 0; !flag && i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				value = (int(line[i]) - 48) * 10
				flag = true
			}
		}
		for i := len(line) - 1; flag && i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				value += int(line[i]) - 48
				flag = false
			}
		}
		result += value
	}

	return result
}

func firstSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?`)
	} else {
		fmt.Println(`--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

	two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?`)
	}
	fmt.Println()
}
