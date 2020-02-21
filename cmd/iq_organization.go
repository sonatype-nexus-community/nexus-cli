package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
	privateiq "github.com/sonatype/gonexus-private/iq"
)

var (
	iqOrganizationsCmd = &cobra.Command{
		Use:   "organizations",
		Short: "Manage Nexus IQ organizations",
		Long:  `List, create, and remove the organizations in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqOrgsList()
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

	iqOrganizationsCreate = &cobra.Command{
		Use:   "create",
		Short: "Create a new Nexus IQ organization",
		Long:  `Create a new organization in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqOrgsCreate(args)
		},
	}

	iqOrganizationsDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete a Nexus IQ organization",
		Long:  `Delete an organization from your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqOrgsDelete(args)
		},
	}
)

func init() {
	IqCommand.AddCommand(iqOrganizationsCmd)
	iqOrganizationsCmd.AddCommand(iqOrganizationsList)
	iqOrganizationsCmd.AddCommand(iqOrganizationsCreate)
	iqOrganizationsCmd.AddCommand(iqOrganizationsDelete)
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

func iqOrgsCreate(names []string) {
	type catcher struct {
		name string
		err  error
	}

	errs := make([]catcher, 0)
	for _, name := range names {
		orgID, err := nexusiq.CreateOrganization(iqClient, name)
		if err != nil {
			errs = append(errs, catcher{name, err})
			continue
		}
		fmt.Printf("Created organization: %s (%s)\n", name, orgID)
	}

	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "error creating organization '%s': %v\n", e.name, e.err)
	}
}

func iqOrgsDelete(names []string) {
	type catcher struct {
		name string
		err  error
	}

	errs := make([]catcher, 0)
	for _, name := range names {
		org, err := nexusiq.GetOrganizationByName(iqClient, name)
		if err == nil {
			err = privateiq.DeleteOrganization(iqClient, org.ID)
		}
		if err != nil {
			errs = append(errs, catcher{name, err})
			continue
		}
		fmt.Printf("Deleted organization: %s (%s)\n", name, org.ID)
	}

	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "error deleting organization '%s': %v\n", e.name, e.err)
	}
}
