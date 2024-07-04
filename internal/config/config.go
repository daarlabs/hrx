package config

import "github.com/daarlabs/hrx/internal/model"

type config struct {
	model.FilePath
	Advice    bool
	Generate  bool
	Handler   bool
	Component bool
	Form      bool
	Page      bool
}

var (
	Config = config{}
)
