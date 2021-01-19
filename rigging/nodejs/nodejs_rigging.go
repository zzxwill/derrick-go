package nodejs

import (
	"os"
	"path/filepath"
)

const Platform = "NodeJS"

type NodeJSRigging struct {
}

func (rig NodeJSRigging) Detect(workspace string) (bool, string) {
	packageJSON := filepath.Join(workspace, "package.json")
	if _, err := os.Stat(packageJSON); err == nil {
		return true, Platform
	}
	return false, ""
}
