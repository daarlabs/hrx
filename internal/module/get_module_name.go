package module

import (
	"os"
	"strings"
)

func GetName(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(fileBytes), "\n")
	for _, line := range parts {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", nil
}

func MustGetName(path string) string {
	name, err := GetName(path)
	if err != nil {
		panic(err)
	}
	return name
}
