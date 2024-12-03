/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package twentyfour

import (
	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var TwentyfourCmd = &cobra.Command{
	Use:   "2024",
	Short: "Adventure of Code 2024",
	Long: `My journey in Advent of Code started in 2024.
I did not finish all the puzzles, but I learned a lot, and I am proud of my progress.
I will continue to improve my skills and try to finish all the puzzles in the future.
You can find out more about the Advent of Code at https://adventofcode.com/2024`,
	Example: "advent-of-code 2024 1st",
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	TwentyfourCmd.PersistentFlags().BoolP("subject", "s", false, "Print the subject of the day before the result")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twentyfourCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
