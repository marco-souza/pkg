package git_test

import (
	"os"
	"path"
	"testing"

	"github.com/marco-souza/pkg/internal/git"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testRepo := "marco-souza/pkg"
	testName := "test_pkg_tmp"

	t.Run("clone invalid repo", func(t *testing.T) {
		err := git.Clone("invalid_repo", testName)
		assert.Error(t, err)
	})

	t.Run("clone valid repo", func(t *testing.T) {
		err := git.Clone(testRepo, testName)
		assert.NoError(t, err)
		assert.DirExists(t, path.Join(testName))

		err = os.RemoveAll(testName)
		assert.NoError(t, err)
	})

	t.Run("clone valid repo with provider", func(t *testing.T) {
		err := git.Clone("github.com/"+testRepo, testName)
		assert.NoError(t, err)
		assert.DirExists(t, path.Join(testName))

		err = os.RemoveAll(testName)
		assert.NoError(t, err)
	})

	t.Run("clone valid repo with protocol", func(t *testing.T) {
		err := git.Clone("https://github.com/"+testRepo, testName)
		assert.NoError(t, err)
		assert.DirExists(t, path.Join(testName))

		err = os.RemoveAll(testName)
		assert.NoError(t, err)
	})
}
