package cmd

import (
	"fmt"
	"os"

	"github.com/cloud-native-application/derrick-go/core"
)

func Run() {
	if err := load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command := Commands()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func preLoad() error {
	fmt.Print(core.DerrickLogo)
	if err := core.InitDirs(); err != nil {
		return fmt.Errorf("failed to init Derrick home, err: %s", err)
	}
	return nil
}

func load() error {
	firstTimeFlag, _ := core.CheckDerrickFirstSetup()
	if firstTimeFlag {
		if err := preLoad(); err != nil {
			return err
		}
	}

	core.LoadRiggings()
	return nil
}
