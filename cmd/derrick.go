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
	fmt.Println(core.DerrickLogo)
	if err := core.InitDirs(); err != nil {
		return fmt.Errorf("failed to init Derrick home, err: %s", err)
	}
	return nil
}

func load() error {
	firstTimeFlag, err := core.CheckDerrickFirstSetup()
	if err != nil {
		return err
	} else if firstTimeFlag {
		if err := preLoad(); err != nil {
			return err
		}
	}
	return nil
}
