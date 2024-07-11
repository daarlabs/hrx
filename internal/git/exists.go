package git

import (
	"errors"
	"os"
	"path/filepath"
)

func Exists(rootdir string) (bool, error) {
	var exist bool
	if err := filepath.Walk(
		rootdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			if info.Name() == DirName {
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
