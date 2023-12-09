package main

import (
	"bufio"
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

func ftGame(game string) bool {
	noSemicolon := strings.Replace(game, ";", " ", -1)
	noComma := strings.Replace(noSemicolon, ",", " ", -1)
	grab := strings.Fields(noComma)

	if len(grab)%2 != 0 {
		panic("grab must be even")
	}

	for i := 0; i < len(grab); i += 2 {
		atoi, err := strconv.Atoi(grab[i])
		check(err)
		if atoi < 12 {
			continue
		} else if strings.Compare(grab[i+1], "red") == 0 && atoi > 12 {
			return false
		} else if strings.Compare(grab[i+1], "green") == 0 && atoi > 13 {
			return false
		} else if strings.Compare(grab[i+1], "blue") == 0 && atoi > 14 {
			return false
		}
	}

	// fmt.Println(game + " is playble")
	return true
}

func main() {
	file, err := os.Open("2nd/input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	start := time.Now()

	for scanner.Scan() {
		mat := strings.Split(scanner.Text(), ":")
		if ftGame(mat[1]) {
			atoi, err := strconv.Atoi(mat[0][5:])
			check(err)
			result += atoi
		}
	}

	end := time.Since(start)

	fmt.Println("Result is:", result)
	fmt.Println("ExecTime is:", end)
}
