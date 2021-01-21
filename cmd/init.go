package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/cloud-native-application/derrick-go/common"
	"github.com/flosch/pongo2"

	"github.com/cloud-native-application/derrick-go/core"
	"github.com/spf13/cobra"
)

var projectPath string

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"ini"},
		Short:   "Detect application's platform and compile the application",
		Long:    "Detect application's platform and compile the application",
		Example: `derrick init`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(projectPath)
		},
	}
	cmd.Flags().StringP("debug", "d", "", "debug mod")
	cmd.Flags().StringVarP(&projectPath, "project-path", "p", "", "Path of a project which is about to detected its source code ")
	return cmd
}

type SuitableRiggings struct {
	Platform       string
	ExtensionPoint core.ExtensionPoint
}

func execute(workspace string) error {
	suitableRiggings := detect(workspace)
	switch len(suitableRiggings) {
	case 0:
		fmt.Println("Failed to detect your application's platform.\nMaybe you can upgrade Derrick to get more platforms supported.")
		return nil
	case 1:
		fmt.Println(fmt.Sprintf("Derrick detect your platform is %s and compile it successfully.", suitableRiggings[0].Platform))
	default:
		// TODO(zzxwill) ask users to choose from one of them
		fmt.Println("More than one rigging can handle the application.")
		return nil
	}
	rig := suitableRiggings[0].ExtensionPoint.Rigging
	detectedContext, err := rig.Compile()
	if err != nil {
		return err
	}
	fmt.Println(detectedContext)
	destDir := filepath.Join(workspace, ".derrick")
	if _, err := os.Stat(destDir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, 0750); err != nil {
			return err
		}
	}
	if err := renderTemplates(rig, detectedContext, destDir); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

func detect(projectPath string) []*SuitableRiggings {
	allRigging := core.LoadRiggings()
	if projectPath == "" {
		projectPath = "./"
	}
	var suitableRiggings []*SuitableRiggings
	for _, rig := range allRigging {
		success, platform := rig.Rigging.Detect(projectPath)
		if success {
			suitableRiggings = append(suitableRiggings,
				&SuitableRiggings{
					Platform:       platform,
					ExtensionPoint: core.ExtensionPoint{Rigging: rig.Rigging},
				})
		}
	}
	return suitableRiggings
}

func renderTemplates(rig common.Rigging, detectedContext map[string]string, destDir string) error {
	// TODO(zzxwill) PkgPath() returns github.com/cloud-native-application/derrick-go/rigging/golang/templates
	// there might be a better solution get the direcotry of the templates
	pkgPath := strings.Join(strings.Split(reflect.TypeOf(rig).PkgPath(), "/")[3:], "/")
	templateDir := filepath.Join(filepath.Clean(pkgPath), "templates")
	var templates []string
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if info != nil && strings.HasSuffix(info.Name(), ".j2") {
			templates = append(templates, info.Name())
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, t := range templates {
		renderedTemplate, err := renderTemplate(filepath.Join(templateDir, t), detectedContext)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(filepath.Join(destDir, t), []byte(renderedTemplate), 0750); err != nil {
			return err
		}
	}
	return nil
}

func renderTemplate(templatePath string, detectedContext map[string]string) (string, error) {
	data, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	tpl, err := pongo2.FromString(string(data))
	if err != nil {
		return "", err
	}
	ctx := make(pongo2.Context, len(detectedContext))
	for k, v := range detectedContext {
		ctx[k] = v
	}

	out, err := tpl.Execute(ctx)
	if err != nil {
		panic(err)
	}
	return out, nil
}
