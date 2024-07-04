package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	ComponentFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
)

type %[2]s struct {
	mirage.Component
}

func (c *%[2]s) Name() string {
	return "%[3]s"
}

func (c *%[2]s) Mount() {}

func (c *%[2]s) Node() Node {
	return Div(Text(c.Name() + " component working!"))
}
`
)

func CreateComponentFileTemplate(packageName, componentCamelName, componentKebabName string) string {
	componentCamelName = strings.TrimSuffix(componentCamelName, model.Component)
	componentKebabName = strings.TrimSuffix(componentKebabName, model.LowComponent)
	return fmt.Sprintf(ComponentFileTemplate, packageName, componentCamelName, componentKebabName)
}
