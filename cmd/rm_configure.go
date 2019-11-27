package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmConfigureCmd = &cobra.Command{
		Use:     "configure",
		Aliases: []string{"config"},
		Short:   "Manage Nexus Repository Manager system configuration",
		Long:    `Manage Nexus Repository Manager system configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			rmStatus()
		},
	}
)

func init() {
	RmCommand.AddCommand(rmConfigureCmd)
}

func rmStatus() {
	fmt.Printf("Server: %s:%d\n", rmServer, rmPort)
	fmt.Printf("Readable: %v\n", nexusrm.StatusReadable(rmClient))
	fmt.Printf("Writable: %v\n", nexusrm.StatusWritable(rmClient))

	state, _ := nexusrm.GetReadOnlyState(rmClient)
	fmt.Println(state)
}
