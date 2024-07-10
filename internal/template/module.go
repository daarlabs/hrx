package template

import (
	"fmt"
)

const (
	ModuleFileTemplate = `module %[1]s

go %[2]s
`
)

func CreateModuleFileTemplate(moduleName, version string) string {
	return fmt.Sprintf(ModuleFileTemplate, moduleName, version)
}
