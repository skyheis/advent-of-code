/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentyfour

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var fifthCmd = &cobra.Command{
	Use:   "5th",
	Short: "Day 5 of Advent of Code 2024",
	Long: `Day 5 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/5`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		// inputPath := "twentyfour/inputs/five.input.test"
		// filecontent, err := os.ReadFile(inputPath)
		// utils.CheckInputFileError(err)

		// start := time.Now()
		// result := dayFiveCommonPart(string(filecontent), dayFourFirstPart)
		// end := time.Since(start)

		if cmd.Flag("subject").Changed {
			fifthSubject(true)
		}
		// utils.PrintResult("one", result, end)

		// start = time.Now()
		// result = dayFourCommonPart(string(filecontent), dayFourSecondPart)
		// end = time.Since(start)

		if cmd.Flag("subject").Changed {
			fifthSubject(false)
		}
		// utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(fifthCmd)
}

// func dayFourCommonPart(filecontent string, part func([][]rune, int, int) int) (result int) {
// 	input := utils.MakeRuneMatrixStr(filecontent)

func fifthSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}

//
