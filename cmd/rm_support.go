package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmSupportCmd = &cobra.Command{
		Use:   "support",
		Short: "Nexus Repository Manager support features",
		Long:  `Nexus Repository Manager support features`,
	}

	rmSupportZipCmd = &cobra.Command{
		Use:   "zip",
		Short: "Create a support zip to send to the support team",
		Long:  "Create a support zip to send to the support team",
		Run: func(cmd *cobra.Command, args []string) {
			rmZip()
		},
	}
)

func init() {
	RmCommand.AddCommand(rmSupportCmd)
	rmSupportCmd.AddCommand(rmSupportZipCmd)
}

func rmZip() {
	zip, name, err := nexusrm.GetSupportZip(rmClient, nexusrm.NewSupportZipOptions())
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(name, zip, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Created %s\n", name)
}
