package encrypt_test

import (
	"github.com/marco-souza/pkg/cmd/encrypt"
	"testing"
)

func TestMain(t *testing.T) {
	// Add your test here
	encrypt.EncryptFile("main.go", "password")
	encrypt.DecryptFile("main.go.gpg", "password")
}
