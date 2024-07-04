package cmd

import (
	"errors"
	"fmt"
	"os"
	
	"github.com/daarlabs/hrx/internal/config"
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
	"github.com/daarlabs/hrx/internal/model"
	"github.com/daarlabs/hrx/internal/template"
	"github.com/daarlabs/hrx/internal/util"
)

func generate(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	gitExists, err := util.GitExists(wd)
	if err != nil {
		return err
	}
	parsed := util.ParseFilePath(path)
	fullpath := wd + "/" + parsed.Path
	if err := os.MkdirAll(fullpath, os.ModePerm); err != nil {
		return err
	}
	filePath := fullpath + "/" + parsed.SnakeName + model.GoExtension
	generatorType := getGenerateType()
	if len(generatorType) == 0 {
		return errors.New(message.InvalidGenerateType)
	}
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Info(fmt.Sprintf("%s: %s", message.AlreadyExists, path))
		switch generatorType {
		case model.Handler:
			if err := generateAnotherHandlerToExistingOne(path, filePath, parsed.CamelName); err != nil {
				return err
			}
		}
		return nil
	}
	if err := os.WriteFile(
		filePath,
		[]byte(createTemplate(generatorType, parsed)),
		os.ModePerm,
	); err != nil {
		return err
	}
	log.Success(fmt.Sprintf("%s: %s", message.Created, path))
	if gitExists {
		if err := util.GitAdd(filePath); err != nil {
			return err
		}
		log.Success(fmt.Sprintf("%s: %s", message.GitAdded, path))
	}
	return nil
}

func getGenerateType() string {
	if config.Config.Handler {
		return model.Handler
	}
	if config.Config.Form {
		return model.Form
	}
	if config.Config.Page {
		return model.Page
	}
	if config.Config.Component {
		return model.Component
	}
	return ""
}

func createTemplate(generateType string, parsed model.FilePath) string {
	switch generateType {
	case model.Handler:
		return template.CreateHandlerFileTemplate(parsed.Package, parsed.CamelName)
	case model.Form:
		return template.CreateFormFileTemplate(parsed.Package, parsed.CamelName)
	case model.Page:
		return template.CreatePageFileTemplate(parsed.Package, parsed.CamelName)
	case model.Component:
		return template.CreateComponentFileTemplate(parsed.Package, parsed.CamelName, parsed.KebabName)
	default:
		return ""
	}
}
