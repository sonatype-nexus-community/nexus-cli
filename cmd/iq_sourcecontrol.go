package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqSourceControlCmd = &cobra.Command{
		Use:     "source-control",
		Aliases: []string{"sc"},
		Short:   "Manage source control configurations in your Nexus IQ instance",
	}

	iqSourceControlCreate = func() *cobra.Command {
		return &cobra.Command{
			Use:     "create",
			Aliases: []string{"c"},
			Short:   "create <appId> <repositoryUrl> <accessToken>",
			Run: func(cmd *cobra.Command, args []string) {
				app, repo, token := args[0], args[1], args[2]
				iqScCreate(app, repo, token)
			},
		}
	}()

	iqSourceControlDelete = func() *cobra.Command {
		var appID, entryID string

		c := &cobra.Command{
			Use:     "delete",
			Aliases: []string{"d"},
			Short:   "deletes a source control entry",
			Run: func(cmd *cobra.Command, args []string) {
				var scEntryID string
				if entryID != "" {
					scEntryID = entryID
				} else {
					scEntry, err := nexusiq.GetSourceControlEntry(iqClient, appID)
					if err != nil {
						panic(err)
					}
					scEntryID = scEntry.ID
				}

				nexusiq.DeleteSourceControlEntry(iqClient, appID, scEntryID)

				fmt.Println("Deleted")
			},
		}

		c.Flags().StringVarP(&appID, "app", "a", "", "")
		c.Flags().StringVarP(&entryID, "id", "i", "", "")

		return c
	}()

	iqSourceControlList = func() *cobra.Command {
		return &cobra.Command{
			Use:     "list",
			Aliases: []string{"l"},
			Short:   "list source control entries",
			Long:    "list source control entries in your Nexus IQ instance",
			Run: func(cmd *cobra.Command, args []string) {
				appID := args[0]
				if appID != "" {
					entry, _ := nexusiq.GetSourceControlEntry(iqClient, appID)
					fmt.Printf("%v\n", entry)
				} else {
					apps, err := nexusiq.GetAllApplications(iqClient)
					if err != nil {
						panic(err)
					}
					for _, app := range apps {
						if entry, err := nexusiq.GetSourceControlEntry(iqClient, app.PublicID); err == nil {
							fmt.Printf("%s: %v\n", app.PublicID, entry)
						}
					}
				}
			},
		}
	}()
)

func init() {
	IqCommand.AddCommand(iqSourceControlCmd)
	iqSourceControlCmd.AddCommand(iqSourceControlCreate)
	iqSourceControlCmd.AddCommand(iqSourceControlDelete)
	iqSourceControlCmd.AddCommand(iqSourceControlList)
}

func iqScCreate(app, repo, token string) {
	err := nexusiq.CreateSourceControlEntry(iqClient, app, repo, token)
	if err != nil {
		panic(err)
	}

	entry, err := nexusiq.GetSourceControlEntry(iqClient, app)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q\n", entry)
}
