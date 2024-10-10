/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/marco-souza/pkg/internal/envs"
	"github.com/spf13/cobra"
)

var env = envs.NewEnv(".env")

// envsCmd represents the envs command
var envsCmd = &cobra.Command{
	Use:   "envs",
	Short: "Manage environment variables",
	Long:  `Manage environment variables. You can get, set and delete environment variables.`,
}

var envsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		value, err := env.GetEnv(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)
	},
}

var envsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments")
			return
		}

		err := env.SetEnv(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("update example file")
		if err := env.GenerateExample(); err != nil {
			fmt.Println(err)
			return
		}

		// as password from user input
		password := getPassphrase()

		fmt.Println("update encrypted file")
		encrypt.EncryptFile(env.Filepath, password)
	},
}

var envsDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		err := env.DetEnv(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("update example file")
		if err := env.GenerateExample(); err != nil {
			fmt.Println(err)
			return
		}

		// as password from user input
		password := getPassphrase()

		fmt.Println("update encrypted file")
		encrypt.EncryptFile(env.Filepath, password)
	},
}

func getPassphrase() string {
	passfile := ".pass"
	passphrase := DEFAULT_PASSPHRASE

	if _, err := os.Stat(passfile); err == nil {
		file, err := os.ReadFile(passfile)
		if err != nil {
			fmt.Println("Error reading passphrase file", err)
			os.Exit(1)
		}

		passphrase = string(file)
	}

	return passphrase
}

func init() {
	envsCmd.AddCommand(envsGetCmd)
	envsCmd.AddCommand(envsSetCmd)
	envsCmd.AddCommand(envsDelCmd)

	rootCmd.AddCommand(envsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
