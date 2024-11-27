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

func oasisConvertion(oasisStr string) []int {
	var e error

	valStr := strings.Fields(oasisStr)
	valNum := make([]int, len(valStr))
	for i, str := range valStr {
		valNum[i], e = strconv.Atoi(str)
		check(e)
	}
	return valNum
}

func oasisPrediction(oasis []int, second bool) int64 {

	if second {
		for i, j := 0, len(oasis)-1; i < j; i, j = i+1, j-1 {
			oasis[i], oasis[j] = oasis[j], oasis[i]
		}
	}

	last := len(oasis) - 1
	for x := 0; x != last+1; x++ {
		i := last
		for ; i > x; i-- {
			oasis[i] = oasis[i] - oasis[i-1]
		}
	}
	oasis = append(oasis, 0)
	x := last + 1
	for i := 0; i != x; i++ {
		x := last + 1
		for ; x > i; x-- {
			oasis[x] = oasis[x] + oasis[x-1]
		}
	}

	return int64(oasis[last+1])
}

func main() {

	fContent, err := os.ReadFile("9th_day/input")
	check(err)

	start := time.Now()

	inputFile := strings.Split(string(fContent), "\n")

	commonSinceStart := time.Since(start)
	endCommon := time.Now()

	var result int64 = 0
	var count int32 = 0
	for _, oasisStr := range inputFile {
		if oasisStr != "" {
			oasis := oasisConvertion(oasisStr)
			result += oasisPrediction(oasis, false)
			count++
		}
	}
	resultFirst := result
	endFirst := time.Since(endCommon)

	result = 0
	count = 0
	for _, oasisStr := range inputFile {
		if oasisStr != "" {
			oasis := oasisConvertion(oasisStr)
			result += oasisPrediction(oasis, true)
			count++
		}
	}
	resultSecond := result
	endSecond := time.Since(endCommon)

	fmt.Println("First part's result: ", resultFirst)
	fmt.Println("ExecTime first part: ", commonSinceStart+endFirst)
	fmt.Println("Second part's result:", resultSecond)
	fmt.Println("ExecTime second part:", commonSinceStart+(endSecond-endFirst))
}
