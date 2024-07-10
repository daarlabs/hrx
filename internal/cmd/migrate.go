package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
	"github.com/daarlabs/hrx/internal/template"
	"github.com/daarlabs/hrx/internal/util"
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
			migrateUpCmd := exec.Command(
				"/bin/bash", "-c", strings.Join([]string{"go", "run", wd + "/" + migrationsDir + "/*.go", "--up"}, " "),
			)
			stdout, err := migrateUpCmd.StdoutPipe()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			stderr, err := migrateUpCmd.StderrPipe()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			if err := migrateUpCmd.Start(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			go func(stdout io.ReadCloser) {
				defer func(stdout io.ReadCloser) {
					if err := stdout.Close(); err != nil {
						log.Error(err)
						os.Exit(1)
					}
				}(stdout)
				scanner := bufio.NewScanner(stdout)
				for scanner.Scan() {
					log.Success(scanner.Text())
				}
				if err := scanner.Err(); err != nil {
					log.Error(err)
				}
			}(stdout)
			go func(stderr io.ReadCloser) {
				defer func(stderr io.ReadCloser) {
					if err := stderr.Close(); err != nil {
						log.Error(err)
						os.Exit(1)
					}
				}(stderr)
				scanner := bufio.NewScanner(stderr)
				for scanner.Scan() {
					log.Error(errors.New(scanner.Text()))
				}
			}(stderr)
			if err := migrateUpCmd.Wait(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
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
			migrateDownCmd := exec.Command(
				"/bin/bash", "-c", strings.Join([]string{"go", "run", wd + "/" + migrationsDir + "/*.go", "--down"}, " "),
			)
			stdout, err := migrateDownCmd.StdoutPipe()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			stderr, err := migrateDownCmd.StderrPipe()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			if err := migrateDownCmd.Start(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			go func(stdout io.ReadCloser) {
				defer func(stdout io.ReadCloser) {
					if err := stdout.Close(); err != nil {
						log.Error(err)
						os.Exit(1)
					}
				}(stdout)
				scanner := bufio.NewScanner(stdout)
				for scanner.Scan() {
					log.Success(scanner.Text())
				}
				if err := scanner.Err(); err != nil {
					log.Error(err)
				}
			}(stdout)
			go func(stderr io.ReadCloser) {
				defer func(stderr io.ReadCloser) {
					if err := stderr.Close(); err != nil {
						log.Error(err)
						os.Exit(1)
					}
				}(stderr)
				scanner := bufio.NewScanner(stderr)
				for scanner.Scan() {
					log.Error(errors.New(scanner.Text()))
				}
			}(stderr)
			if err := migrateDownCmd.Wait(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
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
	return util.GitAdd(path)
}
