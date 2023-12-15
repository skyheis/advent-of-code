package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

type gal struct {
	x, y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func ftAbs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}

func fillSpace(one int, two int, empty []int, much int) int {
	min := 0
	max := 0
	// fmt.Println("fill", one, two)
	// fmt.Println(empty)
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

func main() {
	fContent, err := os.Open("11th_day/input")
	check(err)

	start := time.Now()

	scanner := bufio.NewScanner(fContent)

	var inputRune [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		inputRune = append(inputRune, []rune(line))
	}

	eLines := addEmptyLine(inputRune)
	eColumns := addEmptyColumn(inputRune)

	galaxies := findGalaxies(inputRune)

	totalDist := 0

	commonSinceStart := time.Since(start)
	endCommon := time.Now()

	for i, galax := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			toReach := galaxies[j]
			totalDist += (ftAbs(galax.x-toReach.x) + ftAbs(galax.y-toReach.y))
			totalDist += fillSpace(galax.x, toReach.x, eColumns, 2)
			totalDist += fillSpace(galax.y, toReach.y, eLines, 2)
		}
	}
	resultFirst := totalDist

	endFirst := time.Since(endCommon)

	totalDist = 0
	for i, galax := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			toReach := galaxies[j]
			totalDist += (ftAbs(galax.x-toReach.x) + ftAbs(galax.y-toReach.y))
			totalDist += fillSpace(galax.x, toReach.x, eColumns, 1000000)
			totalDist += fillSpace(galax.y, toReach.y, eLines, 1000000)
		}
	}
	resultSecond := totalDist

	endSecond := time.Since(endCommon)

	fmt.Println("galaxies:", len(galaxies))
	fmt.Println("First part's result: ", resultFirst)
	fmt.Println("ExecTime first part: ", commonSinceStart+endFirst)
	fmt.Println("Second part's result:", resultSecond)
	fmt.Println("ExecTime second part:", commonSinceStart+(endSecond-endFirst))
}
