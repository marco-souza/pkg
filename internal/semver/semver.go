// Semantiv Versioning package
package semver

import (
	"fmt"
	"regexp"
	"strconv"
)

type SemVer struct {
	Major  int
	Minor  int
	Patch  int
	prefix string
}

func (s *SemVer) SetVersion(version string) error {
	validSemVerRegex := `^v?(\d+)\.(\d+)\.(\d+)$`
	if !regexp.MustCompile(validSemVerRegex).MatchString(version) {
		return fmt.Errorf("Invalid version string: %s", version)
	}

	groups := regexp.MustCompile(validSemVerRegex).FindStringSubmatch(version)[1:]
	if groups == nil || len(groups) != 3 {
		return fmt.Errorf("Invalid version groups %v", groups)
	}

	var err error
	if s.Major, err = strconv.Atoi(groups[0]); err != nil {
		return fmt.Errorf("Invalid number: %s", groups[0])
	}

	if s.Minor, err = strconv.Atoi(groups[1]); err != nil {
		return fmt.Errorf("Invalid number: %s", groups[1])
	}

	if s.Patch, err = strconv.Atoi(groups[2]); err != nil {
		return fmt.Errorf("Invalid number: %s", groups[2])
	}

	s.prefix = ""
	if version[0] == 'v' {
		s.prefix = "v"
	}

	return nil
}

func (s *SemVer) GetVersion() string {
	return fmt.Sprintf("%s%d.%d.%d", s.prefix, s.Major, s.Minor, s.Patch)
}

func (s *SemVer) BumpVersion(releaseType string) error {
	switch releaseType {
	case "major":
		s.Major++
		s.Minor = 0
		s.Patch = 0
	case "minor":
		s.Minor++
		s.Patch = 0
	case "patch":
		s.Patch++
	default:
		return fmt.Errorf("Invalid release type %s", releaseType)
	}

	return nil
}

func main() {
	fmt.Println("Hello, World!")
}
