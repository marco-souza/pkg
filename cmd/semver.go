package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marco-souza/pkg/internal/semver"
	"github.com/spf13/cobra"
)

var semverCmd = &cobra.Command{
	Use:   "semver <patch|minor|major>",
	Short: "Manage semantinc semvering from the command line",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flag("file")
		release := "patch"
		if len(args) > 0 {
			release = args[0]
		}

		// read json file
		file, err := os.ReadFile(f.Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		v := map[string]any{}
		if err := json.Unmarshal(file, &v); err != nil {
			fmt.Println(err)
			return
		}

		version := v["version"]
		if version == nil {
			fmt.Println("semver not found in file")
			return
		}

		cur_version, ok := version.(string)
		if !ok {
			fmt.Println("semver is not a string")
			return
		}

		fmt.Println("Current semver:", cur_version)

		s := semver.SemVer{}
		Ensure(s.SetVersion(cur_version))
		Ensure(s.BumpVersion(release))

		// update json file
		v["semver"] = s.GetVersion()
		newFile := Must(json.MarshalIndent(v, "", "    "))

		Ensure(os.WriteFile(f.Value.String(), newFile, 0644))

		fmt.Println("New semver:", s.GetVersion())
	},
}

func init() {
	semverCmd.Flags().StringP("file", "f", "package.json", "json file with version key")
	rootCmd.AddCommand(semverCmd)
}
