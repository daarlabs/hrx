package cmd

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/config"
	"github.com/daarlabs/hrx/internal/git"
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
)

const (
	defaultName = "starter"
)

var (
	newCmd = &cobra.Command{
		Use:     "new",
		Aliases: []string{"n"},
		Short:   "Create new project",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Error(errors.New(message.InvalidName))
				os.Exit(1)
			}
			log.Success("Creating new project - " + args[0])
			wd, err := os.Getwd()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			newProjectDir := filepath.Join(wd, args[0])
			newProjectGitDir := filepath.Join(wd, args[0], git.DirName)
			if _, err := os.Stat(newProjectDir); !os.IsNotExist(err) {
				log.Error(errors.New(message.AlreadyExists))
				os.Exit(1)
			}
			if err := os.MkdirAll(newProjectDir, os.ModePerm); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			git.Clone(config.StarterRepo, newProjectDir)
			if err := os.RemoveAll(newProjectGitDir); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			if err := filepath.Walk(
				newProjectDir, func(path string, info fs.FileInfo, err error) error {
					if !strings.HasSuffix(info.Name(), ".go") &&
						!strings.HasSuffix(info.Name(), ".yaml") &&
						!strings.HasSuffix(info.Name(), ".md") &&
						!strings.HasSuffix(info.Name(), "Tiltfile") {
						return nil
					}
					fileBytes, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					fileContent := string(fileBytes)
					if !strings.Contains(fileContent, defaultName) {
						return nil
					}
					if err := os.WriteFile(
						path,
						[]byte(strings.ReplaceAll(fileContent, defaultName, args[0])),
						os.ModePerm,
					); err != nil {
						return err
					}
					return nil
				},
			); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			log.Success("------------------")
			log.Success("All done!")
			log.Success("------------------")
			log.Success(" - cd " + args[0])
			log.Success(" - run `tilt up` OR `docker compose up`")
		},
	}
)
