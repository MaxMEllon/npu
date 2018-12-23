package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client - npm registory api wraped struct
type Client struct {
	registry string
}

type latestVersion struct {
	Version string `json:"version"`
}

type allVersions struct {
	Versions map[string]string `json:"time"`
}

// NewClient - create Clinet{} instance
func NewClient(adapter ...string) *Client {
	client := new(Client)
	if len(adapter) > 0 {
		client.registry = adapter[0]
	} else {
		client.registry = "https://registry.npmjs.org/"
	}
	return client
}

// GetAllVersions = get all versions of package from registry
func (c *Client) GetAllVersions(moduleName string) (map[string]string, error) {
	url := c.registry + moduleName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failure create http request object %v", err)
	}
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failure request to %s. detail: %v", c.registry, err)
	}
	defer resp.Body.Close()
	data := new(allVersions)
	rawString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failure read response data %v", err)
	}
	json.Unmarshal(rawString, data)
	return data.Versions, nil
}

// GetLatest - get latest version of package from registry
func (c *Client) GetLatest(moduleName string) (string, error) {
	url := c.registry + moduleName + "/latest"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failure create http request object %v", err)
	}
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failure request to %s. detail: %v", c.registry, err)
	}
	defer resp.Body.Close()
	data := new(latestVersion)
	rawString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failure read response data %v", err)
	}
	json.Unmarshal(rawString, data)
	return data.Version, nil
}
