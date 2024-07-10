package workspace

import (
	"os"
)

func DetectWorkspace() (string, bool) {
	wd, err := os.Getwd()
	if err != nil {
		return "", false
	}
	workspacePath := wd + "/" + "/go.work"
	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		return "", false
	}
	return workspacePath, true
}
