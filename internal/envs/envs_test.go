package envs_test

import (
	"os"
	"testing"

	"github.com/marco-souza/pkg/internal/envs"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// create tmp file
	filepath := "/tmp/.env"
	if err := os.WriteFile(filepath, []byte("DB_USER=root"), 0644); err != nil {
		t.Errorf("WriteFile() error = %v", err)
	}

	env := envs.NewEnv(filepath)
	assert.NotNil(t, env)

	t.Run("GenerateExample", func(t *testing.T) {
		assert.NoError(t, env.GenerateExample())
		// check if file exists
		_, err := os.Stat(filepath + ".example")
		assert.NoError(t, err)

		// remove file
		os.Remove(filepath + ".example")
	})

	t.Run("GetEnv", func(t *testing.T) {
		value, err := env.GetEnv("DB_USER")
		assert.NoError(t, err)
		assert.Equal(t, "root", value)
	})

	t.Run("SetEnv", func(t *testing.T) {
		err := env.SetEnv("DB_PASSWORD", "123")
		assert.NoError(t, err)

		value, err := env.GetEnv("DB_PASSWORD")
		assert.NoError(t, err)
		assert.Equal(t, "123", value)
	})

	t.Run("DelEnv", func(t *testing.T) {
		err := env.DetEnv("DB_PASSWORD")
		assert.NoError(t, err)

		_, err = env.GetEnv("DB_PASSWORD")
		assert.Error(t, err)
	})

	os.Remove(filepath)
}
