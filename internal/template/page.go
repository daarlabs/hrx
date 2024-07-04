package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	PageFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
)

func %[2]s(c mirage.Ctx) Node {
	return Div(Text("%[2]s page working!"))
}

`
)

func CreatePageFileTemplate(packageName, pageName string) string {
	pageName = strings.TrimSuffix(pageName, model.Page)
	return fmt.Sprintf(PageFileTemplate, packageName, pageName)
}
