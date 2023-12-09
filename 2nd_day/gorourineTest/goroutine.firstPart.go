package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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
	file, err := os.Open("2nd_day/input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result atomic.Int32
	var wg sync.WaitGroup

	start := time.Now()

	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			mat := strings.Split(line, ":")
			if ftGame(mat[1]) {
				atoi, err := strconv.Atoi(mat[0][5:])
				check(err)
				result.Add(int32(atoi))
			}
			wg.Done()
		}(scanner.Text())
	}

	wg.Wait()

	end := time.Since(start)

	fmt.Println("Result is:", result.Load())
	fmt.Println("ExecTime is:", end)
}
