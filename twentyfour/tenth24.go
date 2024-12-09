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
var tenthCmd = &cobra.Command{
	Use:   "10th",
	Short: "Day 10 of Advent of Code 2024",
	Long: `Day 10 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/10`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/ten.input"
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
	TwentyfourCmd.AddCommand(tenthCmd)
}

func dayTenCommonPart(filecontent []byte, secondPart bool) (result int) {
	// inputFile := strings.Split(string(filecontent), "\n")
	// inputRune := make([][]rune, len(inputFile))

	result = 1
	return
}

func tenthSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
