package semver_test

import (
	"testing"

	"github.com/marco-souza/pkg/internal/semver"
	"github.com/stretchr/testify/assert"
)

func TestSemVer(t *testing.T) {
	s := semver.SemVer{}

	t.Run("Load version", func(t *testing.T) {
		assert.Equal(t, "0.0.0", s.GetVersion())

		err := s.SetVersion("v1.2.3")
		assert.Nil(t, err)

		assert.Equal(t, "v1.2.3", s.GetVersion())
	})

	t.Run("Bump version", func(t *testing.T) {
		s = semver.SemVer{}

		err := s.BumpVersion("major")
		assert.Nil(t, err)
		assert.Equal(t, "1.0.0", s.GetVersion())

		err = s.BumpVersion("minor")
		assert.Nil(t, err)
		assert.Equal(t, "1.1.0", s.GetVersion())

		err = s.BumpVersion("patch")
		assert.Nil(t, err)
		assert.Equal(t, "1.1.1", s.GetVersion())
	})

	t.Run("Invalid version", func(t *testing.T) {
		s = semver.SemVer{}

		err := s.SetVersion("1.2.3")
		assert.Nil(t, err, "")

		err = s.SetVersion("v1.2")
		assert.ErrorContains(t, err, "Invalid")

		err = s.SetVersion("v1.2.3.4")
		assert.ErrorContains(t, err, "Invalid")

		err = s.SetVersion("v1.e.1")
		assert.ErrorContains(t, err, "Invalid")
	})
}
