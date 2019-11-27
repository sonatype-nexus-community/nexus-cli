package cmd

import (
	"encoding/csv"
	"os"

	"github.com/spf13/cobra"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmRepositoriesCmd = &cobra.Command{
		Use:     "repositories",
		Aliases: []string{"repos", "r"},
		Short:   "Manage Nexus Repository Manager repositories",
		Long:    `List, create, and remove the repositories in your Nexus Repository Manager`,
		Run: func(cmd *cobra.Command, args []string) {
			rmListRepos()
		},
	}

	rmRepositoriesList = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List Nexus Repository Manager repositories",
		Long:    `List the repositories in your Nexus Repository Manager`,
		Run: func(cmd *cobra.Command, args []string) {
			rmListRepos()
		},
	}
)

func init() {
	RmCommand.AddCommand(rmRepositoriesCmd)
	rmRepositoriesCmd.AddCommand(rmRepositoriesList)
}

func rmListRepos() {
	w := csv.NewWriter(os.Stdout)

	w.Write([]string{"Name", "Format", "Type", "URL"})
	if repos, err := nexusrm.GetRepositories(rmClient); err == nil {
		for _, r := range repos {
			w.Write([]string{r.Name, r.Format, r.Type, r.URL})
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}
}
