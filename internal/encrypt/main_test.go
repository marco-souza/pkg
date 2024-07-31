package encrypt_test

import (
	"os"
	"testing"

	"github.com/marco-souza/pkg/internal/encrypt"
)

func TestMain(t *testing.T) {
	// Add your test here
	encrypt.EncryptFile("main.go", "password")
	encrypt.DecryptFile("main.go.gpg", "password")
	os.Remove("main.go.gpg")
}
