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
var thirteenthCmd = &cobra.Command{
	Use:   "13th",
	Short: "Day 13 of Advent of Code 2024",
	Long: `Day 13 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/13`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/thirteen.input"
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result64 := dayThirteenCommonPart(bufio.NewScanner(file), false)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			thirteenthSubject(true)
		}
		utils.PrintResult64("one", result64, end)

		file.Seek(0, 0)

		start = time.Now()
		result64 = dayThirteenCommonPart(bufio.NewScanner(file), true)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			thirteenthSubject(false)
		}
		utils.PrintResult64("two", result64, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(thirteenthCmd)
}

func getButtonPressed(ax, ay, bx, by, px, py int64, secondPart bool) (int64, bool) {
	if secondPart {
		px += 10000000000000
		py += 10000000000000
	}

	frac := ax*by - ay*bx

	A := (px*by - py*bx) / frac
	B := (ax*py - ay*px) / frac

	// check if A * ax + B * bx == px and A * ay + B * by == py
	if !secondPart && (A > 100 || B > 100) {
		return 0, false
	}

	if A*ax+B*bx != px || A*ay+B*by != py {
		return 0, false
	}

	return A*3 + B, true
}

func getStringValues(scanner *bufio.Scanner, delimiter string) (int64, int64) {
	splitted := strings.Split(scanner.Text(), delimiter)
	scanner.Scan()

	x, err := strconv.ParseInt(splitted[1][:len(splitted[1])-3], 10, 64)
	utils.CheckError(err)

	y, err := strconv.ParseInt(splitted[2], 10, 64)
	utils.CheckError(err)

	return x, y
}

func dayThirteenCommonPart(scanner *bufio.Scanner, secondPart bool) (result int64) {

	for scanner.Scan() {
		ax, ay := getStringValues(scanner, "+")

		bx, by := getStringValues(scanner, "+")

		px, py := getStringValues(scanner, "=")

		tokens, valid := getButtonPressed(ax, ay, bx, by, px, py, secondPart)
		if valid {
			result += tokens
		}

	}

	return
}

// https://www.youtube.com/watch?v=jBsC34PxzoM
// https://www.youtube.com/watch?v=vXqlIOX2itM

// A*a_x + B*B_x = p_x
// A*a_y + B*b_y = p_y
// A = (p_x*b_y - p_y*b_x) / (a_x*b_y - a_y*b_x)
// B = (a_x*p_y - a_y*p_x) / (a_x*b_y - a_y*b_x)

func thirteenthSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 13: Claw Contraption ---
Next up: the lobby of a resort on a tropical island. The Historians take a moment to admire the hexagonal floor tiles before spreading out.

Fortunately, it looks like the resort has a new arcade! Maybe you can win some prizes from the claw machines?

The claw machines here are a little unusual. Instead of a joystick or directional buttons to control the claw, these machines have two buttons labeled A and B. Worse, you can't just put in a token and play; it costs 3 tokens to push the A button and 1 token to push the B button.

With a little experimentation, you figure out that each machine's buttons are configured to move the claw a specific amount to the right (along the X axis) and a specific amount forward (along the Y axis) each time that button is pressed.

Each machine contains one prize; to win the prize, the claw must be positioned exactly above the prize on both the X and Y axes.

You wonder: what is the smallest number of tokens you would have to spend to win as many prizes as possible? You assemble a list of every machine's button behavior and prize location (your puzzle input). For example:

	Button A: X+94, Y+34
	Button B: X+22, Y+67
	Prize: X=8400, Y=5400

	Button A: X+26, Y+66
	Button B: X+67, Y+21
	Prize: X=12748, Y=12176

	Button A: X+17, Y+86
	Button B: X+84, Y+37
	Prize: X=7870, Y=6450

	Button A: X+69, Y+23
	Button B: X+27, Y+71
	Prize: X=18641, Y=10279

This list describes the button configuration and prize location of four different claw machines.

For now, consider just the first claw machine in the list:

 - Pushing the machine's A button would move the claw 94 units along the X axis and 34 units along the Y axis.
 - Pushing the B button would move the claw 22 units along the X axis and 67 units along the Y axis.
 - The prize is located at X=8400, Y=5400; this means that from the claw's initial position, it would need to move exactly 8400 units along the X axis and exactly 5400 units along the Y axis to be perfectly aligned with the prize in this machine.

The cheapest way to win the prize is by pushing the A button 80 times and the B button 40 times. This would line up the claw along the X axis (because 80*94 + 40*22 = 8400) and along the Y axis (because 80*34 + 40*67 = 5400). Doing this would cost 80*3 tokens for the A presses and 40*1 for the B presses, a total of 280 tokens.

For the second and fourth claw machines, there is no combination of A and B presses that will ever win a prize.

For the third claw machine, the cheapest way to win the prize is by pushing the A button 38 times and the B button 86 times. Doing this would cost a total of 200 tokens.

So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is 480.

You estimate that each button would need to be pressed no more than 100 times to win a prize. How else would someone be expected to play?

Figure out how to win as many prizes as possible. What is the fewest tokens you would have to spend to win all possible prizes?`)
	} else {
		fmt.Println(``)
	}
	fmt.Println()
}
