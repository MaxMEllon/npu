package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/maxmellon/npu/semver"
)

type packages struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// Packages - struct as package.json
type Packages struct {
	Modules map[string]*semver.Semver
}

// Read - read package.json
func Read(path string) (*Packages, error) {
	str, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("failure open file %s, detail: %v", path, err)
	}

	data := new(packages)
	err = json.Unmarshal(str, data)

	if err != nil {
		return nil, fmt.Errorf("failure parse json %s, detail: %v", str, err)
	}

	result := &Packages{
		Modules: make(map[string]*semver.Semver),
	}

	for k, v := range data.Dependencies {
		dt, _ := semver.Parse(v)
		result.Modules[k] = dt
	}

	for k, v := range data.DevDependencies {
		dt, _ := semver.Parse(v)
		result.Modules[k] = dt
	}

	return result, nil
}
