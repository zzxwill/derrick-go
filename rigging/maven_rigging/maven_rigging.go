package rigging

import (
	"os"
	"path/filepath"
)

const GolangPlatform = "Python"

func Detect(workspace string) (bool, string) {
	requirementsTxt := filepath.Join(workspace, "requirements.txt")
	setupPy := filepath.Join(workspace, "setup.py")
	if _, err := os.Stat(requirementsTxt); err == nil {
		if _, err := os.Stat(setupPy); err == nil {
			return true, GolangPlatform
		}
	}
	return false, ""
}
