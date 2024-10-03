package cmd

import (
	"fmt"

	"github.com/marco-souza/pkg/internal/pkg"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:        "create",
	Short:      "Create a package module for go",
	Long:       `Create a package module for go`,
	Args:       cobra.RangeArgs(1, 2),
	ArgAliases: []string{"package", "folder"},
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		message := "creating package " + packageName

		folder := ""
		if len(args) == 2 {
			folder = args[1]
			message += " in " + folder
		}

		fmt.Println(message)
		pkg.CreatePackage(packageName, folder)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
