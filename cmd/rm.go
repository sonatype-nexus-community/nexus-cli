package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// RmCommand is the noun which handles any Nexus Repository actions
	RmCommand = &cobra.Command{
		Use:   "rm",
		Short: "command for managing functionality of Repository Manager.",
		Long:  `command for managing functionality of Repository Manager.`,
	}
)

func init() {
	RootCmd.AddCommand(RmCommand)
}
