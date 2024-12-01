/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var eighthCmd = &cobra.Command{
	Use:   "8th",
	Short: "Day 8 of Advent of Code 2023",
	Long: `Day 8 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/8`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/eight.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayEighthFirstPart(filecontent)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			eighthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result64 := dayEighthSecondPart(filecontent)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			eighthSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(eighthCmd)
}

func dayEighthGetRules(s string) [281]int {
	var rule [281]int

	for i, val := range s {
		if val == 'R' {
			rule[i] = 1
		}
	}
	return rule
}

func dayEighthGetTurns(inp []string, m map[string]int) [746][2]int {
	var turns [746][2]int

	for i, str := range inp {
		if i > 1 && str != "" {
			turns[i-2][0] = m[str[7:10]]
			turns[i-2][1] = m[str[12:15]]
		}
	}
	return turns
}

func dayEightItsOver(val int, b map[int]bool) bool {
	return b[val]
}

func dayEighthFirstPart(filecontent []byte) (result int) {
	inputFile := strings.Split(string(filecontent), "\n")

	//format the rules so are 0 and 1
	rule := dayEighthGetRules(inputFile[0])

	//fill map so each str is a num
	m := make(map[string]int)
	var init, over int
	for i, str := range inputFile {
		if i > 1 && str != "" {
			m[str[:3]] = i - 2
			if str[:3] == "AAA" {
				init = i - 2
			} else if str[:3] == "ZZZ" {
				over = i - 2
			}
		}
	}

	//convert rule in num as map
	turns := dayEighthGetTurns(inputFile, m)

	//init the calc
	result = 0
	for x := init; x != over; result++ {
		x = turns[x][rule[result%len(rule)]]
	}

	return
}

func dayEighthSecondPart(filecontent []byte) (result int64) {
	inputFile := strings.Split(string(filecontent), "\n")

	//format the rules so are 0 and 1
	rule := dayEighthGetRules(inputFile[0])

	//fill map so each str is valm
	m := make(map[string]int)
	var init [6]int
	over := make(map[int]bool)
	var countA, countZ int

	for i, str := range inputFile {
		if i > 1 && str != "" {
			m[str[:3]] = i - 2

			if str[2] == 'A' {
				init[countA] = i - 2
				countA++
			} else if str[2] == 'Z' {
				over[i-2] = true
				countZ++
			}
		}
	}

	//convert rule in num as map
	turns := dayEighthGetTurns(inputFile, m)

	//init the calc to find the first end for each init point
	var steps [6]int64
	for i, cur := range init {
		for !dayEightItsOver(cur, over) {
			cur = turns[cur][rule[steps[i]%int64(len(rule))]]
			steps[i]++
		}
	}

	//find the gcd (greater commod divsor) [mcd (massimo comune divisiore)]
	result = steps[0]
	for i := range steps {
		result = result / utils.Mcd(result, steps[i]) * steps[i]
	}

	return
}

func eighthSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
