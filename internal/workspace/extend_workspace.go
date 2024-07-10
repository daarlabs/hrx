package workspace

import (
	"fmt"
	"os"
	"strings"
)

func ExtendWorkspace(path, module string) error {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(fileBytes)
	content = strings.Replace(content, ")", fmt.Sprintf("\t%s\n)", module), 1)
	return os.WriteFile(path, []byte(content), os.ModePerm)
}
