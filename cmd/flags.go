package cmd

import (
	"fmt"
	"os"
)

// encrypt and decrypt commands
var passphraseFlag string

func Must[T any](v T, err error) T {
	if err != nil {
		if err != nil {
			fmt.Println("Error getting required value", err)
			os.Exit(1)
		}
	}
	return v
}

func Ensure(err error) {
	if err != nil {
		fmt.Println("Error getting required value", err)
		os.Exit(1)
	}
}
