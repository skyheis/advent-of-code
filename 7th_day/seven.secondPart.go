package main

import (
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

type Hand struct {
	hVal int
	bet  int
	typo rune
}

var cards = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func calcHighers(m map[rune]rune) (rune, rune) {
	var fst rune = 0
	var snd rune = 0
	var tmp rune = 0
	var jly rune = 0

	for key, value := range m {
		if key == 'J' {
			jly = value
		} else if value > fst {
			fst = value
			tmp = key
		}
	}
	delete(m, tmp)
	for key, value := range m {
		if key != 'J' && value > snd {
			snd = value
		}
	}
	return fst + jly, snd
}

func getHandValue(s string) int {
	value := 0
	for _, card := range s {
		for i, val := range cards {
			if card == val {
				value = value*100 + i + 1
			}
		}
	}
	return value
}

func getTypo(s string) rune {
	m := make(map[rune]rune)

	for _, card := range s {
		_, exist := m[card]
		if exist {
			m[card]++
		} else {
			m[card] = 1
		}
	}
	fst, snd := calcHighers(m)
	switch {
	case fst == 5:
		return 7
	case fst == 4:
		return 6
	case fst == 3 && snd == 2:
		return 5
	case fst == 3:
		return 4
	case fst == 2 && snd == 2:
		return 3
	case fst == 2:
		return 2
	default:
		return 1
	}
}

func main() {
	fContent, err := os.ReadFile("7th_day/input")
	check(err)

	start := time.Now()

	inputHands := strings.Split(string(fContent), "\n")

	var allHands [1000]Hand
	var e error

	for i, str := range inputHands {
		if i < 1000 {
			allHands[i].hVal = getHandValue(str[:5])
			allHands[i].bet, e = strconv.Atoi(str[6:])
			check(e)
			allHands[i].typo = getTypo(str[:5])
			// fmt.Println(allHands[i])
		}
	}

	var tmp Hand
	sorted := false

	for !sorted {
		sorted = true
		for i, hand := range allHands {
			if i == len(allHands)-1 {
				break
			}
			if hand.typo > allHands[i+1].typo {
				tmp = allHands[i+1]
				allHands[i+1] = allHands[i]
				allHands[i] = tmp
				sorted = false
			} else if hand.typo == allHands[i+1].typo {
				if hand.hVal > allHands[i+1].hVal {
					tmp = allHands[i+1]
					allHands[i+1] = allHands[i]
					allHands[i] = tmp
					sorted = false
				}
			}
		}
	}
	// fmt.Println(allHands)

	result := 0
	for i, typo := range allHands {
		result += typo.bet * (i + 1)
	}

	end := time.Since(start)

	fmt.Println("The result is: ", result)
	fmt.Println("ExecTime is:", end)
}
