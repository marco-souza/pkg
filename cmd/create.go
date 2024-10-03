/* Create a package module
 *
 * $ pkg create <package-name> <folder>
 */
package cmd

import (
	"fmt"

	"github.com/marco-souza/pkg/internal/pkg"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a package module for go",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Println("package not found")

		case 1, 2:
			packageName := args[0]
			message := "creating package " + packageName

			folder := ""
			if len(args) == 2 {
				folder = args[1]
				message += " in " + folder
			}

			fmt.Println(message)
			pkg.CreatePackage(packageName, folder)

		default:
			fmt.Println("too many arguments")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
