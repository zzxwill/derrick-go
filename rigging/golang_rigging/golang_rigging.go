package rigging

import (
	"os"
	"path/filepath"
)

const PythonPlatform = "Python"

//type PythonRigging struct {
//	Platform string
//}
//
//func (python PythonRigging) GetTemplateDir() string {
//	return reflect.TypeOf(PythonPlatform).PkgPath()
//}
//
//func (python PythonRigging) GetTemplateName() string {
//	return PythonPlatform
//}

func Detect(workspace string) (bool, string) {
	requirementsTxt := filepath.Join(workspace, "requirements.txt")
	setupPy := filepath.Join(workspace, "setup.py")
	if _, err := os.Stat(requirementsTxt); err == nil {
		if _, err := os.Stat(setupPy); err == nil {
			return true, PythonPlatform
		}
	}
	return false, ""
}
