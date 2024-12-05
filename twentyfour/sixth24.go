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
var sixthCmd = &cobra.Command{
	Use:   "6th",
	Short: "Day 6 of Advent of Code 2024",
	Long: `Day 6 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/6`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/six.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := daySixCommonPart(string(filecontent), daySixFirstPart)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			sixthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = daySixCommonPart(string(filecontent), daySixSecondPart)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			sixthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(sixthCmd)
}

func daySixFirstPart(input [][]rune, x, y int) (parResult int) {
	parResult = 0

	return
}

func daySixSecondPart(input [][]rune, x, y int) (parResult int) {

	return 0
}

func daySixCommonPart(filecontent string, part func([][]rune, int, int) int) (result int) {

	return
}

func sixthSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}

//
