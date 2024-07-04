package cmd

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
	"github.com/daarlabs/hrx/internal/model"
	"github.com/daarlabs/hrx/internal/template"
)

func generateAnotherHandlerToExistingOne(path, filePath, camelName string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := os.WriteFile(
		filePath,
		[]byte(string(fileBytes)+template.CreateHandlerTemplate(
			camelName, strings.Count(string(fileBytes), model.MirageHandlerType)+1,
		)),
		os.ModePerm,
	); err != nil {
		return err
	}
	log.Success(fmt.Sprintf("%s: %s", message.ExistingAdded, path))
	return nil
}
