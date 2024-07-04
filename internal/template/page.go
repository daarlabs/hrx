package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	PageFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
)

func %[2]sPage(c hiro.Ctx) Node {
	return Div(Text("%[2]s page working!"))
}

`
)

func CreatePageFileTemplate(packageName, pageName string) string {
	pageName = strings.TrimSuffix(pageName, model.Page)
	return fmt.Sprintf(PageFileTemplate, packageName, pageName)
}
