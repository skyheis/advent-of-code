package utils

import (
	"fmt"
	"os"
)

func CheckInputFileError(e error) {
	if e != nil {
		fmt.Println("Sorry, input file invalid:", e)
		os.Exit(1)
	}
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func CheckPanic(e error) {
	if e != nil {
		panic(e)
	}
}
