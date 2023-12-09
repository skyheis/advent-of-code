package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSeeds(s string) []int {
	strings := strings.Split(s[7:], " ")
	var intArr []int
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		check(err)
		intArr = append(intArr, num)
	}
	return intArr
}

func getRanges(t string) (int, int, int) {
	var e error
	splitted := strings.Split(t, " ")
	var dr, sr, r int
	if len(splitted) == 3 {
		dr, e = strconv.Atoi(splitted[0])
		check(e)
		sr, e = strconv.Atoi(splitted[1])
		check(e)
		r, e = strconv.Atoi(splitted[2])
		check(e)
	} else {
		panic("ranges must be 3")
	}
	return dr, sr, r
}

func execChange(source *[]int, fullDr []int, fullSr []int, fullR []int) {
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

func main() {
	file, error := os.Open("5th_day/input")
	check(error)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	start := time.Now()

	scanner.Scan()

	threeNil := func() ([]int, []int, []int) {
		return nil, nil, nil
	}
	fullDr, fullSr, fullR := threeNil()

	source := getSeeds(scanner.Text())
	fmt.Println("Seeds to be planted:", len(source))

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			execChange(&source, fullDr, fullSr, fullR)
			fullDr, fullSr, fullR = threeNil()
			if !scanner.Scan() {
				break
			}
			fmt.Println("Init", scanner.Text())
		} else {
			dr, sr, r := getRanges(text)
			fullDr = append(fullDr, dr)
			fullSr = append(fullSr, sr)
			fullR = append(fullR, r)
		}
	}
	execChange(&source, fullDr, fullSr, fullR)
	fmt.Println("Sorting")
	slices.Sort(source)

	result := source[0]

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
