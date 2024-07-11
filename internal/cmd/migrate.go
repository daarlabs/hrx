package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/git"
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
	"github.com/daarlabs/hrx/internal/template"
	"github.com/daarlabs/hrx/internal/util"
	"github.com/daarlabs/hrx/internal/workspace"
)

const (
	migrationsDir = "migrations"
)

var (
	migrateCmd = &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"m"},
		Short:   "Control migrations",
		Long:    "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			_, ok := workspace.DetectWorkspace()
			if ok {
				return
			}
			log.Error(errors.New(message.InvalidWorkspace))
			os.Exit(1)
		},
	}
)

var (
	migrateInitCmd = &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Create migrations folder with main migrator file",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			if err := initMigrator(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

var (
	migrateUpCmd = &cobra.Command{
		Use:     "up",
		Aliases: []string{"u"},
		Short:   "Migrate up",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			util.Exec(
				"/bin/bash",
				"-c",
				strings.Join([]string{"go", "run", wd + "/" + migrationsDir + "/*.go", "--up"}, " "),
			)
		},
	}
)

var (
	migrateDownCmd = &cobra.Command{
		Use:     "down",
		Aliases: []string{"d"},
		Short:   "Migrate down",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			util.Exec(
				"/bin/bash",
				"-c",
				strings.Join([]string{"go", "run", wd + "/" + migrationsDir + "/*.go", "--down"}, " "),
			)
		},
	}
)

func init() {
	migrateCmd.AddCommand(migrateInitCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
}

func initMigrator() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := wd + "/" + migrationsDir
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	if err := ensureModule(dir); err != nil {
		return err
	}
	path := dir + "/main.go"
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if len(fileBytes) == 0 {
			log.Info(fmt.Sprintf("%s: %s", message.AlreadyExistsNoContent, path))
			if err := os.WriteFile(
				path,
				[]byte(template.MigratorFileContent),
				os.ModePerm,
			); err != nil {
				return err
			}
			return nil
		}
		log.Info(fmt.Sprintf("%s: %s", message.AlreadyExists, path))
		return nil
	}
	if err := os.WriteFile(
		path,
		[]byte(template.MigratorFileContent),
		os.ModePerm,
	); err != nil {
		return err
	}
	log.Success(fmt.Sprintf("%s: %s", message.Created, path))
	return git.Add(path)
}
