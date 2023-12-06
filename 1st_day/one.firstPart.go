package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, error := os.Open("day_one/input")
	check(error)
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	fmt.Println("The result is: ", result)
}
