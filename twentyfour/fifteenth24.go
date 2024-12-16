/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var fifteenthCmd = &cobra.Command{
	Use:   "15th",
	Short: "Day 15 of Advent of Code 2024",
	Long: `Day 15 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/15`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/fifteen.input.test2"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result64 := dayFifteenFirstPart(filecontent)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			fifteenSubject(true)
		}
		utils.PrintResult64("one", result64, end)

		start = time.Now()
		result64 = dayFifteenSecondPart(filecontent)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			fifteenSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(fifteenthCmd)
}

const (
	XV_ROBOT = '@'
	XV_WALL  = '#'
	XV_BOX   = 'O'
	XV_BOXL  = '['
	XV_BOXR  = ']'
	XV_EMPTY = '.'
	XV_UP    = '^'
	XV_DOWN  = 'v'
	XV_LEFT  = '<'
	XV_RIGHT = '>'
)

func findRobot(area [][]rune) (pos utils.Coordinates) {
	for y, line := range area {
		for x, robot := range line {
			if robot == XV_ROBOT {
				pos.X = x
				pos.Y = y
				return
			}
		}
	}
	return pos
}

func drawRobotInArea(area [][]rune, robot *utils.Coordinates, move func(utils.Coordinates) utils.Coordinates) {
	area[robot.Y][robot.X] = XV_EMPTY
	*robot = move(*robot)
	area[robot.Y][robot.X] = XV_ROBOT
}

func getDirection(dir byte) (dx, dy int, newDir byte) {
	switch dir {
	case XV_UP:
		dy = -1
		newDir = 0
	case XV_RIGHT:
		dx = 1
		newDir = 1
	case XV_DOWN:
		dy = 1
		newDir = 2
	case XV_LEFT:
		dx = -1
		newDir = 3
	}

	return
}

func robotMoveBoxes(area [][]rune, robot utils.Coordinates, rules string) {
	if len(rules) == 0 {
		return
	}

	dir := rules[0]

	moveMap := []func(utils.Coordinates) utils.Coordinates{
		utils.Coordinates.Up,
		utils.Coordinates.Right,
		utils.Coordinates.Down,
		utils.Coordinates.Left,
	}

	dx, dy, dir := getDirection(dir)

	nextX, nextY := robot.X+dx, robot.Y+dy
	if area[nextY][nextX] == XV_BOX {
		i := 1
		for area[robot.Y+i*dy][robot.X+i*dx] != XV_EMPTY && area[robot.Y+i*dy][robot.X+i*dx] != XV_WALL {
			i++
		}
		if area[robot.Y+i*dy][robot.X+i*dx] == XV_EMPTY {
			area[robot.Y+i*dy][robot.X+i*dx] = XV_BOX
			drawRobotInArea(area, &robot, moveMap[dir])
		}
	} else if area[nextY][nextX] == XV_EMPTY {
		drawRobotInArea(area, &robot, moveMap[dir])
	}

	robotMoveBoxes(area, robot, rules[1:])
}

func dayFifteenFirstPart(filecontent []byte) (result int64) {
	strslice := strings.Split(string(filecontent), "\n")

	i := 0
	var area [][]rune
	for _, str := range strslice {
		if str == "" {
			break
		}
		area = append(area, []rune(str))
		i++
	}

	var rules string
	for _, str := range strslice[i+1:] {
		rules += str
	}

	robotMoveBoxes(area, findRobot(area), rules)

	for y, row := range area {
		for x, spot := range row {
			if spot == XV_BOX {
				result += int64(100*y + x)
			}
		}
	}

	return
}

func robotMoveBigBoxes(area [][]rune, robot utils.Coordinates, rules string) {
	if len(rules) == 0 {
		return
	}

	dir := rules[0]

	moveMap := []func(utils.Coordinates) utils.Coordinates{
		utils.Coordinates.Up,
		utils.Coordinates.Right,
		utils.Coordinates.Down,
		utils.Coordinates.Left,
	}

	dx, dy, dir := getDirection(dir)

	nextX, nextY := robot.X+dx, robot.Y+dy
	if area[nextY][nextX] == XV_BOXL || area[nextY][nextX] == XV_BOXR {

		// for area[robot.Y+i*dy][robot.X+i*dx] != XV_EMPTY && area[robot.Y+i*dy][robot.X+i*dx] != XV_WALL {
		// 	i++
		// }
		// if area[robot.Y+i*dy][robot.X+i*dx] == XV_EMPTY {
		// 	area[robot.Y+i*dy][robot.X+i*dx] = XV_BOX
		// 	drawRobotInArea(area, &robot, moveMap[dir])
		// }
	} else if area[nextY][nextX] == XV_EMPTY {
		drawRobotInArea(area, &robot, moveMap[dir])
	}

	robotMoveBigBoxes(area, robot, rules[1:])
}

func dayFifteenSecondPart(filecontent []byte) (result int64) {
	var rules string
	matrix := utils.MakeRuneMatrixByte(filecontent)

	h := 0
	for _, row := range matrix {
		if len(row) == 0 {
			break
		}
		h++
	}

	area := make([][]rune, h)
	for y, row := range matrix {
		for _, spot := range row {
			if spot == XV_WALL {
				area[y] = append(area[y], XV_WALL)
				area[y] = append(area[y], XV_WALL)
			} else if spot == XV_EMPTY {
				area[y] = append(area[y], XV_EMPTY)
				area[y] = append(area[y], XV_EMPTY)
			} else if spot == XV_BOX {
				area[y] = append(area[y], XV_BOXL)
				area[y] = append(area[y], XV_BOXR)
			} else if spot == XV_ROBOT {
				area[y] = append(area[y], XV_ROBOT)
				area[y] = append(area[y], XV_EMPTY)
			}
		}
	}

	for ; h < len(matrix); h++ {
		rules += string(matrix[h])
	}

	utils.PrintRuneMatrix(area)
	fmt.Println(rules)
	robotMoveBigBoxes(area, findRobot(area), rules)

	// for y, row := range area {
	// 	for x, spot := range row {
	// 		if spot == XV_BOX {
	// 			result += int64(100*y + x)
	// 		}
	// 	}
	// }

	return
}

func fifteenSubject(partOne bool) {
	if partOne {
		fmt.Println(``)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}
