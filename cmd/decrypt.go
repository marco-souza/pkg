package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:        "decrypt",
	Short:      "Decrypt a file",
	Long:       `Decrypt a file using the passphrase provided.`,
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		filepath := path.Join(os.Getenv("PWD"), filename)
		passphrase := Must(encrypt.GetPassphase(passphraseFlag))

		// TODO: add timeout in case passphrase is not valid
		if err := encrypt.DecryptFile(filepath, passphrase); err != nil {
			fmt.Println("Error decrypting file", err)
			os.Exit(1)
		}

		fmt.Println("File decrypted successfully:", filepath[:len(filepath)-4])
	},
}

func init() {
	decryptCmd.Flags().StringVarP(&passphraseFlag, "passphrase", "p", "", "Passphrase to use for encryption (optional, uses .pass file as fallback)")
	rootCmd.AddCommand(decryptCmd)
}
