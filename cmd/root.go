/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pkg",
	Short: "CLI tool to speed-up software development",
	Long: `pkg is a CLI tool to speed-up software development by providing
	a set of commands to automate common tasks, such as creating a package module,
	encrypting and decrypting files, and more`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
