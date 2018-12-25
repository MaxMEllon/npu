package semver_test

import (
	"reflect"
	"testing"

	"github.com/maxmellon/nvu/semver"
)

func TestGreaterThan(t *testing.T) {
	suite := []struct {
		left     *semver.Semver
		right    *semver.Semver
		expected bool
	}{
		{
			left:     &semver.Semver{Prefix: "", Major: 1, Miner: 1, Patch: 1, Valid: true, Raw: "1.1.1"},
			right:    &semver.Semver{Prefix: "", Major: 1, Miner: 1, Patch: 2, Valid: true, Raw: "1.1.2"},
			expected: true,
		},
	}

	for _, s := range suite {
		actual := s.left.GreaterThanPatch(s.right)
		if actual != s.expected {
			t.Errorf("expected %#v, but got %#v", s.expected, actual)
		}
	}
}

func TestParse(t *testing.T) {
	suite := []struct {
		data     string
		expected *semver.Semver
	}{
		{
			data:     "1.1.1",
			expected: &semver.Semver{Prefix: "", Major: 1, Miner: 1, Patch: 1, Valid: true, Raw: "1.1.1"},
		},
		{
			data:     "~1.1.1",
			expected: &semver.Semver{Prefix: "~", Major: 1, Miner: 1, Patch: 1, Valid: true, Raw: "~1.1.1"},
		},
		{
			data:     "^100.10000.1000000",
			expected: &semver.Semver{Prefix: "^", Major: 100, Miner: 10000, Patch: 1000000, Valid: true, Raw: "^100.10000.1000000"},
		},
		{
			data:     "23.1",
			expected: &semver.Semver{Prefix: "", Major: 0, Miner: 0, Patch: 0, Valid: false, Raw: "23.1"},
		},
		/* TODOs
		{
			data: "*":
			expected: &semver.Semver{Prefix: "*", Major: -1, Miner: -1, Patch: -1, Valid: true, Raw: "*"},
		},
		{
			data: "22.x":
			expected: &semver.Semver{Prefix: "", Major: 22, Miner: -1, Patch: -1, Valid: true, Raw: "22.x"},
		},
		*/
	}

	for _, s := range suite {
		actual, _ := semver.Parse(s.data)
		if !reflect.DeepEqual(actual, s.expected) {
			t.Errorf("expected %#v, but got %#v", s.expected, actual)
		}
	}
}
