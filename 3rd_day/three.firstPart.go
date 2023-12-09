package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func replaceChar(file_mat *[]string, x int, y int) {
	s := (*file_mat)[y]
	r := []rune(s)
	r[x] = '.'
	s = string(r)
	(*file_mat)[y] = s
}

func ftCheckSymbol(file_mat *[]string, symbolFlag *bool, x int, y int) {
	if *symbolFlag || x < 0 || y < 0 || x >= len((*file_mat)[y]) || y >= len(*file_mat) {
		return
	} else if (*file_mat)[y][x] == '.' {
		return
	} else if unicode.IsDigit(rune((*file_mat)[y][x])) {
		replaceChar(file_mat, x, y)
		ftCheckSymbol(file_mat, symbolFlag, x+1, y)
		ftCheckSymbol(file_mat, symbolFlag, x-1, y-1)
		ftCheckSymbol(file_mat, symbolFlag, x, y-1)
		ftCheckSymbol(file_mat, symbolFlag, x+1, y-1)
		ftCheckSymbol(file_mat, symbolFlag, x-1, y)
		ftCheckSymbol(file_mat, symbolFlag, x-1, y+1)
		ftCheckSymbol(file_mat, symbolFlag, x, y+1)
		ftCheckSymbol(file_mat, symbolFlag, x+1, y+1)
	} else {
		*symbolFlag = true
	}
}

func missingPart(file_mat []string) int {
	result := 0
	// valid := 0
	// invalid := 0

	var width int
	height := len(file_mat)

	if height > 0 {
		width = len(file_mat[0])
	} else {
		panic("File is empty")
	}

	for y := 0; y < height-1; y++ {
		for x := 0; x < width; x++ {
			s := 0
			for x+s < width && unicode.IsDigit(rune(file_mat[y][x+s])) {
				s++
			}

			if s > 0 {
				num, err := strconv.Atoi(file_mat[y][x : x+s])
				check(err)
				symbolFlag := false
				ftCheckSymbol(&file_mat, &symbolFlag, x, y)
				if symbolFlag {
					result += num
					// valid++
				} //else {
				// invalid++
				//}
				x += s
			}
		}
	}
	// fmt.Println("Valid:", valid, "Invalid:", invalid)

	return result
}

func main() {
	filecontent, err := os.ReadFile("3rd_day/input")
	check(err)

	file_mat := strings.Split(string(filecontent), "\n")

	start := time.Now()

	result := missingPart(file_mat)

	end := time.Since(start)

	fmt.Println("Result is:", result)
	fmt.Println("ExecTime is:", end)
}
