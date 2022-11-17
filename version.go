package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type semversion int
type semver string

var (
	str              = func(i int) string { return strconv.Itoa(i) }
	num              = func(s string) (int, error) { return strconv.Atoi(s) }
	ErrInvalidPrefix = errors.New("missing version prefix")
	ErrEmptyVersion  = errors.New("version string is empty")
)

const (
	maj semversion = iota
	min
	patch
	maxVersions = 3
	prefix      = "v"
)

// Version represents a build version with semantic versioning
type version struct {
	major int
	minor int
	patch int
}

// String implements the Stringer interface on semver
// and returns a string representation
// This is very useful when printing Version structs
func (s semver) String() string {
	return string(s)
}

// String implements the Stringer interface on version
// and returns a string representation
// This is very useful when printing Version structs
func (v version) String() string {
	return fmt.Sprintf("%s%d.%d.%d", prefix, v.Major(), v.Minor(), v.Patch())
}

// Major returns a int representation of the major version
func (v version) Major() int {
	return v.major
}

// Minor returns a int representation of the minor version
func (v version) Minor() int {
	return v.minor
}

// Patch returns an int representation of the patch version
func (v version) Patch() int {
	return v.patch
}

// MajorString returns a string representation of the major version
func (v version) MajorString() string {
	return str(v.major)
}

// MinorString returns a string representation of the minor version
func (v version) MinorString() string {
	return str(v.minor)
}

// PatchString returns a string representation of the patch version
func (v version) PatchString() string {
	return str(v.patch)
}

// NewVersion takes a string representation of a semantic version
// and returns a version struct with Major, Minor and Patch fields
func NewVersion(vers string) (*version, error) {
	res := &version{}
	if err := unmarshal(vers, res); err != nil {
		return nil, err
	}
	return res, nil
}

// unmarshal takes a string representation of a semantic version and serializes it
// into a target version struct
func unmarshal(vers string, target *version) error {
	if ok, err := semver(vers).validate(prefix); !ok {
		return fmt.Errorf("validate version prefix: %v", err)
	}
	// Remove prefix
	vers = vers[1:]
	// Should always expect length of 3: maj.min.patch
	out := strings.Split(vers, ".")
	if len(out) < maxVersions || len(out) > maxVersions {
		return fmt.Errorf("version does not conform to semantic version format: %s", vers)
	}

	res := make([]int, 0, maxVersions)
	for i, v := range out {
		if v == "" {
			return fmt.Errorf("unmarshal: %s has empty versions", semversion(i))
		}
		val, err := num(v)
		if err != nil {
			return fmt.Errorf("unmarshal: %s: %v", v, err)
		}
		res = append(res, val)
	}

	target.major = res[maj]
	target.minor = res[min]
	target.patch = res[patch]

	return nil
}

// validate takes a string semantic version and expected target prefix
// and checks whether target prefix matches semantic prefix
func (s semver) validate(target string) (bool, error) {
	if s.String() == "" {
		return false, ErrEmptyVersion
	}
	prefix := s.String()[0]
	if string(prefix) != target {
		return false, ErrInvalidPrefix
	}
	return true, nil
}

// String implements the Stringer interface to return string representations
// of vers enums
func (v semversion) String() string {
	switch v {
	case maj:
		return "major"
	case min:
		return "minor"
	case patch:
		return "patch"
	default:
		return "major"
	}
}
