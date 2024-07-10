package template

import (
	"fmt"
)

const (
	PropsFileTemplate = `package %[1]s

type Props struct {
}

`
)

func CreatePropsFileTemplate(packageName string) string {
	return fmt.Sprintf(PropsFileTemplate, packageName)
}
