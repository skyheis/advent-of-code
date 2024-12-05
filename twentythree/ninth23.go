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

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var ninthCmd = &cobra.Command{
	Use:   "9th",
	Short: "Day 9 of Advent of Code 2023",
	Long: `Day 9 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/9`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/nine.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayNineCommonPart(filecontent, false)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			ninthSubject(true)
		}
		utils.PrintResult64("one", result, end)

		start = time.Now()
		result64 := dayNineCommonPart(filecontent, true)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			ninthSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(ninthCmd)
}

func oasisConvertion(oasisStr string) []int {
	var e error

	valStr := strings.Fields(oasisStr)
	valNum := make([]int, len(valStr))
	for i, str := range valStr {
		valNum[i], e = strconv.Atoi(str)
		utils.CheckPanic(e)
	}
	return valNum
}

func oasisPrediction(oasis []int, second bool) int64 {

	if second {
		for i, j := 0, len(oasis)-1; i < j; i, j = i+1, j-1 {
			oasis[i], oasis[j] = oasis[j], oasis[i]
		}
	}

	last := len(oasis) - 1
	for x := 0; x != last+1; x++ {
		i := last
		for ; i > x; i-- {
			oasis[i] = oasis[i] - oasis[i-1]
		}
	}
	oasis = append(oasis, 0)
	x := last + 1
	for i := 0; i != x; i++ {
		x := last + 1
		for ; x > i; x-- {
			oasis[x] = oasis[x] + oasis[x-1]
		}
	}

	return int64(oasis[last+1])
}

func dayNineCommonPart(filecontent []byte, secondPart bool) int64 {
	var result int64 = 0
	var count int32 = 0

	inputFile := strings.Split(string(filecontent), "\n")

	for _, oasisStr := range inputFile {
		if oasisStr != "" {
			oasis := oasisConvertion(oasisStr)
			result += oasisPrediction(oasis, secondPart)
			count++
		}
	}

	return result
}

func ninthSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
