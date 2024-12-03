/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"fmt"
	"os"

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
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		// start := time.Now()
		// result := dayThreeFirstPart(bufio.NewScanner(file))
		// end := time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(true)
		}
		// utils.PrintResult("one", result, end)

		file.Seek(0, 0)

		// start = time.Now()
		// result = dayThreeSecondPart(bufio.NewScanner(file))
		// end = time.Since(start)

		if cmd.Flag("subject").Changed {
			thirdSubject(false)
		}
		// utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(thirdCmd)
}

func thirdSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}
