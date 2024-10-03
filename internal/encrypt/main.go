package encrypt

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/openpgp" // TODO: replace with: https://github.com/ProtonMail/gopenpgp
)

// encrypt a file like gpg
func EncryptFile(filepath, passphrase string) error {
	if filepath == "" {
		fmt.Println("please provide a file path")
		return fmt.Errorf("no file path provided")
	}

	fmt.Println("encrypting file", filepath)

	// read filepath
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	// create file .gpg
	outputFile, err := os.Create(filepath + ".gpg")
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}

	// create a new writer
	w, err := openpgp.SymmetricallyEncrypt(outputFile, []byte(passphrase), nil, nil)
	if err != nil {
		return fmt.Errorf("error creating writer: %w", err)
	}

	if _, err := io.Copy(w, file); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	// close the writer
	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing writer: %w", err)
	}

	return nil
}

func DecryptFile(filepath, passphrase string) error {
	if filepath == "" {
		return fmt.Errorf("no file path provided")
	}

	if !strings.Contains(filepath, ".gpg") {
		filepath += ".gpg"
	}

	output := filepath[:len(filepath)-4]
	fmt.Println("decrypting file", filepath, "to", output)

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
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
		return fmt.Errorf("error reading message: %w", err)
	}

	// create file without .gpg
	outputFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}

	if _, err := io.Copy(outputFile, msgDetails.UnverifiedBody); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}
