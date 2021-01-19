package golang

import (
	"os"
	"path/filepath"
)

const Platform = "Golang"

type GolangRigging struct {
}

func (rig GolangRigging) Detect(workspace string) (bool, string) {
	gofile := filepath.Join(workspace, ".go")
	if _, err := os.Stat(gofile); err == nil {
		return true, Platform
	}
	return false, ""
}
