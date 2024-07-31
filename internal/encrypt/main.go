package encrypt

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/openpgp" // TODO: replace with: https://github.com/ProtonMail/gopenpgp
)

// encrypt a file like gpg
func EncryptFile(filepath, passphrase string) {
	if filepath == "" {
		fmt.Println("please provide a file path")
		return
	}

	// read filepath
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	// create file .gpg
	outputFile, err := os.Create(filepath + ".gpg")
	if err != nil {
		fmt.Println("error creating output file", err)
		return
	}

	// create a new writer
	w, err := openpgp.SymmetricallyEncrypt(outputFile, []byte(passphrase), nil, nil)
	if err != nil {
		fmt.Println("error creating writer", err)
		return
	}

	if _, err := io.Copy(w, file); err != nil {
		fmt.Println("error copying file", err)
		return
	}

	// close the writer
	if err := w.Close(); err != nil {
		fmt.Println("error closing writer", err)
		return
	}
}

func DecryptFile(filepath, passphrase string) {
	if filepath == "" {
		fmt.Println("please provide a file path")
		return
	}

	if !strings.Contains(filepath, ".gpg") {
		filepath += ".gpg"
	}

	// read filepath
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	// decrypt symmetrically encrypted file with passphrase
	promptFunc := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if !symmetric {
			return nil, fmt.Errorf("not a symmetric file")
		}
		return []byte(passphrase), nil
	}

	msgDetails, err := openpgp.ReadMessage(file, nil, promptFunc, nil)
	if err != nil {
		fmt.Println("error reading message", err)
		return
	}

	fmt.Println("decrypting file", msgDetails.LiteralData.FileName, "to", filepath[:len(filepath)-4])

	// create file without .gpg
	outputFile, err := os.Create(filepath[:len(filepath)-4])
	if err != nil {
		fmt.Println("error creating output file", err)
		return
	}

	if _, err := io.Copy(outputFile, msgDetails.UnverifiedBody); err != nil {
		fmt.Println("error copying file", err)
		return
	}
}
