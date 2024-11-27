/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"advent-of-code/twentythree"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "My solutions for Advent of Code",
	Long: `Advent of Code is an Advent calendar of small programming puzzles for a variety
of skill sets and skill levels that can be solved in any programming language you like.
You can find out more about the Advent of Code at https://adventofcode.com/`,
}

func main() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code.yaml)")
	rootCmd.AddCommand(twentythree.TwentythreeCmd)

	defaultHelpTemplate := rootCmd.HelpTemplate()

	rootCmd.SetHelpTemplate(helpTemplate)
	twentythree.TwentythreeCmd.SetHelpTemplate(defaultHelpTemplate)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const helpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}

Years Commands:{{range .Commands}}{{if and (ne .Name "help") (ne .Name "completion")}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

Others Commands:{{range .Commands}}{{if or (eq .Name "help") (eq .Name "completion")}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
`
