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
var sixthCmd = &cobra.Command{
	Use:   "6th",
	Short: "Day 6 of Advent of Code 2024",
	Long: `Day 6 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/6`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/six.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := daySixCommonPart(string(filecontent), true)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			sixthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = daySixCommonPart(string(filecontent), false)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			sixthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(sixthCmd)
}

type coordinates struct {
	x int
	y int
}

const (
	VI_STEP  = 'X'
	VI_OBST  = '#'
	VI_UP    = '^'
	VI_RIGHT = '>'
	VI_DOWN  = 'v'
	VI_LEFT  = '<'
	VI_TURN  = 'O'
)

func findGuard(labMap [][]rune) (pos coordinates) {
	for i, line := range labMap {
		for k, guard := range line {
			if guard == VI_UP {
				pos.x = k
				pos.y = i
				return
			}
		}
	}
	return pos
}

func countWalked(labMap [][]rune) (result int) {
	for _, line := range labMap {
		for _, spot := range line {
			if spot == VI_STEP || spot == VI_TURN {
				result++
			}
		}
	}

	return
}

func newGuardPosition(labMap [][]rune, x, y int) (int, int, rune) {
	steppedOn := '.'

	switch labMap[y][x] {
	case VI_UP:
		if y == 0 {
			labMap[y][x] = VI_STEP
			return -1, -1, steppedOn
		}
		if labMap[y-1][x] == VI_OBST {
			labMap[y][x] = VI_TURN
			x++
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_RIGHT
		} else {
			labMap[y][x] = VI_STEP
			y--
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_UP
		}

	case VI_RIGHT:
		if x == len(labMap[y])-1 {
			labMap[y][x] = VI_STEP
			return -1, -1, steppedOn
		}
		if labMap[y][x+1] == VI_OBST {
			labMap[y][x] = VI_TURN
			y++
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_DOWN
		} else {
			labMap[y][x] = VI_STEP
			x++
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_RIGHT
		}

	case VI_DOWN:
		if y == len(labMap)-2 {
			labMap[y][x] = VI_STEP
			return -1, -1, steppedOn
		}
		if labMap[y+1][x] == VI_OBST {
			labMap[y][x] = VI_TURN
			x--
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_LEFT
		} else {
			labMap[y][x] = VI_STEP
			y++
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_DOWN
		}

	default:
		if x == 0 {
			labMap[y][x] = VI_STEP
			return -1, -1, steppedOn
		}
		if labMap[y][x-1] == VI_OBST {
			labMap[y][x] = VI_TURN
			y--
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_UP
		} else {
			labMap[y][x] = VI_STEP
			x--
			steppedOn = labMap[y][x]
			labMap[y][x] = VI_LEFT
		}
	}

	return x, y, steppedOn
}

func isFrontObst(labMap [][]rune, curSpot coordinates) bool {
	switch labMap[curSpot.y][curSpot.x] {
	case VI_UP:
		curSpot.y--
	case VI_RIGHT:
		curSpot.x++
	case VI_DOWN:
		curSpot.y++
	default:
		curSpot.x--
	}

	return labMap[curSpot.y][curSpot.x] == VI_OBST
}

func followTheGuard(labMap [][]rune, curSpot coordinates) bool {
	var steppedOn rune
	curSpot.x, curSpot.y, steppedOn = newGuardPosition(labMap, curSpot.x, curSpot.y)

	if curSpot.x == -1 && curSpot.y == -1 {
		return true
	} else if steppedOn == VI_TURN && isFrontObst(labMap, curSpot) {
		labMap[curSpot.y][curSpot.x] = VI_TURN
		return false
	}

	return followTheGuard(labMap, curSpot)
}

func addObstacle(filecontent string, startPoint, newObst coordinates) bool {
	newLab := utils.MakeRuneMatrixStr(filecontent)
	newLab[newObst.y][newObst.x] = VI_OBST
	// if !followTheGuard(newLab, startPoint) {
	// 	fmt.Println("loop!", newObst)
	// 	return true
	// }

	return !followTheGuard(newLab, startPoint)
}

func howToLoopGuard(filecontent string, labMap [][]rune, startSpot coordinates) (result int) {
	labMap[startSpot.y][startSpot.x] = VI_UP

	for y, line := range labMap {
		for x, spot := range line {
			if spot == VI_STEP || spot == VI_TURN {
				if addObstacle(filecontent, startSpot, coordinates{x: x, y: y}) {
					result++
				}
			}
		}
	}

	return
}

func daySixCommonPart(filecontent string, first bool) (result int) {
	labMap := utils.MakeRuneMatrixStr(filecontent)

	startSpot := findGuard(labMap)
	followTheGuard(labMap, startSpot)

	// labMap[startSpot.y][startSpot.x] = VI_UP

	utils.WriteFileFromRuneMatrix(labMap, "sasa")

	if first {
		result = countWalked(labMap)
	} else {
		result = howToLoopGuard(filecontent, labMap, startSpot)
	}

	return
}

func sixthSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 6: Guard Gallivant ---
The Historians use their fancy device again, this time to whisk you all away to the North Pole prototype suit manufacturing lab... in the year 1518! It turns out that having direct access to history is very convenient for a group of historians.

You still have to be careful of time paradoxes, and so it will be important to avoid anyone from 1518 while The Historians search for the Chief. Unfortunately, a single guard is patrolling this part of the lab.

Maybe you can work out where the guard will go ahead of time so that The Historians can search safely?

You start by making a map (your puzzle input) of the situation. For example:

	....#.....
	.........#
	..........
	..#.......
	.......#..
	..........
	.#..^.....
	........#.
	#.........
	......#...

The map shows the current position of the guard with ^ (to indicate the guard is currently facing up from the perspective of the map). Any obstructions - crates, desks, alchemical reactors, etc. - are shown as #.

Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly following these steps:

 - If there is something directly in front of you, turn right 90 degrees.
 - Otherwise, take a step forward.
Following the above protocol, the guard moves up several times until she reaches an obstacle (in this case, a pile of failed suit prototypes):

	....#.....
	....^....#
	..........
	..#.......
	.......#..
	..........
	.#........
	........#.
	#.........
	......#...

Because there is now an obstacle in front of the guard, she turns right before continuing straight in her new facing direction:

	....#.....
	........>#
	..........
	..#.......
	.......#..
	..........
	.#........
	........#.
	#.........
	......#...

Reaching another obstacle (a spool of several very long polymers), she turns right again and continues downward:

	....#.....
	.........#
	..........
	..#.......
	.......#..
	..........
	.#......v.
	........#.
	#.........
	......#...

This process continues for a while, but the guard eventually leaves the mapped area (after walking past a tank of universal solvent):

	....#.....
	.........#
	..........
	..#.......
	.......#..
	..........
	.#........
	........#.
	#.........
	......#v..

By predicting the guard's route, you can determine which specific positions in the lab will be in the patrol path. Including the guard's starting position, the positions visited by the guard before leaving the area are marked with an X:

	....#.....
	....XXXXX#
	....X...X.
	..#.X...X.
	..XXXXX#X.
	..X.X.X.X.
	.#XXXXXXX.
	.XXXXXXX#.
	#XXXXXXX..
	......#X..

In this example, the guard will visit 41 distinct positions on your map.

Predict the path of the guard. How many distinct positions will the guard visit before leaving the mapped area?`)
	} else {
		fmt.Println(`--- Part Two ---
While The Historians begin working around the guard's patrol route, you borrow their fancy device and step outside the lab. From the safety of a supply closet, you time travel through the last few months and record the nightly status of the lab's guard post on the walls of the closet.

Returning after what seems like only a few seconds to The Historians, they explain that the guard's patrol area is simply too large for them to safely search the lab without getting caught.

Fortunately, they are pretty sure that adding a single new obstruction won't cause a time paradox. They'd like to place the new obstruction in such a way that the guard will get stuck in a loop, making the rest of the lab safe to search.

To have the lowest chance of creating a time paradox, The Historians would like to know all of the possible positions for such an obstruction. The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.

In the above example, there are only 6 different positions where a new obstruction would cause the guard to get stuck in a loop. The diagrams of these six situations use O to mark the new obstruction, | to show a position where the guard moves up/down, - to show a position where the guard moves left/right, and + to show a position where the guard moves both up/down and left/right.

Option one, put a printing press next to the guard's starting position:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	....|..#|.
	....|...|.
	.#.O^---+.
	........#.
	#.........
	......#...

Option two, put a stack of failed suit prototypes in the bottom right quadrant of the mapped area:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	..+-+-+#|.
	..|.|.|.|.
	.#+-^-+-+.
	......O.#.
	#.........
	......#...

Option three, put a crate of chimney-squeeze prototype fabric next to the standing desk in the bottom right quadrant:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	..+-+-+#|.
	..|.|.|.|.
	.#+-^-+-+.
	.+----+O#.
	#+----+...
	......#...

Option four, put an alchemical retroencabulator near the bottom left corner:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	..+-+-+#|.
	..|.|.|.|.
	.#+-^-+-+.
	..|...|.#.
	#O+---+...
	......#...

Option five, put the alchemical retroencabulator a bit to the right instead:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	..+-+-+#|.
	..|.|.|.|.
	.#+-^-+-+.
	....|.|.#.
	#..O+-+...
	......#...

Option six, put a tank of sovereign glue right next to the tank of universal solvent:

	....#.....
	....+---+#
	....|...|.
	..#.|...|.
	..+-+-+#|.
	..|.|.|.|.
	.#+-^-+-+.
	.+----++#.
	#+----++..
	......#O..

It doesn't really matter what you choose to use as an obstacle so long as you and The Historians can put it into position without the guard noticing. The important thing is having enough options that you can find one that minimizes time paradoxes, and in this example, there are 6 different positions you could choose.

You need to get the guard stuck in a loop by adding a single new obstruction. How many different positions could you choose for this obstruction?`)
	}
	fmt.Println()
}

//
