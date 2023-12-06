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

func ftGame(game string) int {
	noSemicolon := strings.Replace(game, ";", " ", -1)
	noComma := strings.Replace(noSemicolon, ",", " ", -1)
	grab := strings.Fields(noComma)

	if len(grab)%2 != 0 {
		panic("grab must be even")
	}

	minRed := 0
	minGreen := 0
	minBlue := 0

	for i := 0; i < len(grab); i += 2 {
		atoi, err := strconv.Atoi(grab[i])
		check(err)
		if strings.Compare(grab[i+1], "red") == 0 && atoi > minRed {
			minRed = atoi
		} else if strings.Compare(grab[i+1], "green") == 0 && atoi > minGreen {
			minGreen = atoi
		} else if strings.Compare(grab[i+1], "blue") == 0 && atoi > minBlue {
			minBlue = atoi
		}
	}

	return minRed * minGreen * minBlue
}

func main() {
	file, err := os.Open("day_two/input")
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
			result.Add(int32(ftGame(mat[1])))
			wg.Done()
		}(scanner.Text())
	}

	wg.Wait()

	end := time.Since(start)

	fmt.Println("Result is:", result.Load())
	fmt.Println("ExecTime is:", end)
}
