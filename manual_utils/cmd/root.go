package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for the CLI
var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "A simple CLI application",
	Long:  "This is a CLI application built with Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my CLI!")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}