package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/factory"
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
	"github.com/daarlabs/hrx/internal/model"
	"github.com/daarlabs/hrx/internal/module"
	"github.com/daarlabs/hrx/internal/template"
	"github.com/daarlabs/hrx/internal/util"
	"github.com/daarlabs/hrx/internal/workspace"
)

var (
	generateName, generateDir, generateApp string
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate project files",
		Long:    "",
	}
)

var (
	generateComponentCmd = &cobra.Command{
		Use:     "component",
		Aliases: []string{"c"},
		Short:   "Generate smart component",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			verifyFlags()
			if err := generate(model.Component, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	generateFormCmd = &cobra.Command{
		Use:     "form",
		Aliases: []string{"f"},
		Short:   "Generate logic form",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			verifyFlags()
			if err := generate(model.Form, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	generateHandlerCmd = &cobra.Command{
		Use:     "handler",
		Aliases: []string{"h"},
		Short:   "Generate handler",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			verifyFlags()
			if err := generate(model.Handler, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	generatePageCmd = &cobra.Command{
		Use:     "page",
		Aliases: []string{"p"},
		Short:   "Generate page (view) with props",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			verifyFlags()
			if err := generate(model.Page, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	generateRouteCmd = &cobra.Command{
		Use:     "route",
		Aliases: []string{"r"},
		Short:   "Generate route",
		Long:    "Includes route (endpoint), handler and page with props, all in one command",
		Run: func(cmd *cobra.Command, args []string) {
			verifyFlags()
			if err := generate(model.Route, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	generateMigrationCmd = &cobra.Command{
		Use:     "migration",
		Aliases: []string{"m"},
		Short:   "Generate new migration",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(generateDir) == 0 {
				generateDir = "./" + migrationsDir
			}
			if err := generate(model.Migration, generateDir, generateName); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	generateCmd.PersistentFlags().StringVarP(&generateName, "name", "n", "", "")
	generateCmd.PersistentFlags().StringVarP(&generateDir, "dir", "d", "", "")
	generateCmd.PersistentFlags().StringVarP(&generateApp, "app", "a", "", "")
	generateCmd.AddCommand(generateComponentCmd)
	generateCmd.AddCommand(generateFormCmd)
	generateCmd.AddCommand(generateHandlerCmd)
	generateCmd.AddCommand(generatePageCmd)
	generateCmd.AddCommand(generateRouteCmd)
	generateCmd.AddCommand(generateMigrationCmd)
}

func verifyFlags() {
	if len(generateName) == 0 {
		log.Error(errors.New(message.InvalidName))
		os.Exit(1)
	}
	if len(generateDir) == 0 && len(generateApp) == 0 {
		log.Error(errors.New(message.InvalidPath))
		os.Exit(1)
	}
	if len(generateApp) > 0 {
		generateDir = "./app/" + generateApp
	}
}

func generate(generateType, inputDir, inputName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := wd + "/" + util.NormalizeDir(inputDir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	if err := ensureModule(dir); err != nil {
		return err
	}
	switch generateType {
	case model.Route:
		if err := createFile(model.Page, factory.CreateFileInfo(model.Page, wd, inputDir, inputName)); err != nil {
			return err
		}
		if err := createFile(model.Props, factory.CreateFileInfo(model.Props, wd, inputDir, inputName)); err != nil {
			return err
		}
		if err := createFile(
			model.HandlerPage,
			factory.CreateFileInfo(model.Handler, wd, inputDir, inputName),
		); err != nil {
			return err
		}
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	case model.Handler:
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	case model.Form:
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	case model.Page:
		if err := createFile(model.Props, factory.CreateFileInfo(model.Props, wd, inputDir, inputName)); err != nil {
			return err
		}
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	case model.Component:
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	case model.Migration:
		return createFile(generateType, factory.CreateFileInfo(generateType, wd, inputDir, inputName))
	default:
		return errors.New(message.InvalidGenerateType)
	}
}

func createFile(generateType string, info model.FileInfo) error {
	if err := os.MkdirAll(info.Dir, os.ModePerm); err != nil {
		return err
	}
	if _, err := os.Stat(info.Path); !os.IsNotExist(err) {
		fileBytes, err := os.ReadFile(info.Path)
		if err != nil {
			return err
		}
		if len(fileBytes) == 0 {
			log.Info(fmt.Sprintf("%s: %s", message.AlreadyExistsNoContent, info.Path))
			if err := os.WriteFile(
				info.Path,
				[]byte(createTemplate(generateType, info)),
				os.ModePerm,
			); err != nil {
				return err
			}
			return nil
		}
		log.Info(fmt.Sprintf("%s: %s", message.AlreadyExists, info.Path))
		return nil
	}
	if err := os.WriteFile(
		info.Path,
		[]byte(createTemplate(generateType, info)),
		os.ModePerm,
	); err != nil {
		return err
	}
	log.Success(fmt.Sprintf("%s: %s", message.Created, info.Path))
	return util.GitAdd(info.Path)
}

func createTemplate(generateType string, info model.FileInfo) string {
	switch generateType {
	case model.Route:
		return template.CreateRouteFileTemplate(info.Package, info.Module, info.SnakeName, info.CamelName)
	case model.Handler:
		return template.CreateHandlerFileTemplate(info.Package, info.CamelName)
	case model.HandlerPage:
		return template.CreateHandlerPageFileTemplate(
			info,
			factory.CreateFileInfo(model.Page, info.Wd, strings.TrimPrefix(info.ModuleDir, info.Wd), info.SnakeName),
		)
	case model.Form:
		return template.CreateFormFileTemplate(info.Package, info.CamelName)
	case model.Page:
		return template.CreatePageFileTemplate(info.Package, info.CamelName)
	case model.Props:
		return template.CreatePropsFileTemplate(info.Package)
	case model.Component:
		return template.CreateComponentFileTemplate(info.Package, info.CamelName, info.KebabName)
	case model.Migration:
		return template.MigrationFileContent
	default:
		return ""
	}
}

func ensureModule(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	_, moduleOk := module.DetectModule(dir)
	if !moduleOk {
		moduleName, err := module.CreateModule(dir)
		if err != nil {
			return err
		}
		workspacePath, workspaceOk := workspace.DetectWorkspace()
		if workspaceOk {
			if err := workspace.ExtendWorkspace(workspacePath, moduleName); err != nil {
				return err
			}
		}
	}
	return nil
}
