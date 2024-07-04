package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	ComponentFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
)

type %[2]sComponent struct {
	hiro.Component
}

func (c *%[2]sComponent) Name() string {
	return "%[3]s"
}

func (c *%[2]sComponent) Mount() {}

func (c *%[2]sComponent) Node() Node {
	return Div(Text(c.Name() + " component working!"))
}
`
)

func CreateComponentFileTemplate(packageName, componentCamelName, componentKebabName string) string {
	componentCamelName = strings.TrimSuffix(componentCamelName, model.Component)
	componentKebabName = strings.TrimSuffix(componentKebabName, model.LowComponent)
	return fmt.Sprintf(ComponentFileTemplate, packageName, componentCamelName, componentKebabName)
}
