package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRules(s string) [281]int {
	var rule [281]int

	for i, val := range s {
		if val == 'R' {
			rule[i] = 1
		}
	}
	return rule
}

func getTurns(inp []string, m map[string]int) [746][2]int {
	var turns [746][2]int

	for i, str := range inp {
		if i > 1 && str != "" {
			turns[i-2][0] = m[str[7:10]]
			turns[i-2][1] = m[str[12:15]]
		}
	}
	return turns
}

func itsOver(val int, b map[int]bool) bool {
	return b[val] || b[val] || b[val] || b[val] || b[val] || b[val]
}

func mcd(a int64, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return mcd(b, a%b)
}

func main() {
	fContent, err := os.ReadFile("8th_day/input")
	check(err)

	start := time.Now()

	inputFile := strings.Split(string(fContent), "\n")

	//format the rules so are 0 and 1
	rule := getRules(inputFile[0])

	//fill map so each str is valm
	m := make(map[string]int)
	var init [6]int
	over := make(map[int]bool)
	var countA, countZ int

	for i, str := range inputFile {
		if i > 1 && str != "" {
			m[str[:3]] = i - 2

			if str[2] == 'A' {
				init[countA] = i - 2
				countA++
			} else if str[2] == 'Z' {
				over[i-2] = true
				countZ++
			}
		}
	}

	//convert rule in num as map
	turns := getTurns(inputFile, m)

	//init the calc to find the first end for each init point
	var steps [6]int64
	for i, cur := range init {
		for !itsOver(cur, over) {
			cur = turns[cur][rule[steps[i]%int64(len(rule))]]
			steps[i]++
		}
	}

	//find the gcd (greater commod divsor) [mcd (massimo comune divisiore)]
	result := steps[0]
	for i := range steps {
		result = result / mcd(result, steps[i]) * steps[i]
	}

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
