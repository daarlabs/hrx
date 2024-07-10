package util

import "strings"

func NormalizeDir(dir string) string {
	dir = strings.TrimPrefix(dir, ".")
	dir = strings.TrimPrefix(dir, "/")
	dir = strings.TrimSuffix(dir, "/")
	return dir
}
