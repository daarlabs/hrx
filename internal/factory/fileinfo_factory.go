package factory

import (
	"fmt"
	"time"
	
	"github.com/iancoleman/strcase"
	
	"github.com/daarlabs/hrx/internal/model"
	"github.com/daarlabs/hrx/internal/module"
	"github.com/daarlabs/hrx/internal/util"
)

func CreateFileInfo(generateType, wd, inputDir, inputName string) model.FileInfo {
	dir := util.NormalizeDir(inputDir)
	r := model.FileInfo{
		Wd:        wd,
		ModuleDir: wd + "/" + dir,
		SnakeName: strcase.ToSnake(inputName),
		KebabName: strcase.ToKebab(inputName),
		CamelName: strcase.ToCamel(inputName),
	}
	modulePath, moduleExists := module.DetectModule(dir)
	if moduleExists {
		r.Module = module.MustGetName(modulePath)
	}
	switch generateType {
	case model.Migration:
		r.Dir = wd + "/" + dir
		if len(inputName) == 0 {
			r.Path = r.Dir + "/" + fmt.Sprintf("%d", time.Now().UnixNano()) + model.GoExtension
		}
		if len(inputName) > 0 {
			r.Path = r.Dir + "/" + fmt.Sprintf("%d_%s", time.Now().UnixNano(), r.SnakeName) + model.GoExtension
		}
	case model.Page:
		r.Package = r.SnakeName + "_" + generateType
		r.Dir = wd + "/" + dir + "/" + model.Handler + "/" + inputName + "_" + model.Handler + "/" + inputName + "_" + generateType
		r.Path = r.Dir + "/" + inputName + "_" + generateType + model.GoExtension
	case model.Props:
		r.Package = r.SnakeName + "_" + model.Page
		r.Dir = wd + "/" + dir + "/" + model.Handler + "/" + inputName + "_" + model.Handler + "/" + inputName + "_" + model.Page
		r.Path = r.Dir + "/" + inputName + "_" + model.Page + "_" + generateType + model.GoExtension
	default:
		r.Package = r.SnakeName + "_" + generateType
		r.Dir = wd + "/" + dir + "/" + generateType + "/" + inputName + "_" + generateType
		r.Path = r.Dir + "/" + inputName + "_" + generateType + model.GoExtension
	}
	return r
}
