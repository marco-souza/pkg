package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/pkg/internal/encrypt"
	"github.com/marco-souza/pkg/internal/pkg"
)

const DEFAULT_PASSPHRASE = "password" // TODO : get this from .pkg_pass file & -p flag

func Execute() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("command not found\n")
	}

	cmd := os.Args[1]
	switch cmd {
	case "create":
		if len(os.Args) < 3 {
			return fmt.Errorf("package not found\n")
		}
		packageName := os.Args[2]

		if len(os.Args) < 4 {
			fmt.Println("W: folder not found")
		}
		folder := os.Args[3]

		pkg.CreatePackage(packageName, folder)

	case "encrypt":
		if len(os.Args) < 3 {
			return fmt.Errorf("file not found")
		}

		filepath := path.Join(os.Args[2])
		encrypt.EncryptFile(filepath, DEFAULT_PASSPHRASE)

	case "decrypt":
		if len(os.Args) < 3 {
			return fmt.Errorf("file not found")
		}

		filepath := path.Join(os.Args[2])
		encrypt.DecryptFile(filepath, DEFAULT_PASSPHRASE)
	}

	return nil
}

func Fallback() {
	help()
}

func help() {
	fmt.Println("usage: pkg <cmd>")
	fmt.Println("")
	fmt.Println("commands:")
	fmt.Println("  $ pkg create <package-name> [folder] - create a go package in specified folder")
	fmt.Println("  $ pkg encrypt <filename> - encrypt a file into a new <filename>.gpg file")
	fmt.Println("  $ pkg decrypt <filename> - decrypt a gpg file into a file with the same name but without .gpg")
	fmt.Println("")
}
