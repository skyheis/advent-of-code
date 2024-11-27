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

func setDest(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = -1
	}
	return arr
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

func fillEmptyDest(d *[]int, s []int) {
	for i, val := range *d {
		if val == -1 {
			(*d)[i] = s[i]
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
	source := getSeeds(scanner.Text())
	destination := setDest(len(source))

	for scanner.Scan() {

		text := scanner.Text()
		if text == "" {
			// fmt.Println("\ndest is:", destination, "| source is:", source)
			fillEmptyDest(&destination, source)
			source = destination
			destination = nil
			destination = setDest(len(source))
			// fmt.Println("current:", source[0])
			if !scanner.Scan() {
				break
			}
		} else {
			dr, sr, r := getRanges(text)
			for i, val := range source {
				if destination[i] == -1 && val >= sr && val < sr+r {
					// fmt.Println("in this range:", text, "| new value:", dr+(val-sr))
					destination[i] = dr + (val - sr)
				}
			}
		}
	}
	fillEmptyDest(&destination, source)
	slices.Sort(destination)

	result := destination[0]

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
