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
		return c.Response().Render(
			Div(Text("%[2]s handler working!")),
		)
	}
}

`
	
	HandlerTemplate = `func %[1]s() hiro.Handler {
	return func(c hiro.Ctx) error {
		return c.Response().Render(
			Div(Text("%[1]s handler working!")),
		)
	}
}
`
)

func CreateHandlerFileTemplate(packageName, handlerName string) string {
	handlerName = strings.TrimSuffix(handlerName, model.Handler)
	return fmt.Sprintf(HandlerFileTemplate, packageName, handlerName)
}

func CreateHandlerTemplate(handlerName string, index int) string {
	handlerName = strings.TrimSuffix(handlerName, model.Handler)
	handlerName += fmt.Sprintf("%d", index)
	return "\n" + fmt.Sprintf(HandlerTemplate, handlerName)
}
