package platform

import (
	"fmt"
	"os"
	"regexp"
)

type PackageNameDetector struct {
}

func (detector PackageNameDetector) Execute() (map[string]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile("/src/.*")
	if err != nil {
		return nil, err
	}
	result := re.FindAllString(cwd, -1)
	if len(result) == 0 {
		return nil, fmt.Errorf("the source code is not in GOPATH")
	}
	return map[string]string{"project_folder": result[0]}, nil
}

func (detector PackageNameDetector) Name() string {
	return "PackageNameDetector"
}
