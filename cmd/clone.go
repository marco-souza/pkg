package cmd

import (
	"fmt"

	"github.com/marco-souza/pkg/internal/git"
	"github.com/spf13/cobra"
)

type repository struct {
	protocol string
	registry string
	repo     string
	user     string
}

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <repository> <name>",
	Short: "Use a repository as a template for a new project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]
		name := args[1]

		fmt.Println("Cloning repository", repo, "into", name)

		if err := git.Clone(repo, name); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Repository cloned successfully")
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
