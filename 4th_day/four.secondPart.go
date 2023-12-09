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

func strToSortInt(s []string) []int {
	var intArr []int
	for _, str := range s {
		num, err := strconv.Atoi(str)
		check(err)
		intArr = append(intArr, num)
	}
	slices.Sort(intArr)
	return intArr
}

func main() {
	file, error := os.Open("4th_day/input")
	check(error)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	start := time.Now()

	var hand [192]int

	for i := range hand {
		hand[i] = 1
	}

	for i := 0; scanner.Scan(); i++ {
		cardValue := 0
		text := scanner.Text()
		card := strings.Split(text[9:], "|")
		winning := strToSortInt(strings.Fields(card[0]))
		playable := strToSortInt(strings.Fields(card[1]))

		for _, num := range winning {
			_, found := slices.BinarySearch(playable, num)
			if found {
				cardValue++
			}
		}

		for ; cardValue > 0; cardValue-- {
			hand[i+cardValue] += hand[i]
		}
	}

	result := func() int {
		ret := 0
		for _, num := range hand {
			ret += num
		}
		return ret
	}

	end := time.Since(start)

	fmt.Println("The result is: ", result())
	fmt.Println("ExecTime is:", end)
}
