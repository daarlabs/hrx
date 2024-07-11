package module

import (
	"os"
	"strings"
	
	"github.com/daarlabs/hrx/internal/git"
	"github.com/daarlabs/hrx/internal/template"
)

func CreateModule(dir string) (string, error) {
	parts := strings.Split(dir, "/")
	version, err := GetLatestVersion()
	if err != nil {
		return "", err
	}
	name := parts[len(parts)-1]
	modulePath := dir + "/go.mod"
	if err := os.WriteFile(
		modulePath,
		[]byte(template.CreateModuleFileTemplate(name, version)),
		os.ModePerm,
	); err != nil {
		return "", err
	}
	return name, git.Add(modulePath)
}

func MustCreateModule(dir string) string {
	name, err := CreateModule(dir)
	if err != nil {
		panic(err)
	}
	return name
}
