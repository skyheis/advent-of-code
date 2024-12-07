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
var seventhCmd = &cobra.Command{
	Use:   "7th",
	Short: "Day 7 of Advent of Code 2024",
	Long: `Day 7 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/7`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/seven.input.test"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := daySevenCommonPart(string(filecontent), true)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			seventhSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = daySevenCommonPart(string(filecontent), false)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			seventhSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(seventhCmd)
}

func daySevenCommonPart(filecontent string, first bool) (result int) {

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
