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

func oasisPrediction(oasis []int) int64 {

	for i, j := 0, len(oasis)-1; i < j; i, j = i+1, j-1 {
		oasis[i], oasis[j] = oasis[j], oasis[i]
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
	// fmt.Println("fine oasis:", oasis, oasis[last+1])
	// fmt.Println("curr oasis:", oasis[x], x)
	return int64(oasis[last+1])
}

func main() {
	fContent, err := os.ReadFile("9th_day/input")
	check(err)

	start := time.Now()

	inputFile := strings.Split(string(fContent), "\n")

	var result int64 = 0

	count := 0
	for _, oasisStr := range inputFile {
		if oasisStr != "" {
			oasis := oasisConvertion(oasisStr)
			result += oasisPrediction(oasis)
			count++
		}
	}

	end := time.Since(start)

	fmt.Println("The result is: ", result, "for", count, "lines")
	fmt.Println("ExecTime is:", end)
}
