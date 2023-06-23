package versions

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major int
	minor int
	patch int
}

type IVersionExporter interface {
	Load() (Version, error)
	Set(Version) error
}

func NewVersion(version string) (Version, error) {
	v := Version{}

	version = strings.Replace(version, " ", "", -1)
	version = strings.Replace(version, "\n", "", -1)
	version = strings.Replace(version, "\r", "", -1)

	components := strings.Split(version, ".")
	if len(components) != 3 {
		return v, fmt.Errorf("invalid version format")
	}

	major, err := strconv.Atoi(components[0])
	if err != nil {
		return v, fmt.Errorf("invalid major version")
	}
	v.major = major

	minor, err := strconv.Atoi(components[1])
	if err != nil {
		return v, fmt.Errorf("invalid minor version")
	}
	v.minor = minor

	patch, err := strconv.Atoi(components[2])
	if err != nil {
		return v, fmt.Errorf("invalid patch version")
	}
	v.patch = patch

	return v, nil
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func (v Version) IsGreaterThan(other Version) bool {
	if v.major > other.major {
		return true
	}
	if v.major < other.major {
		return false
	}

	if v.minor > other.minor {
		return true
	}
	if v.minor < other.minor {
		return false
	}

	return v.patch > other.patch
}

func (v Version) IsEqual(other Version) bool {
	return v.major == other.major && v.minor == other.minor && v.patch == other.patch
}
