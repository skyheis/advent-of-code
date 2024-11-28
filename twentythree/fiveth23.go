/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var FivethCmd = &cobra.Command{
	Use:   "5th",
	Short: "Day 5 of Advent of Code 2023",
	Long: `Day 5 of Advent of Code 2023, my very first day in Advent of Code!!!
Here the problem statement: https://adventofcode.com/2023/day/5`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentythree/inputs/five.input"
		file, error := os.Open(inputPath)
		utils.CheckInputFileError(error)
		defer file.Close()

		start := time.Now()
		result := dayFiveFirstPart(bufio.NewScanner(file))
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			fivethSubject(true)
		}
		utils.PrintResult("one", result, end)

		file.Seek(0, 0)

		start = time.Now()
		result = dayFiveSecondPart(bufio.NewScanner(file))
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			fivethSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentythreeCmd.AddCommand(FivethCmd)
}

func getSeedsFirstPart(s string) []int {
	strings := strings.Split(s[7:], " ")
	var intArr []int
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		utils.CheckPanic(err)
		intArr = append(intArr, num)
	}
	return intArr
}

func getRangesFirstPart(t string) (int, int, int) {
	var e error
	splitted := strings.Split(t, " ")
	var dr, sr, r int
	if len(splitted) == 3 {
		dr, e = strconv.Atoi(splitted[0])
		utils.CheckPanic(e)
		sr, e = strconv.Atoi(splitted[1])
		utils.CheckPanic(e)
		r, e = strconv.Atoi(splitted[2])
		utils.CheckPanic(e)
	} else {
		panic("ranges must be 3")
	}
	return dr, sr, r
}

func execChangeFirstPart(source *[]int, fullDr []int, fullSr []int, fullR []int) {
	if fullR != nil {
		for i, val := range *source {
			the_val := val
			for x, vSr := range fullSr {
				if val >= vSr && val <= vSr+fullR[x] {
					(*source)[i] = fullDr[x] + (the_val - vSr)
					break
				}
			}
		}
	}
}

func dayFiveFirstPart(scanner *bufio.Scanner) int {

	scanner.Scan()

	threeNil := func() ([]int, []int, []int) {
		return nil, nil, nil
	}
	fullDr, fullSr, fullR := threeNil()

	source := getSeedsFirstPart(scanner.Text())
	// fmt.Println("Seeds to be planted:", len(source))

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			execChangeFirstPart(&source, fullDr, fullSr, fullR)
			fullDr, fullSr, fullR = threeNil()
			if !scanner.Scan() {
				break
			}
			// fmt.Println("Init", scanner.Text())
		} else {
			dr, sr, r := getRangesFirstPart(text)
			fullDr = append(fullDr, dr)
			fullSr = append(fullSr, sr)
			fullR = append(fullR, r)
		}
	}
	execChangeFirstPart(&source, fullDr, fullSr, fullR)
	slices.Sort(source)

	return source[0]
}

// ////////////////
func getSeedsSecondPart(s string) []int {
	strings := strings.Split(s[7:], " ")
	var intArr []int
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		utils.CheckPanic(err)
		intArr = append(intArr, num)
	}
	size := 0
	for i, val := range intArr {
		if i%2 == 1 {
			size += val
		}
	}

	intBigArr := make([]int, size)
	bigI := 0
	for i, val := range intArr {
		if i%2 == 0 {
			rng := intArr[i+1]
			for j := 0; j < rng; j++ {
				intBigArr[bigI+j] = val + j
			}
			bigI += rng
		}
	}

	return intBigArr
}

func getRangesSecondPart(t string) (int, int, int) {
	var e error
	splitted := strings.Split(t, " ")
	var dr, sr, r int
	if len(splitted) == 3 {
		dr, e = strconv.Atoi(splitted[0])
		utils.CheckPanic(e)
		sr, e = strconv.Atoi(splitted[1])
		utils.CheckPanic(e)
		r, e = strconv.Atoi(splitted[2])
		utils.CheckPanic(e)
	} else {
		panic("ranges must be 3")
	}
	return dr, sr, r
}

func execChangeSecondPart(source *[]int, destRange []int, sourRange []int, intRange []int) {
	if intRange != nil {
		for i, val := range *source {
			for x, vSr := range sourRange {
				if val >= vSr && val <= vSr+intRange[x] {
					(*source)[i] = destRange[x] + (val - vSr)
				}
			}
		}
	}
}

func dayFiveSecondPart(scanner *bufio.Scanner) int {

	scanner.Scan()
	threeNil := func() ([]int, []int, []int) {
		return nil, nil, nil
	}
	destRange, sourRange, intRange := threeNil()

	source := getSeedsSecondPart(scanner.Text())
	// fmt.Println("Seeds to be planted:", len(source))

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			// changeTime := time.Now()
			execChangeSecondPart(&source, destRange, sourRange, intRange)
			// endChangeTime := time.Since(changeTime)
			// fmt.Println("Time to change values", endChangeTime)

			destRange, sourRange, intRange = threeNil()
			if !scanner.Scan() {
				break
			}
			fmt.Println("Init", scanner.Text())
		} else {
			dr, sr, r := getRangesSecondPart(text)
			destRange = append(destRange, dr)
			sourRange = append(sourRange, sr)
			intRange = append(intRange, r)
		}
	}

	execChangeSecondPart(&source, destRange, sourRange, intRange)
	// fmt.Println("Sorting")
	slices.Sort(source)

	return source[0]
}

/////////////////

func fivethSubject(partOne bool) {
	if partOne {
		fmt.Println()
	} else {
		fmt.Println()
	}
	fmt.Println()
}
