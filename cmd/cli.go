package cmd

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/cloud-native-application/derrick-go/core"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

// Commands will contain all commands
func Commands() *cobra.Command {
	// ioStream := util.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	cmd := &cobra.Command{
		Use:                "derrick-go",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("🐳 A tool to help you containerize application in seconds\n\n")
			cmd.Println("Flags:")
			cmd.Println("  -h, --help   help for derrick")
			cmd.Println()
			cmd.Println(`Use "derrick [command] --help" for more information about a command.`)
		},
		SilenceUsage: true,
	}

	cmd.AddCommand(
		// Getting Start

		NewVersionCommand(),
	)

	// this is for mute klog
	fset := flag.NewFlagSet("logs", flag.ContinueOnError)
	klog.InitFlags(fset)
	_ = fset.Set("v", "-1")

	return cmd
}

// PrintHelpByTag print custom defined help message
func PrintHelpByTag(cmd *cobra.Command, all []*cobra.Command, tag string) {
	cmd.Printf("  %s:\n\n", tag)
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true

	cmd.Println(table.String())
	cmd.Println()
}

// NewVersionCommand print client version
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints out build version information",
		Long:  "Prints out build version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`Version: %v
GitRevision: %v
GolangVersion: %v
`,
				core.Version,
				core.GitRevision,
				runtime.Version())
		},
	}
}
