/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentythree

import (
	"github.com/spf13/cobra"
)

// twentythreeCmd represents the twentythree command
var TwentythreeCmd = &cobra.Command{
	Use:   "2023",
	Short: "Adventure of Code 2023",
	Long: `My journey in Advent of Code started in 2023.
I did not finish all the puzzles, but I learned a lot, and I am proud of my progress.
I will continue to improve my skills and try to finish all the puzzles in the future.
You can find out more about the Advent of Code at https://adventofcode.com/2023`,
	Example: "advent-of-code 2023 1st",
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	TwentythreeCmd.PersistentFlags().BoolP("subject", "s", false, "Print the subject of the day before the result")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twentythreeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
