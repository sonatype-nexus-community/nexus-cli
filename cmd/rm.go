package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "command for managing functionality of Repository Manager.",
		Long:  `command for managing functionality of Repository Manager.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rm called")
		},
	}
)

func init() {
	rootCmd.AddCommand(rmCmd)
}
