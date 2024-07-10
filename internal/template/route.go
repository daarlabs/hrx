package template

import (
	"fmt"
)

const (
	RouteFileTemplate = `package %[1]s

import (
	"net/http"
	
	"%[2]s/handler/%[3]s_handler"
	
	. "github.com/daarlabs/hirokit/hiro"
)

func %[4]sRoute(app Hiro) {
	app.Route(
		"/%[3]s",
		%[3]s_handler.%[4]s(),
		Method(http.MethodGet, http.MethodPost),
		Name("%[3]s"),
	)
}

`
)

func CreateRouteFileTemplate(packageName, moduleName, handlerPackageName, handlerName string) string {
	return fmt.Sprintf(
		RouteFileTemplate,
		packageName,
		moduleName,
		handlerPackageName,
		handlerName,
	)
}
