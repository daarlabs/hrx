package module

import (
	"os"
)

func DetectModule(dir string) (string, bool) {
	modulePath := dir + "/" + "go.mod"
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		return "", false
	}
	return modulePath, true
}
