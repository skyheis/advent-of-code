package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func digitOrNumber(line string, numbs [9]string) int {

	for i := 0; i < len(numbs); i++ {
		if strings.HasPrefix(line, numbs[i]) {
			return i + 1
		}
	}
	if unicode.IsDigit(rune(line[0])) {
		return int(line[0]) - 48
	}
	return 0
}

func main() {
	file, error := os.Open("1st_day/input")
	check(error)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	numbs := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		value := 0
		flag := false
		line := scanner.Text()

		for i := 0; !flag && i < len(line); i++ {
			digit := digitOrNumber(line[i:], numbs)
			if digit > 0 {
				value = digit * 10
				flag = true
			}
		}
		for i := len(line) - 1; flag && i >= 0; i-- {
			digit := digitOrNumber(line[i:], numbs)
			if digit > 0 {
				value += digit
				flag = false
			}
		}
		result += value
	}

	fmt.Println("The result is: ", result)
}
