package golang

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloud-native-application/derrick-go/common"

	"github.com/cloud-native-application/derrick-go/detectors/general"
	"github.com/cloud-native-application/derrick-go/detectors/image"
	platform "github.com/cloud-native-application/derrick-go/detectors/platform/golang"
)

const (
	Platform             = "Golang"
	Meta                 = "Meta"
	Dockerfile           = "Dockerfile.j2"
	Jenkinsfile          = "Jenkinsfile.j2"
	DockerCompose        = "docker-compose.yml.j2"
	KubernetesDeployment = "kubernetes-deployment.yaml.j2"
)

type GolangRigging struct {
}

func (rig GolangRigging) Detect(workspace string) (bool, string) {
	var detected bool
	err := filepath.Walk(workspace, func(workspace string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			detected = true
			return io.EOF
		}
		return nil
	})
	if err == io.EOF && detected {
		return true, Platform
	}
	return false, ""
}

func (rig GolangRigging) Compile() (map[string]string, error) {
	dr := &common.DetectorReport{
		Nodes: map[string]common.DetectorReport{},
		Store: map[string]string{},
	}
	if err := dr.RegisterDetector(general.ImageRepoDetector{}, Meta); err != nil {
		return nil, err
	}
	if err := dr.RegisterDetector(image.GolangVersionDetector{}, Dockerfile); err != nil {
		return nil, err
	}
	if err := dr.RegisterDetector(platform.PackageNameDetector{}, Dockerfile); err != nil {
		return nil, err
	}

	if err := dr.RegisterDetector(general.ImageRepoDetector{}, Jenkinsfile); err != nil {
		return nil, err
	}

	if err := dr.RegisterDetector(general.ImageRepoDetector{}, DockerCompose); err != nil {
		return nil, err
	}

	if err := dr.RegisterDetector(general.ImageRepoDetector{}, KubernetesDeployment); err != nil {
		return nil, err
	}
	if err := dr.RegisterDetector(general.DerrickDetector{}, KubernetesDeployment); err != nil {
		return nil, err
	}
	return dr.GenerateReport(), nil
}
