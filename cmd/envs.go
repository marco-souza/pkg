package cmd

import (
	"fmt"

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
	Use:   "get <key>",
	Short: "Get an environment variable",
	Args:  cobra.ExactArgs(1),
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
	Use:   "set <key> <value>",
	Short: "Set an environment variable",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Println("update encrypted file")
		passphrase := Must(encrypt.GetPassphase(passphraseFlag))
		encrypt.EncryptFile(env.Filepath, passphrase)
	},
}

var envsDelCmd = &cobra.Command{
	Use:   "del <key>",
	Short: "Delete an environment variable",
	Args:  cobra.ExactArgs(1),
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

		fmt.Println("update encrypted file")
		passphrase := Must(encrypt.GetPassphase(passphraseFlag))
		encrypt.EncryptFile(env.Filepath, passphrase)
	},
}

func init() {
	envsCmd.Flags().StringVarP(&passphraseFlag, "passphrase", "p", "", "Passphrase to use for encryption (optional, uses .pass file as fallback)")

	envsCmd.AddCommand(envsGetCmd)
	envsCmd.AddCommand(envsSetCmd)
	envsCmd.AddCommand(envsDelCmd)

	rootCmd.AddCommand(envsCmd)
}
