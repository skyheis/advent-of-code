package one

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func secondPart(scanner *bufio.Scanner) int {
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

	return result
}

func firstPart(scanner *bufio.Scanner) int {
	result := 0

	for scanner.Scan() {
		value := 0
		flag := false
		line := scanner.Text()

		for i := 0; !flag && i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				value = (int(line[i]) - 48) * 10
				flag = true
			}
		}
		for i := len(line) - 1; flag && i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				value += int(line[i]) - 48
				flag = false
			}
		}
		result += value
	}

	return result
}

func dayOne() {
	fileOne, error := os.Open("input")
	check(error)
	defer fileOne.Close()

	firstStart := time.Now()
	resultFirst := firstPart(bufio.NewScanner(fileOne))
	firstEnd := time.Since(firstStart)

	fileTwo, error := os.Open("input")
	check(error)
	defer fileTwo.Close()

	secondStart := time.Now()
	resultSecond := secondPart(bufio.NewScanner(fileTwo))
	secondEnd := time.Since(secondStart)

	fmt.Println("First part's result: ", resultFirst)
	fmt.Println("ExecTime first part: ", firstEnd)
	fmt.Println("Second part's result:", resultSecond)
	fmt.Println("ExecTime second part:", secondEnd)
}
