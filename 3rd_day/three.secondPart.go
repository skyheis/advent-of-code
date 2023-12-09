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

func ftCheckSymbol(file_mat *[]string, symbolFlag *bool, num *int, x int, y int) {
	if *symbolFlag || x < 0 || y < 0 || x >= len((*file_mat)[y]) || y >= len(*file_mat) {
		return
	} else if (*file_mat)[y][x] == '.' {
		return
	} else if unicode.IsDigit(rune((*file_mat)[y][x])) {

		r := 0
		for x+r < len((*file_mat)[y]) && unicode.IsDigit(rune((*file_mat)[y][x+r])) {
			r++
		}
		l := 0
		for x-l >= 0 && unicode.IsDigit(rune((*file_mat)[y][x-l])) {
			l++
		}
		atoi, err := strconv.Atoi((*file_mat)[y][x-l+1 : x+r])
		for i := x - l + 1; i < x+r; i++ {
			replaceChar(file_mat, i, y)
		}
		check(err)
		if *num == 0 {
			*num = atoi
		} else {
			*num *= atoi
			*symbolFlag = true
		}

	} else if (*file_mat)[y][x] == '*' {
		replaceChar(file_mat, x, y)
		ftCheckSymbol(file_mat, symbolFlag, num, x+1, y)
		ftCheckSymbol(file_mat, symbolFlag, num, x-1, y-1)
		ftCheckSymbol(file_mat, symbolFlag, num, x, y-1)
		ftCheckSymbol(file_mat, symbolFlag, num, x+1, y-1)
		ftCheckSymbol(file_mat, symbolFlag, num, x-1, y)
		ftCheckSymbol(file_mat, symbolFlag, num, x-1, y+1)
		ftCheckSymbol(file_mat, symbolFlag, num, x, y+1)
		ftCheckSymbol(file_mat, symbolFlag, num, x+1, y+1)
	}
}

//unicode.IsDigit(rune((*file_mat)[y][x]))

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
			if file_mat[y][x] == '*' {
				num := 0
				symbolFlag := false
				ftCheckSymbol(&file_mat, &symbolFlag, &num, x, y)
				if symbolFlag {
					result += num
					// valid++
				} //else {
				// invalid++
				//}
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
