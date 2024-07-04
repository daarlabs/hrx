package util

import (
	"strings"
	
	"github.com/iancoleman/strcase"
	
	"github.com/daarlabs/hrx/internal/model"
)

func ParseFilePath(path string) model.FilePath {
	var result model.FilePath
	path = strings.TrimPrefix(path, ".")
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	n := len(parts)
	if n > 0 {
		if n > 1 {
			result.Path = strings.Join(parts[:n-1], "/")
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
