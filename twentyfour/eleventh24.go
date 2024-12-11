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
var eleventhCmd = &cobra.Command{
	Use:   "11th",
	Short: "Day 11 of Advent of Code 2024",
	Long: `Day 11 of Advent of Code 2024, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/11`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/eleven.input"

		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayElevenFirstPart(filecontent)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			eleventhSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result64 := dayElevenSecondPart(filecontent)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			eleventhSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(eleventhCmd)
}

func splitEvenStone(stone int) (stone1, stone2 int) {
	digits := utils.CountDigitsInt(stone)

	stone1 = stone / utils.TenTimes(digits/2)
	stone2 = stone % utils.TenTimes(digits/2)

	return
}

func addStones(stones map[int]int, key, value int) {
	if _, present := stones[key]; !present {
		stones[key] = 0
	}
	stones[key] += value
}

func blinkMapStone(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			addStones(newStones, 1, count)
		} else if utils.CountDigitsInt(stone)%2 == 0 {
			stone1, stone2 := splitEvenStone(stone)

			addStones(newStones, stone1, count)
			addStones(newStones, stone2, count)
		} else {
			addStones(newStones, stone*2024, count)
		}

	}

	return newStones
}

func dayElevenSecondPart(filecontent []byte) (result int64) {
	inputStr := strings.Fields(string(filecontent))

	stoneMap := make(map[int]int)

	for _, stoneStr := range inputStr {
		stone, err := strconv.Atoi(stoneStr)
		utils.CheckError(err)

		stoneMap[stone] = 1
	}

	for i := 0; i < 75; i++ {
		stoneMap = blinkMapStone(stoneMap)
	}

	for _, value := range stoneMap {
		result += int64(value)
	}

	return
}

func blinkSplitStone(stone int, blink int) (result int) {

	if blink == 0 {
		return 1
	}

	if stone == 0 {
		result += blinkSplitStone(1, blink-1)
	} else if utils.CountDigitsInt(stone)%2 == 0 {
		stone1, stone2 := splitEvenStone(stone)

		result += blinkSplitStone(stone1, blink-1)
		result += blinkSplitStone(stone2, blink-1)
	} else {
		result += blinkSplitStone(stone*2024, blink-1)
	}

	return
}

func dayElevenFirstPart(filecontent []byte) (result int) {
	inputStr := strings.Fields(string(filecontent))

	for _, stoneStr := range inputStr {

		stone, err := strconv.Atoi(stoneStr)
		utils.CheckError(err)

		result += blinkSplitStone(stone, 25)
	}

	return
}

// func dayNineCommonPart(filecontent []byte, secondPart bool) int {
// inputFile := strings.Split(string(filecontent), "\n")
// var result int = 0

func eleventhSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}
