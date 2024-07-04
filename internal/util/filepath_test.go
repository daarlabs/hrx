package util

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestParseFilePath(t *testing.T) {
	t.Run(
		"basic", func(t *testing.T) {
			r := ParseFilePath("app/test/handler/foo-bar")
			assert.Equal(t, "app/test/handler", r.Path)
			assert.Equal(t, "handler", r.Package)
			assert.Equal(t, "FooBar", r.CamelName)
			assert.Equal(t, "foo-bar", r.KebabName)
		},
	)
	t.Run(
		"short", func(t *testing.T) {
			r := ParseFilePath("foo-bar")
			assert.Equal(t, "", r.Path)
			assert.Equal(t, "main", r.Package)
			assert.Equal(t, "FooBar", r.CamelName)
			assert.Equal(t, "foo-bar", r.KebabName)
		},
	)
	t.Run(
		"path equal to package", func(t *testing.T) {
			r := ParseFilePath("./handler/foo-bar")
			assert.Equal(t, "handler", r.Path)
			assert.Equal(t, "handler", r.Package)
			assert.Equal(t, "FooBar", r.CamelName)
			assert.Equal(t, "foo-bar", r.KebabName)
		},
	)
}
