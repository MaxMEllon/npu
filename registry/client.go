package registry

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pkg/errors"
)

// Client - npm registry api wrapped struct
type Client struct {
	registry *url.URL
}

type latestVersion struct {
	Version string `json:"version"`
}

type allVersions struct {
	Versions map[string]string `json:"time"`
}

var DefaultRegistryURL = "https://registry.npmjs.org"

// NewClient - create Client{} instance
func NewClient(adapter ...string) (*Client, error) {
	rawURL := DefaultRegistryURL
	if len(adapter) > 0 {
		rawURL = adapter[0]
	}
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse registry url: %v", rawURL)
	}
	return &Client{registry: u}, nil
}

func (c *Client) get(ctx context.Context, p string, dest interface{}) error {
	u := *c.registry
	u.Path = path.Join(u.Path, p)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to create http request")
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to request registry")
	}
	defer resp.Body.Close()
	defer io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("unexpected status code. expected: %v, actual: %v", http.StatusOK, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return errors.Wrap(err, "unexpected response body")
	}
	return nil
}

// GetAllVersions = get all versions of package from registry
func (c *Client) GetAllVersions(moduleName string) (map[string]string, error) {
	var data allVersions
	if err := c.get(context.TODO(), moduleName, &data); err != nil {
		return nil, err
	}
	return data.Versions, nil
}

// GetLatest - get latest version of package from registry
func (c *Client) GetLatest(moduleName string) (string, error) {
	var data latestVersion
	if err := c.get(context.TODO(), path.Join(moduleName, "latest"), &data); err != nil {
		return "", err
	}
	return data.Version, nil
}
