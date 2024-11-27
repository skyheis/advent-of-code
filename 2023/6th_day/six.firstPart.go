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

func getTimes(m []string) ([]int, []int) {
	var e error

	s1 := strings.Fields(m[0][9:])
	raceTime := make([]int, len(s1))
	for i, val := range s1 {
		raceTime[i], e = strconv.Atoi(val)
		check(e)
	}

	s1 = strings.Fields(m[1][9:])
	distance := make([]int, len(s1))
	for i, val := range s1 {
		distance[i], e = strconv.Atoi(val)
		check(e)
	}

	fmt.Println(raceTime)
	fmt.Println(distance)
	return raceTime, distance
}

func main() {
	filecontent, err := os.ReadFile("6th_day/input")
	check(err)

	start := time.Now()

	result := 1
	raceTime, distance := getTimes(strings.Split(string(filecontent), "\n"))

	for i, val := range raceTime {
		x := 1
		y := val - 1
		toBeat := distance[i]
		for x*(val-x) <= toBeat {
			x++
		}
		// fmt.Println("x is", x)
		for y*(val-y) <= toBeat {
			y--
		}
		// fmt.Println("y is", y)
		// fmt.Println("result of", i+1, "is", y-x+1)
		result *= y - x + 1
	}

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
