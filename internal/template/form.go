package template

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/hrx/internal/model"
)

const (
	FormFileTemplate = `package %[1]s

import (
	. "github.com/daarlabs/hirokit/form"
	"github.com/daarlabs/hirokit/hiro"
)

type %[2]sForm struct {
	Form
	Change  Field[string]
}

func Create%[2]sForm(c hiro.Ctx) (%[2]sForm, error) {
	f := c.Create().Form(
		Add("change").Id("").Label(c.Translate("")).With(Text(), Validate.Required()),
	)
	return Build[%[2]sForm](f)
}

func MustCreate%[2]sForm(c hiro.Ctx) %[2]sForm {
	f, err := Create%[2]sForm(c)
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
