package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
)

const (
	gitDirName = ".git"
)

func GitExists(rootdir string) (bool, error) {
	var exist bool
	if err := filepath.Walk(
		rootdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			if info.Name() == gitDirName {
				exist = true
				return filepath.SkipDir
			}
			return nil
		},
	); err != nil && !errors.Is(err, filepath.SkipDir) {
		return false, err
	}
	return exist, nil
}

func GitAdd(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	gitExists, err := GitExists(wd)
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
