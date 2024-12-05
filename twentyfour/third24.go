/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var thirdCmd = &cobra.Command{
	Use:   "3rd",
	Short: "Day 3 of Advent of Code 2024",
	Long: `Day 3 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/three.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayThreeFirstPart(string(filecontent))
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = dayThreeSecondPart(string(filecontent))
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(thirdCmd)
}

func dayThreeSecondPart(filecontent string) (result int) {
	inputFile := strings.Split(string(filecontent), "do()")
	var new []string
	for _, val := range inputFile {
		new = append(new, strings.Split(string(val), "don'")...)
	}

	var removeDont []string
	for _, value := range new {
		if strings.Compare(value[:3], "t()") == 0 {
			continue
		}
		removeDont = append(removeDont, value)
	}

	for _, val := range removeDont {
		result += dayThreeFirstPart(val)
	}

	return
}

func dayThreeFirstPart(filecontent string) (result int) {
	inputFile := strings.Split(string(filecontent), "mul(")

	for _, value := range inputFile {

		i := utils.CountDigits(value)
		x, err := strconv.Atoi(value[:i])
		if err != nil || x > 999 {
			continue
		}

		if i < len(value) && value[i] == ',' {
			i++
		} else {
			continue
		}

		k := utils.CountDigits(value[i:])
		y, err := strconv.Atoi(value[i : i+k])
		if err != nil || y > 999 {
			continue
		}

		if value[i+k] != ')' {
			continue
		}

		result += x * y
	}

	return result
}

func thirdSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 3: Mull It Over ---
"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

	xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?`)
	} else {
		fmt.Println(`--- Part Two ---
As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

 - The do() instruction enables future mul instructions.
 - The don't() instruction disables future mul instructions.

Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

	xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))

This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions are disabled because there is a don't() instruction before them. The other mul instructions function normally, including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?`)
	}
	fmt.Println()
}
