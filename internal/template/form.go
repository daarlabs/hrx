package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	FormFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/arcanum/form"
	"github.com/daarlabs/arcanum/mirage"
)

type %[2]s struct {
	Form
	Change  Field[string]
}

func Create%[2]s(c mirage.Ctx) (%[2]s, error) {
	f := c.Create().Form(
		Add("change").Id("").Label(c.Translate("")).With(Text(), Validate.Required()),
	)
	return Build[%[2]s](f)
}

func MustCreate%[2]s(c mirage.Ctx) %[2]s {
	f, err := Create%[2]s(c)
	if err != nil {
		panic(err)
	}
	return f
}
`
)

func CreateFormFileTemplate(packageName, formName string) string {
	formName = strings.TrimSuffix(formName, model.Form)
	return fmt.Sprintf(FormFileTemplate, packageName, formName)
}
