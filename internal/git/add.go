package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
)

func Add(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	gitExists, err := Exists(wd)
	if err != nil {
		return err
	}
	if gitExists {
		cmd := exec.Command("git", "add", path)
		if err := cmd.Run(); err != nil {
			return err
		}
		log.Success(fmt.Sprintf("%s: %s", message.GitAdded, "."+strings.TrimPrefix(path, wd)))
	}
	return nil
}
