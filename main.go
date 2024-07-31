package main

import (
	"fmt"
	"os"

	"github.com/marco-souza/pkg/cmd/create"
)

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
	default:
		fmt.Println("command not found")
	}
}
