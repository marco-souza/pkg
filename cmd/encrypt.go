package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:        "encrypt",
	Short:      "Encrypt a file",
	Long:       `Encrypt a file using the passphrase provided.`,
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		filepath := path.Join(os.Getenv("PWD"), filename)
		passphrase := Must(encrypt.GetPassphase(passphraseFlag))

		if err := encrypt.EncryptFile(filepath, passphrase); err != nil {
			fmt.Println("Error encrypting file", err)
			os.Exit(1)
		}

		fmt.Println("File encrypted successfully:", filepath+".gpg")
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&passphraseFlag, "passphrase", "p", "", "Passphrase to use for encryption (optional, uses .pass file as fallback)")
	rootCmd.AddCommand(encryptCmd)
}
