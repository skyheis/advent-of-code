package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTimes(m []string) (int, int) {
	var e error

	s1 := strings.Fields(m[0][9:])
	s0 := strings.Join(s1, "")
	raceTime, e := strconv.Atoi(s0)
	check(e)

	s1 = strings.Fields(m[1][9:])
	s0 = strings.Join(s1, "")
	distance, e := strconv.Atoi(s0)
	check(e)

	return raceTime, distance
}

func main() {
	filecontent, err := os.ReadFile("6th_day/example")
	check(err)

	start := time.Now()

	raceTime, distance := getTimes(strings.Split(string(filecontent), "\n"))

	x := 1
	y := raceTime - 1
	for x*(raceTime-x) <= distance {
		x++
	}
	fmt.Println("min is", x)
	for y*(raceTime-y) <= distance {
		y--
	}
	fmt.Println("max is", y)
	result := y - x + 1

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
