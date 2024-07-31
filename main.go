package main

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/cmd/create"
	"github.com/marco-souza/pkg/cmd/encrypt"
)

const DEFAULT_PASSPHRASE = "password" // TODO : get this from .pkg_pass file & -p flag

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command not found")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("package not found")
			os.Exit(1)
		}
		packageName := os.Args[2]

		if len(os.Args) < 4 {
			fmt.Println("folder not found")
		}

		folder := os.Args[3]
		create.CreatePackage(packageName, folder)

	case "encrypt":
		if len(os.Args) < 3 {
			fmt.Println("file not found")
			os.Exit(1)
		}

		filepath := path.Join(os.Args[2])
		encrypt.EncryptFile(filepath, DEFAULT_PASSPHRASE)

	case "decrypt":
		if len(os.Args) < 3 {
			fmt.Println("file not found")
			os.Exit(1)
		}

		filepath := path.Join(os.Args[2])
		encrypt.DecryptFile(filepath, DEFAULT_PASSPHRASE)

	default:
		fmt.Println("command not found")
	}
}
