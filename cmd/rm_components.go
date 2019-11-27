package cmd

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/spf13/cobra"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmComponentsCmd = &cobra.Command{
		Use:     "components",
		Aliases: []string{"artifacts", "c"},
		Short:   "Manage Nexus Repository Manager components and assets",
		Long:    `List, create, and remove the components in your Nexus Repository Manager`,
		Run: func(cmd *cobra.Command, args []string) {
			rmListRepoComponents(args)
		},
	}

	rmComponentsList = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List Nexus Repository Manager components",
		Long:    `List the components in your Nexus Repository Manager in a table`,
		Run: func(cmd *cobra.Command, args []string) {
			rmListRepoComponents(args)
		},
	}
)

func init() {
	RmCommand.AddCommand(rmComponentsCmd)
	rmComponentsCmd.AddCommand(rmComponentsList)
}

func rmListRepoComponents(repos []string) {
	w := csv.NewWriter(os.Stdout)

	w.Write([]string{"Repository", "Group", "Name", "Version", "Tags"})
	if len(repos) == 0 {
		all, _ := nexusrm.GetRepositories(rmClient)
		for _, r := range all {
			repos = append(repos, r.Name)
		}
	}

	for _, repo := range repos {
		if components, err := nexusrm.GetComponents(rmClient, repo); err == nil {
			for _, c := range components {
				w.Write([]string{c.Repository, c.Group, c.Name, c.Version, strings.Join(c.Tags, ";")})
			}
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}
}
