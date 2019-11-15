package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqOrganizationsCmd = &cobra.Command{
		Use:   "organizations",
		Short: "Manage Nexus IQ organizations",
		Long:  `List, create, and remove the organizations in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("organizations called")
		},
	}

	iqOrganizationsList = &cobra.Command{
		Use:   "list",
		Short: "List Nexus IQ organizations",
		Long:  `List the organizations in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqOrgsList()
		},
	}
)

func init() {
	iqCmd.AddCommand(iqOrganizationsCmd)
	iqOrganizationsCmd.AddCommand(iqOrganizationsList)
}

func iqOrgsList() {
	iq := newIQClient()

	fmt.Printf("%s, %s\n", "Name", "ID")
	if orgs, err := nexusiq.GetAllOrganizations(iq); err == nil {
		for _, o := range orgs {
			fmt.Printf("%s, %s\n", o.Name, o.ID)
		}
	}
}
