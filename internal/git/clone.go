package git

import (
	"github.com/daarlabs/hrx/internal/util"
)

func Clone(repo, dir string) {
	util.Exec("git", "clone", repo, dir, "--quiet")
}
