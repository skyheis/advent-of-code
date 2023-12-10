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

func main() {
	fContent, err := os.ReadFile("8th_day/input")
	check(err)

	start := time.Now()

	inputFile := strings.Split(string(fContent), "\n")

	//format the rules so are 0 and 1
	rule := getRules(inputFile[0])

	//fill map so each str is a num
	m := make(map[string]int)
	var init, over int
	for i, str := range inputFile {
		if i > 1 && str != "" {
			m[str[:3]] = i - 2
			if str[:3] == "AAA" {
				init = i - 2
			} else if str[:3] == "ZZZ" {
				over = i - 2
			}
		}
	}

	//convert rule in num as map
	turns := getTurns(inputFile, m)

	//init the calc
	steps := 0
	for x := init; x != over; steps++ {
		x = turns[x][rule[steps%len(rule)]]
	}

	end := time.Since(start)

	result := steps

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
