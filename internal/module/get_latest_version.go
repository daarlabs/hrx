package module

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type model struct {
	Version string `json:"version"`
}

const (
	endpoint = "https://go.dev/dl/?mode=json"
)

func GetLatestVersion() (string, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer func() { _ = res.Body.Close() }()
	m := make([]model, 0)
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return "", err
	}
	if len(m) == 0 {
		return "", errors.New("no version found")
	}
	version := m[0].Version
	version = strings.TrimPrefix(version, "go")
	version = version[:strings.LastIndex(version, ".")]
	return version, nil
}

func MustGetLatestVersion() string {
	version, err := GetLatestVersion()
	if err != nil {
		panic(err)
	}
	return version
}
