package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/spf13/cobra"
)

const DEFAULT_PASSPHRASE = "password"

var decryptCmd = &cobra.Command{
	Use:        "decrypt",
	Short:      "Decrypt a file",
	Long:       `Decrypt a file using the passphrase provided.`,
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

		// TODO: add timeout in case passphrase is not valid
		if err := encrypt.DecryptFile(filepath, passphrase); err != nil {
			fmt.Println("Error decrypting file", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
