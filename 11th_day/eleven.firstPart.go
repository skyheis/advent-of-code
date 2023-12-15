package main

import (
	"bufio"
	"fmt"
	"os"
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

func addEmptyLine(inputRune [][]rune) [][]rune {
	var empty [][]rune
	flag := true

	for _, line := range inputRune {
		for _, char := range line {
			if char == '#' {
				flag = false
			}
		}
		empty = append(empty, line)
		if flag {
			empty = append(empty, line)
		}
		flag = true
	}
	return empty
}

func addEmptyColumn(inputFile [][]rune) [][]rune {
	// var space [][]rune
	space := make([][]rune, len(inputFile))

	for i := 0; i < len(inputFile); i++ {
		space = append(space, make([]rune, 0, len(inputFile[0])))
	}

	for x := 0; x < len(inputFile[0]); x++ {
		flag := true
		for y := 0; y < len(inputFile); y++ {
			if inputFile[y][x] == '#' {
				flag = false
			}
			space[y] = append(space[y], rune(inputFile[y][x]))
		}
		if flag {
			for y := 0; y < len(inputFile); y++ {
				space[y] = append(space[y], '.')
			}

		}
	}
	return space
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

	inputRune = addEmptyLine(inputRune)
	inputRune = addEmptyColumn(inputRune)

	galaxies := findGalaxies(inputRune)

	totalDist := 0
	fmt.Println("galaxies:", len(galaxies))
	for i, galax := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			toReach := galaxies[j]
			totalDist += (ftAbs(galax.x-toReach.x) + ftAbs(galax.y-toReach.y))
		}
	}

	end := time.Since(start)

	result := totalDist
	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
