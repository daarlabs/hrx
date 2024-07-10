package template

const (
	MigratorFileContent = `package main

import (
	"github.com/daarlabs/hirokit/esquel"
	"github.com/daarlabs/hirokit/esquel/migrator"
)

var manager = new(migrator.Manager)

func main() {
	m := migrator.New(
		".",
		map[string]*esquel.DB{},
		manager.GetAll(),
	)
	m.MustRun()
}
`
)

const (
	MigrationFileContent = `package main

import "github.com/daarlabs/hirokit/esquel/migrator"

func init() {
	manager.Add().
		Up(
			func(c migrator.Control) {

			},
		).
		Down(
			func(c migrator.Control) {
			
			},
		)
}
`
)
