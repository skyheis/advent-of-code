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

func execChange(source *[]int, destRange []int, sourRange []int, intRange []int) {
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
	destRange, sourRange, intRange := threeNil()

	source := getSeeds(scanner.Text())
	fmt.Println("Seeds to be planted:", len(source))

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			changeTime := time.Now()
			execChange(&source, destRange, sourRange, intRange)
			endChangeTime := time.Since(changeTime)
			fmt.Println("Time to change values", endChangeTime)

			destRange, sourRange, intRange = threeNil()
			if !scanner.Scan() {
				break
			}
			fmt.Println("Init", scanner.Text())
		} else {
			dr, sr, r := getRanges(text)
			destRange = append(destRange, dr)
			sourRange = append(sourRange, sr)
			intRange = append(intRange, r)
		}
	}

	execChange(&source, destRange, sourRange, intRange)
	fmt.Println("Sorting")
	slices.Sort(source)

	result := source[0]

	end := time.Since(start)

	fmt.Println("The result is:", result)
	fmt.Println("ExecTime is:", end)
}
