package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bl",
	Short: "Backlog CLI - GitHub CLI-like interface for Backlog",
	Long: `bl is a command line interface for Backlog that provides a GitHub CLI-like experience.
It allows you to interact with Backlog issues, pull requests, projects, and more from the terminal.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Backlog CLI (bl) - Use 'bl --help' for more information")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().String("format", "table", "output format (table, json, csv)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
