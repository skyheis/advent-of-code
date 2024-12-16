/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var fourteenCmd = &cobra.Command{
	Use:   "14th",
	Short: "Day 14 of Advent of Code 2024",
	Long: `Day 14 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/14`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/fourteen.input"
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result64 := dayFourteenFirstPart(bufio.NewScanner(file))
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			fourteenSubject(true)
		}
		utils.PrintResult64("one", result64, end)

		file.Seek(0, 0)

		start = time.Now()
		result64 = dayFourteenSecondPart(bufio.NewScanner(file))
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			fourteenSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(fourteenCmd)
}

type robot struct {
	position utils.Coordinates
	velocity utils.Coordinates
}

func buildRobot(str string) (r robot) {
	var err error

	split := strings.Fields(str)

	r.position.X, err = strconv.Atoi(strings.Split(split[0], ",")[0][2:])
	utils.CheckError(err)
	r.position.Y, err = strconv.Atoi(strings.Split(split[0], ",")[1])
	utils.CheckError(err)

	r.velocity.X, err = strconv.Atoi(strings.Split(split[1], ",")[0][2:])
	utils.CheckError(err)
	r.velocity.Y, err = strconv.Atoi(strings.Split(split[1], ",")[1])
	utils.CheckError(err)

	return
}

func createMatrix(robots []robot, seconds int) (matrix [103][101]int) {
	for _, robot := range robots {
		robot.position.X = (robot.position.X + robot.velocity.X*seconds) % 101
		robot.position.Y = (robot.position.Y + robot.velocity.Y*seconds) % 103

		if robot.position.X < 0 {
			robot.position.X = 101 + robot.position.X
		}
		if robot.position.Y < 0 {
			robot.position.Y = 103 + robot.position.Y
		}

		matrix[robot.position.Y][robot.position.X] += 1
	}

	return
}

func dayFourteenSecondPart(scanner *bufio.Scanner) (result int64) {
	const w = 101
	const h = 103

	var matrix [h][w]int

	robots := make([]robot, 500)
	i := 0
	for scanner.Scan() {
		robots[i] = buildRobot(scanner.Text())
		i++
	}

	for sec := 6510; sec < 6515; sec++ {
		matrix = createMatrix(robots, sec)
		fileName := fmt.Sprintf("%d.txt", sec)
		file, err := os.Create(fileName)
		if err != nil {
			utils.CheckError(err)
		}

		writer := bufio.NewWriter(file)
		for _, row := range matrix {
			for _, cell := range row {
				if cell == 0 {
					fmt.Fprintf(writer, " ")
				} else {
					fmt.Fprintf(writer, "%d", cell)
				}
			}
			fmt.Fprintln(writer)
		}
		writer.Flush()
		file.Close()
	}

	return 6512

}

func dayFourteenFirstPart(scanner *bufio.Scanner) (result int64) {
	const w = 101
	const h = 103

	seconds := 100
	var matrix [h][w]int

	for scanner.Scan() {
		robot := buildRobot(scanner.Text())

		robot.position.X = (robot.position.X + robot.velocity.X*seconds) % w
		robot.position.Y = (robot.position.Y + robot.velocity.Y*seconds) % h

		if robot.position.X < 0 {
			robot.position.X = w + robot.position.X
		}
		if robot.position.Y < 0 {
			robot.position.Y = h + robot.position.Y
		}

		matrix[robot.position.Y][robot.position.X] += 1

	}

	var res [4]int

	for y, row := range matrix {
		for x, cell := range row {

			if y < h/2 && x < w/2 {
				res[0] += cell
			} else if y < h/2 && x > w/2 {
				res[1] += cell
			} else if y > h/2 && x < w/2 {
				res[2] += cell
			} else if y > h/2 && x > w/2 {
				res[3] += cell
			}

		}
	}

	result = 1
	for _, r := range res {
		result *= int64(r)
	}

	return
}

func fourteenSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 14: Restroom Redoubt ---
One of The Historians needs to use the bathroom; fortunately, you know there's a bathroom near an unvisited location on their list, and so you're all quickly teleported directly to the lobby of Easter Bunny Headquarters.

Unfortunately, EBHQ seems to have "improved" bathroom security again after your last visit. The area outside the bathroom is swarming with robots!

To get The Historian safely to the bathroom, you'll need a way to predict where the robots will be in the future. Fortunately, they all seem to be moving on the tile floor in predictable straight lines.

You make a list (your puzzle input) of all of the robots' current positions (p) and velocities (v), one robot per line. For example:

	p=0,4 v=3,-3
	p=6,3 v=-1,-3
	p=10,3 v=-1,2
	p=2,0 v=2,-1
	p=0,0 v=1,3
	p=3,0 v=-2,-2
	p=7,6 v=-1,-3
	p=3,0 v=-1,-2
	p=9,3 v=2,3
	p=7,3 v=-1,2
	p=2,4 v=2,-3
	p=9,5 v=-3,-3

Each robot's position is given as p=x,y where x represents the number of tiles the robot is from the left wall and y represents the number of tiles from the top wall (when viewed from above). So, a position of p=0,0 means the robot is all the way in the top-left corner.

Each robot's velocity is given as v=x,y where x and y are given in tiles per second. Positive x means the robot is moving to the right, and positive y means the robot is moving down. So, a velocity of v=1,-2 means that each second, the robot moves 1 tile to the right and 2 tiles up.

The robots outside the actual bathroom are in a space which is 101 tiles wide and 103 tiles tall (when viewed from above). However, in this example, the robots are in a space which is only 11 tiles wide and 7 tiles tall.

The robots are good at navigating over/under each other (due to a combination of springs, extendable legs, and quadcopters), so they can share the same tile and don't interact with each other. Visually, the number of robots on each tile in this example looks like this:

	1.12.......
	...........
	...........
	......11.11
	1.1........
	.........1.
	.......1...

These robots have a unique feature for maximum bathroom security: they can teleport. When a robot would run into an edge of the space they're in, they instead teleport to the other side, effectively wrapping around the edges. Here is what robot p=2,4 v=2,-3 does for the first few seconds:

Initial state:

	...........
	...........
	...........
	...........
	..1........
	...........
	...........

After 1 second:

	...........
	....1......
	...........
	...........
	...........
	...........
	...........

After 2 seconds:

	...........
	...........
	...........
	...........
	...........
	......1....
	...........

After 3 seconds:

	...........
	...........
	........1..
	...........
	...........
	...........
	...........

After 4 seconds:

	...........
	...........
	...........
	...........
	...........
	...........
	..........1

After 5 seconds:

	...........
	...........
	...........
	.1.........
	...........
	...........
	...........

The Historian can't wait much longer, so you don't have to simulate the robots for very long. Where will the robots be after 100 seconds?

In the above example, the number of robots on each tile after 100 seconds has elapsed looks like this:

......2..1.
...........
1..........
.11........
.....1.....
...12......
.1....1....

To determine the safest area, count the number of robots in each quadrant after 100 seconds. Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant, so the only relevant robots are:

..... 2..1.
..... .....
1.... .....
           
..... .....
...12 .....
.1... 1....

In this example, the quadrants contain 1, 3, 4, and 1 robot. Multiplying these together gives a total safety factor of 12.

Predict the motion of the robots in your list within a space which is 101 tiles wide and 103 tiles tall. What will the safety factor be after exactly 100 seconds have elapsed?`)
	} else {
		fmt.Println(`--- Part Two ---
During the bathroom break, someone notices that these robots seem awfully similar to ones built and used at the North Pole. If they're the same type of robots, they should have a hard-coded Easter egg: very rarely, most of the robots should arrange themselves into a picture of a Christmas tree.

What is the fewest number of seconds that must elapse for the robots to display the Easter egg?`)
	}
	fmt.Println()
}
