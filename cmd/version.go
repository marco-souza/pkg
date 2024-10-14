package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marco-souza/pkg/internal/semver"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version <patch|minor|major>",
	Short: "Manage semantinc versioning from the command line",
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

		v := map[string]interface{}{}
		if err := json.Unmarshal(file, &v); err != nil {
			fmt.Println(err)
			return
		}

		version := v["version"]
		if version == nil {
			fmt.Println("Version not found in file")
			return
		}

		sVersion, ok := version.(string)
		if !ok {
			fmt.Println("Version is not a string")
			return
		}

		fmt.Println("Current version:", sVersion)

		s := semver.SemVer{}
		if err := s.SetVersion(sVersion); err != nil {
			fmt.Println(err)
			return
		}

		if err := s.BumpVersion(release); err != nil {
			fmt.Println(err)
			return
		}

		// update json file
		v["version"] = s.GetVersion()
		newFile, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := os.WriteFile(f.Value.String(), newFile, 0644); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("New version:", s.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().StringP("file", "f", "package.json", "Versioning file must be a json file with a version key present")
}
