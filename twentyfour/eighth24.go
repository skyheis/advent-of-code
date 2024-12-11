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
var eighthCmd = &cobra.Command{
	Use:   "8th",
	Short: "Day 8 of Advent of Code 2024",
	Long: `Day 8 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/8`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/eight.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayEighCommonPart(string(filecontent), false)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			eighthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = dayEighCommonPart(string(filecontent), true)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			eighthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(eighthCmd)
}

func inMatrix(area [][]rune, node utils.Coordinates) bool {
	return node.Y >= 0 && node.Y < len(area) && node.X >= 0 && node.X < len(area[0])
}

func addAntinode(area [][]rune, overlap map[utils.Coordinates]bool, antenna, node utils.Coordinates, bonus bool) {
	var antinode utils.Coordinates

	deltaX := utils.Abs(antenna.X - node.X)
	deltaY := utils.Abs(antenna.Y - node.Y)
	if antenna.X > node.X {
		antinode.X = antenna.X + deltaX
	} else {
		antinode.X = antenna.X - deltaX
	}
	if antenna.Y > node.Y {
		antinode.Y = antenna.Y + deltaY
	} else {
		antinode.Y = antenna.Y - deltaY
	}

	if inMatrix(area, antinode) {
		if area[antinode.Y][antinode.X] == '.' {
			area[antinode.Y][antinode.X] = '#'
		} else if area[antinode.Y][antinode.X] != '#' {
			overlap[antinode] = true
		}

		if bonus {
			addAntinode(area, overlap, antinode, antenna, bonus)
		}
	}

}

func checkOtherAntennas(area [][]rune, overlap map[utils.Coordinates]bool, antenna rune, cordAntenna utils.Coordinates, bonus bool) {

	for y, line := range area {
		for x, spot := range line {
			if spot == antenna && (cordAntenna != utils.Coordinates{X: x, Y: y}) {
				addAntinode(area, overlap, cordAntenna, utils.Coordinates{X: x, Y: y}, bonus)
			}
		}
	}
}

func dayEighCommonPart(filecontent string, bonus bool) (result int) {
	overlap := make(map[utils.Coordinates]bool)
	area := utils.MakeRuneMatrixStr(filecontent)
	area = area[:len(area)-1] //add

	for y, line := range area {
		for x, antenna := range line {
			if antenna != '.' && antenna != '#' {
				checkOtherAntennas(area, overlap, antenna, utils.Coordinates{X: x, Y: y}, bonus)
			}
		}
	}

	utils.PrintRuneMatrix(area)

	if bonus {
		result = (len(area) - 1) * len(area[0])
		result -= utils.CountInMatrix(area, '.')
	} else {
		result = len(overlap)
		result += utils.CountInMatrix(area, '#')
	}
	return
}

func eighthSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 8: Resonant Collinearity ---
You find yourselves on the roof of a top-secret Easter Bunny installation.

While The Historians do their thing, you take a look at the familiar huge antenna. Much to your surprise, it seems to have been reconfigured to emit a signal that makes people 0.1% more likely to buy Easter Bunny brand Imitation Mediocre Chocolate as a Christmas gift! Unthinkable!

Scanning across the city, you find that there are actually many such antennas. Each antenna is tuned to a specific frequency indicated by a single lowercase letter, uppercase letter, or digit. You create a map (your puzzle input) of these antennas. For example:

	............
	........0...
	.....0......
	.......0....
	....0.......
	......A.....
	............
	............
	........A...
	.........A..
	............
	............

The signal only applies its nefarious effect at specific antinodes based on the resonant frequencies of the antennas. In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but only when one of the antennas is twice as far away as the other. This means that for any pair of antennas with the same frequency, there are two antinodes, one on either side of them.

So, for these two antennas with frequency a, they create the two antinodes marked with #:

	..........
	...#......
	..........
	....a.....
	..........
	.....a....
	..........
	......#...
	..........
	..........

Adding a third antenna with the same frequency creates several more antinodes. It would ideally add four antinodes, but two are off the right side of the map, so instead it adds only two:

	..........
	...#......
	#.........
	....a.....
	........a.
	.....a....
	..#.......
	......#...
	..........
	..........

Antennas with different frequencies don't create antinodes; A and a count as different frequencies. However, antinodes can occur at locations that contain antennas. In this diagram, the lone antenna with frequency capital A creates no antinodes but has a lowercase-a-frequency antinode at its location:

	..........
	...#......
	#.........
	....a.....
	........a.
	.....a....
	..#.......
	......A...
	..........
	..........

The first example has antennas with two different frequencies, so the antinodes they create look like this, plus an antinode overlapping the topmost A-frequency antenna:

	......#....#
	...#....0...
	....#0....#.
	..#....0....
	....0....#..
	.#....A.....
	...#........
	#......#....
	........A...
	.........A..
	..........#.
	..........#.

Because the topmost A-frequency antenna overlaps with a 0-frequency antinode, there are 14 total unique locations that contain an antinode within the bounds of the map.

Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?`)
	} else {
		fmt.Println(`Watching over your shoulder as you work, one of The Historians asks if you took the effects of resonant harmonics into your calculations.

Whoops!

After updating your model, it turns out that an antinode occurs at any grid position exactly in line with at least two antennas of the same frequency, regardless of distance. This means that some of the new antinodes will occur at the position of each antenna (unless that antenna is the only one of its frequency).

So, these three T-frequency antennas now create many antinodes:

	T....#....
	...T......
	.T....#...
	.........#
	..#.......
	..........
	...#......
	..........
	....#.....
	..........

In fact, the three T-frequency antennas are all exactly in line with two antennas, so they are all also antinodes! This brings the total number of antinodes in the above example to 9.

The original example now has 34 antinodes, including the antinodes that appear on every antenna:

	##....#....#
	.#.#....0...
	..#.#0....#.
	..##...0....
	....0....#..
	.#...#A....#
	...#..#.....
	#....#.#....
	..#.....A...
	....#....A..
	.#........#.
	...#......##

Calculate the impact of the signal using this updated model. How many unique locations within the bounds of the map contain an antinode?`)
	}
	fmt.Println()
}

//
