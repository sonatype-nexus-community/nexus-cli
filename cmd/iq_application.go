package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqApplicationsCmd = &cobra.Command{
		Use:   "applications",
		Short: "Manage Nexus IQ applications",
		Long:  `List, create, and remove the applications in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("application called")
		},
	}

	iqApplicationsList = &cobra.Command{
		Use:   "list",
		Short: "List Nexus IQ applications",
		Long:  `List the applications in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqAppsList()
		},
	}
)

func init() {
	iqCmd.AddCommand(iqApplicationsCmd)
	iqApplicationsCmd.AddCommand(iqApplicationsList)
}

func iqAppsList() {
	iq := newIQClient()

	fmt.Printf("%s, %s, %s, %s\n", "Name", "Public ID", "ID", "Organization ID")
	orgsID2Name, _, _ := iqOrgsIDMap(iq)
	if apps, err := nexusiq.GetAllApplications(iq); err == nil {
		for _, a := range apps {
			fmt.Printf("%s, %s, %s, %s\n", a.Name, a.PublicID, a.ID, orgsID2Name[a.OrganizationID])
		}
	}
}
