package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	HandlerFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
)

func %[2]s() hiro.Handler {
	return func(c hiro.Ctx) error {
		c.Page().Set().Title("%[2]s")
		return c.Response().Render(
			Div(Text("%[2]s handler working!")),
		)
	}
}
`
	HandlerPageFileTemplate = `package %[1]s

import (
	"%[3]s"
	"github.com/daarlabs/hirokit/hiro"
)

func %[2]s() hiro.Handler {
	return func(c hiro.Ctx) error {
		c.Page().Set().Title("%[5]s")
		return c.Response().Render(
			%[4]s.%[5]sPage(c, %[4]s.Props{}),
		)
	}
}
`
)

func CreateHandlerFileTemplate(packageName, handlerName string) string {
	handlerName = strings.TrimSuffix(handlerName, model.Handler)
	return fmt.Sprintf(HandlerFileTemplate, packageName, handlerName)
}

func CreateHandlerPageFileTemplate(handlerInfo, pageInfo model.FileInfo) string {
	pagePath := strings.TrimPrefix(pageInfo.Dir, handlerInfo.Wd)
	pagePath = pagePath[strings.Index(pagePath, pageInfo.Module):]
	return fmt.Sprintf(
		HandlerPageFileTemplate,
		handlerInfo.Package,
		handlerInfo.CamelName,
		pagePath,
		pageInfo.Package,
		pageInfo.CamelName,
	)
}
