package util

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
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

func GitAdd(filepath string) error {
	cmd := exec.Command("git", "add", filepath)
	return cmd.Run()
}
