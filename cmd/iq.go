package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	iqCmd = &cobra.Command{
		Use:   "iq",
		Short: "command for managing functionality of Nexus IQ",
		Long:  `command for managing functionality of Nexus IQ`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("iq called")
		},
	}
)

func init() {
	rootCmd.AddCommand(iqCmd)
}
