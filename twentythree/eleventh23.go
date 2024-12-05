/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentythree

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var eleventhCmd = &cobra.Command{
	Use:   "11th",
	Short: "Day 11 of Advent of Code 2023",
	Long: `Day 11 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/11`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/eleven.input"

		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result := dayElevenCommonPart(bufio.NewScanner(file), 2)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			eleventhSubject(true)
		}
		utils.PrintResult("one", result, end)

		file.Seek(0, 0)

		start = time.Now()
		result = dayElevenCommonPart(bufio.NewScanner(file), 1000000)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			eleventhSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(eleventhCmd)
}

type gal struct {
	x, y int
}

func addEmptyLine(inputRune [][]rune) []int {
	var empty []int

	for y, line := range inputRune {
		flag := true
		for _, char := range line {
			if char == '#' {
				flag = false
			}
		}
		if flag {
			empty = append(empty, y)
		}
	}
	return empty
}

func addEmptyColumn(inputFile [][]rune) []int {
	var empty []int

	for x := 0; x < len(inputFile[0]); x++ {
		flag := true
		for y := 0; y < len(inputFile); y++ {
			if inputFile[y][x] == '#' {
				flag = false
			}
		}
		if flag {
			empty = append(empty, x)
		}
	}
	return empty
}

func findGalaxies(space [][]rune) []gal {
	var galaxies []gal
	for y, line := range space {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, gal{x, y})
			}
		}
	}
	return galaxies
}

func fillSpace(one int, two int, empty []int, much int) int {
	min := 0
	max := 0

	if one > two {
		min = two
		max = one
	} else {
		min = one
		max = two
	}
	spaces := 0
	for i := 0; i < max-min; i++ {
		if slices.Contains(empty, i+min) {
			// fmt.Println(i+min, "is in", empty)
			spaces += much - 1
		}
	}
	return spaces
}

func dayElevenCommonPart(scanner *bufio.Scanner, expand int) (result int) {

	var inputRune [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		inputRune = append(inputRune, []rune(line))
	}

	eLines := addEmptyLine(inputRune)
	eColumns := addEmptyColumn(inputRune)

	galaxies := findGalaxies(inputRune)

	result = 0
	for i, galax := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			toReach := galaxies[j]
			result += (utils.Abs(galax.x-toReach.x) + utils.Abs(galax.y-toReach.y))
			result += fillSpace(galax.x, toReach.x, eColumns, expand)
			result += fillSpace(galax.y, toReach.y, eLines, expand)
		}
	}

	return
}

// func dayNineCommonPart(filecontent []byte, secondPart bool) int64 {
// inputFile := strings.Split(string(filecontent), "\n")
// var result int64 = 0

func eleventhSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
