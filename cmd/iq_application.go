package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqApplicationsCmd = &cobra.Command{
		Use:   "applications",
		Short: "Manage Nexus IQ applications",
		Long:  `List, create, and remove the applications in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqAppsList()
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

	iqApplicationsCreate = func() *cobra.Command {
		var name, id, organization string

		c := &cobra.Command{
			Use:   "create",
			Short: "Create a new Nexus IQ application",
			Long:  `Create an application in your Nexus IQ Server`,
			Run: func(cmd *cobra.Command, args []string) {
				iqAppsCreate(name, id, organization)
			},
		}

		c.Flags().StringVarP(&name, "name", "n", "", "The name of the new application")
		c.Flags().StringVarP(&id, "id", "i", "", "The identifier of the new application")
		c.Flags().StringVarP(&organization, "organization", "o", "", "The organization where the new application will be added to")

		c.MarkFlagRequired("name")
		c.MarkFlagRequired("id")
		c.MarkFlagRequired("organization")

		return c
	}

	iqApplicationsDelete = func() *cobra.Command {
		c := &cobra.Command{
			Use:   "delete",
			Short: "Delete a Nexus IQ application",
			Long:  `Delete an application from your Nexus IQ Server`,
			Run: func(cmd *cobra.Command, args []string) {
				iqAppsDelete(args)
			},
		}

		return c
	}
)

func init() {
	iqCmd.AddCommand(iqApplicationsCmd)
	iqApplicationsCmd.AddCommand(iqApplicationsList)
	iqApplicationsCmd.AddCommand(iqApplicationsCreate())
	iqApplicationsCmd.AddCommand(iqApplicationsDelete())
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

func iqAppsCreate(name, id, organization string) {
	iq := newIQClient()

	org, err := nexusiq.GetOrganizationByName(iq, organization)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding organization '%s': %v\n", organization, err)
		return
	}

	if _, err = nexusiq.CreateApplication(iq, name, id, org.ID); err != nil {
		fmt.Fprintf(os.Stderr, "error creating application '%s': %v\n", name, err)
		return
	}
	fmt.Printf("Created application %s (%s) in %s\n", name, id, organization)
}

func iqAppsDelete(ids []string) {
	iq := newIQClient()

	type catcher struct {
		id  string
		err error
	}

	errs := make([]catcher, 0)
	for _, id := range ids {
		if err := nexusiq.DeleteApplication(iq, id); err != nil {
			errs = append(errs, catcher{id, err})
			continue
		}
		fmt.Printf("Deleted application with ID %s\n", id)
	}

	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "error deleting application with ID '%s': %v\n", e.id, e.err)
	}
}
