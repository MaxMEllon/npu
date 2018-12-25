package semver

import (
	"fmt"
	"strconv"
	"strings"
)

// Semver -
type Semver struct {
	Prefix string
	Major  int
	Miner  int
	Patch  int
	Valid  bool
	Raw    string
}

// Parse - parse string as semver
func Parse(version string) (*Semver, error) {
	var semver = new(Semver)
	semver.Raw = version
	if isValidSemver(version) {
		semver.Valid = true
		splited := strings.Split(version, ".")
		if strings.HasPrefix(splited[0], "^") {
			semver.Prefix = "^"
		}
		if strings.HasPrefix(splited[0], "~") {
			semver.Prefix = "~"
		}
		if semver.Prefix != "" {
			major, err := strconv.Atoi(strings.Join(strings.Split(splited[0], "")[1:], ""))
			if err != nil {
				return &Semver{Valid: false, Raw: version}, fmt.Errorf("failure parse 'major' %v", err)
			}
			semver.Major = major
		} else {
			major, err := strconv.Atoi(splited[0])
			if err != nil {
				return &Semver{Valid: false, Raw: version}, fmt.Errorf("failure parse 'major' %v", err)
			}
			semver.Major = major
		}
		miner, err := strconv.Atoi(splited[1])
		if err != nil {
			return &Semver{Valid: false, Raw: version}, fmt.Errorf("failure parse 'miner' %v", err)
		}
		semver.Miner = miner
		patch, err := strconv.Atoi(splited[2])
		if err != nil {
			return &Semver{Valid: false, Raw: version}, fmt.Errorf("failure parse 'patch' %v", err)
		}
		semver.Patch = patch
	} else {
		semver.Valid = false
	}
	return semver, nil
}

// Eq - compare a other *Semver struct
func (s *Semver) Eq(target *Semver) bool {
	return s.Major == target.Major && s.Miner == target.Miner && s.Patch == target.Patch
}

// StrictEq - strict compare a other *Semver struct, actual compares Valid, Raw, and Prefix
func (s *Semver) StrictEq(target *Semver) bool {
	return s.Eq(target) && s.Valid == target.Valid && s.Prefix == target.Prefix && s.Raw == target.Raw
}

// GreaterThanPatch - 3.0.1 < 3.0.2 => true
func (s *Semver) GreaterThanPatch(target *Semver) bool {
	return s.Major == target.Major && s.Miner == target.Miner && s.Patch < target.Patch
}

func isValidSemver(version string) bool {
	return len(strings.Split(version, ".")) == 3
}
