package php

import (
	"os"
	"path/filepath"
)

const Platform = "PHP"

type PHPRigging struct {
}

func (rig PHPRigging) Detect(workspace string) (bool, string) {
	composer := filepath.Join(workspace, "composer.json")
	if _, err := os.Stat(composer); err == nil {
		return true, Platform
	}
	return false, ""
}
