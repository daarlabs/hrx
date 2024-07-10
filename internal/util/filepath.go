package util

import (
	"os"
	"strings"
	
	"github.com/iancoleman/strcase"
	
	"github.com/daarlabs/hrx/internal/model"
)

func ParsePath(path string) model.ParsedPath {
	var result model.ParsedPath
	wd, _ := os.Getwd()
	result.Wd = wd
	path = strings.TrimPrefix(path, ".")
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	n := len(parts)
	if n > 0 {
		if n > 1 {
			result.Dir = strings.Join(parts[:n-1], "/")
			result.Package = strings.Join(parts[n-2:n-1], "/")
		}
		if n < 2 {
			result.Package = model.MainPackage
		}
		name := strings.TrimSuffix(parts[n-1], model.GoExtension)
		result.SnakeName = strcase.ToSnake(name)
		result.KebabName = strcase.ToKebab(name)
		result.CamelName = strcase.ToCamel(name)
	}
	return result
}
