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
		fmt.Println(`--- Day 11: Plutonian Pebbles ---
The ancient civilization on Pluto was known for its ability to manipulate spacetime, and while The Historians explore their infinite corridors, you've noticed a strange set of physics-defying stones.

At first glance, they seem like normal stones: they're arranged in a perfectly straight line, and each stone has a number engraved on it.

The strange part is that every time you blink, the stones change.

Sometimes, the number engraved on a stone changes. Other times, a stone might split in two, causing all the other stones to shift over a bit to make room in their perfectly straight line.

As you observe them for a while, you find that the stones have a consistent behavior. Every time you blink, the stones each simultaneously change according to the first applicable rule in this list:

 - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
 - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
 - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

No matter how the stones change, their order is preserved, and they stay on their perfectly straight line.

How will the stones evolve if you keep blinking at them? You take a note of the number engraved on each stone in the line (your puzzle input).

If you have an arrangement of five stones engraved with the numbers 0 1 10 99 999 and you blink once, the stones transform as follows:

 - The first stone, 0, becomes a stone marked 1.
 - The second stone, 1, is multiplied by 2024 to become 2024.
 - The third stone, 10, is split into a stone marked 1 followed by a stone marked 0.
 - The fourth stone, 99, is split into two stones marked 9.
 - The fifth stone, 999, is replaced by a stone marked 2021976.

So, after blinking once, your five stones would become an arrangement of seven stones engraved with the numbers 1 2024 1 0 9 9 2021976.

Here is a longer example:

	Initial arrangement:
	125 17

	After 1 blink:
	253000 1 7

	After 2 blinks:
	253 0 2024 14168

	After 3 blinks:
	512072 1 20 24 28676032

	After 4 blinks:
	512 72 2024 2 0 2 4 2867 6032

	After 5 blinks:
	1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

	After 6 blinks:
	2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2

In this example, after blinking six times, you would have 22 stones. After blinking 25 times, you would have 55312 stones!

Consider the arrangement of stones in front of you. How many stones will you have after blinking 25 times?`)
	} else {
		fmt.Println(`--- Part Two ---
The Historians sure are taking a long time. To be fair, the infinite corridors are very large.

How many stones would you have after blinking a total of 75 times?`)
	}
	fmt.Println()
}
