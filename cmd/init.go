package cmd

import (
	"fmt"

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
			suitableRiggings := detect(projectPath)
			switch len(suitableRiggings) {
			case 0:
				fmt.Println("Failed to detect your application's platform.\nMaybe you can upgrade Derrick to get more platforms supported.")
				return nil
			case 1:
				fmt.Println(fmt.Sprintf("Derrick detect your platform is %s and compile it successfully.", suitableRiggings[0].Platform))
				return nil
			default:
				// TODO(zzxwill) ask users to choose from one of them
				fmt.Println("More than one rigging can handle the application.")
			}
			return fmt.Errorf("unknown issues happened")
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
