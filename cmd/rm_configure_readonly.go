package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmReadonlyCmd = &cobra.Command{
		Use:     "read-only",
		Aliases: []string{"ro"},
		Short:   "Manage Nexus Repository Manager read-only mode",
		Long:    `Manage Nexus Repository Manager read-only mode`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			rmStatus()
		},
	}

	rmReadonlyEnable = &cobra.Command{
		Use:     "enable",
		Aliases: []string{"e"},
		Short:   "enables read-only mode",
		Run: func(cmd *cobra.Command, args []string) {
			rmReadOnlyToggle()
		},
	}

	rmReadonlyRelease = func() *cobra.Command {
		var force bool

		c := &cobra.Command{
			Use:     "release",
			Aliases: []string{"r"},
			Short:   "releases from read-only mode",
			Run: func(cmd *cobra.Command, args []string) {
				rmReadOnly(false, force)
			},
		}

		c.Flags().BoolVarP(&force, "force", "f", false, "")

		return c
	}()
)

func init() {
	rmConfigureCmd.AddCommand(rmReadonlyCmd)
	rmReadonlyCmd.AddCommand(rmReadonlyEnable)
	rmReadonlyCmd.AddCommand(rmReadonlyRelease)
}

func rmReadOnly(enable, forceRelease bool) {
	if enable {
		nexusrm.ReadOnlyEnable(rmClient)
	} else {
		nexusrm.ReadOnlyRelease(rmClient, forceRelease)
	}
}

func rmReadOnlyToggle() {
	state, err := nexusrm.GetReadOnlyState(rmClient)
	if err != nil {
		return
	}
	if state.Frozen {
		rmReadOnly(false, false)
	} else {
		rmReadOnly(true, false)
	}
}
