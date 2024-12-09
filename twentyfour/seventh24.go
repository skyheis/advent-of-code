/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var seventhCmd = &cobra.Command{
	Use:   "7th",
	Short: "Day 7 of Advent of Code 2024",
	Long: `Day 7 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/7`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/seven.input"
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result := daySevenCommonPart(bufio.NewScanner(file), false)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			seventhSubject(true)
		}
		utils.PrintResult64("one", result, end)

		file.Seek(0, 0)

		start = time.Now()
		result = daySevenCommonPart(bufio.NewScanner(file), true)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			seventhSubject(false)
		}
		utils.PrintResult64("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(seventhCmd)
}

func daySevenConvertInputLine(scannerTest string) (result int64, operators []int64) {

	splitted := strings.Split(scannerTest, ": ")

	result, err := strconv.ParseInt(splitted[0], 10, 64)
	utils.CheckError(err)

	operatorStr := strings.Fields(splitted[1])

	for _, val := range operatorStr {
		valInt, err := strconv.ParseInt(val, 10, 64)
		utils.CheckError(err)
		operators = append(operators, valInt)
	}

	return
}

func combineInt64(a, b int64) int64 {
	aStr := strconv.FormatInt(a, 10)
	bStr := strconv.FormatInt(b, 10)

	combinedStr := aStr + bStr

	combined, err := strconv.ParseInt(combinedStr, 10, 64)
	utils.CheckError(err)

	return combined
}

func tryOperCombo(numbers []int64, operators []rune) int64 {
	result := numbers[0]
	for i, oper := range operators {
		switch oper {
		case '+':
			result += numbers[i+1]
		case '*':
			result *= numbers[i+1]
		case '|':
			result = combineInt64(result, numbers[i+1])

		default:
			panic(fmt.Sprintln("no operator case:", oper))
		}
	}
	return result
}

func generateAllCases(n int, ops []rune) [][]rune {
	if n == 0 {
		return [][]rune{{}}
	}
	subOperators := generateAllCases(n-1, ops)
	var operators [][]rune
	for _, sub := range subOperators {
		for _, op := range ops {
			operators = append(operators, append([]rune{op}, sub...))
		}
	}
	return operators
}

func daySevenIsValidResult(result int64, numbers []int64, ops []rune) bool {
	operNumb := len(numbers) - 1
	operCases := generateAllCases(operNumb, ops)
	for _, operators := range operCases {
		if tryOperCombo(numbers, operators) == result {
			return true
		}
	}
	return false
}

func daySevenCommonPart(scanner *bufio.Scanner, two bool) (result int64) {

	ops := []rune{'+', '*'}
	if two {
		ops = append(ops, '|')
	}

	for scanner.Scan() {
		calculation, operators := daySevenConvertInputLine(scanner.Text())
		if daySevenIsValidResult(calculation, operators, ops) {
			result += calculation
		}
	}

	return
}

func seventhSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}

//
