package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Packages - struct as package.json
type Packages struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// Read - read package.json
func Read(path string) (*Packages, error) {
	str, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("failure open file %s, detail: %v", path, err)
	}

	packages := new(Packages)
	err = json.Unmarshal(str, packages)

	if err != nil {
		return nil, fmt.Errorf("failure parse json %s, detail: %v", str, err)
	}

	return packages, nil
}
