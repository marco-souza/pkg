/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:        "encrypt",
	Short:      "Encrypt a file",
	Long:       `Encrypt a file using the passphrase provided.`,
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		filepath := path.Join(os.Getenv("PWD"), filename)

		// get password from .pass file
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

		fmt.Println("encrypt file", filepath)
		encrypt.EncryptFile(filepath, passphrase)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
