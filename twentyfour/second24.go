/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentyfour

// import (
// 	"advent-of-code/utils"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/spf13/cobra"
// )

// // twentyfourCmd represents the twentyfour command
// var secondCmd = &cobra.Command{
// 	Use:   "1st",
// 	Short: "Day 1 of Advent of Code 2024",
// 	Long: `Day 1 of Advent of Code 2024, finally my second Advent of Code!!!
// Here the problem statement: https://adventofcode.com/2024/day/1`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Executing %s\n\n", cmd.Short)

// 		inputPath := "twentyfour/inputs/one.input"
// 		file, error := os.Open(inputPath)
// 		utils.CheckInputFileError(error)
// 		defer file.Close()

// 		start := time.Now()
// 		result := dayTwoFirstPart(bufio.NewScanner(file))
// 		end := time.Since(start)

// 		if cmd.Flag("subject").Changed {
// 			firstSubject(true)
// 		}
// 		utils.PrintResult("one", result, end)

// 		file.Seek(0, 0)

// 		start = time.Now()
// 		result = dayTwoSecondPart(bufio.NewScanner(file))
// 		end = time.Since(start)

// 		if cmd.Flag("subject").Changed {
// 			firstSubject(false)
// 		}
// 		utils.PrintResult("two", result, end)

// 	},
// }

// func init() {
// 	TwentyfourCmd.AddCommand(secondCmd)
// }

// func dayTwoSecondPart(scanner *bufio.Scanner) (result int) {

// 	return
// }

// func dayTwoFirstPart(scanner *bufio.Scanner) (result int) {

// 	return
// }

// func secondSubject(partOne bool) {
// 	if partOne {
// 		fmt.Println()
// 	} else {
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }
