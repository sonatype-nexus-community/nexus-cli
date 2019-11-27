package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
)

var (
	rmClient nexusrm.RM

	// RmCommand is the noun which handles any Nexus Repository actions
	RmCommand = func() *cobra.Command {
		c := &cobra.Command{
			Use:     "rm",
			Aliases: []string{"r"},
			Short:   "Subcommand for managing functionality of Nexus Repository Manager.",
			Long:    `Subcommand for managing functionality of Nexus Repository Manager.`,
			PersistentPreRun: func(cmd *cobra.Command, args []string) {
				rmClient = newRMClient()
			},
		}

		c.PersistentFlags().StringVarP(&rmUser, "user", "u", "", "your Nexus Repository Manager user name.")
		c.PersistentFlags().StringVarP(&rmPassword, "password", "p", "", "your Nexus Repository Manager password.")
		c.PersistentFlags().StringVarP(&rmServer, "server", "s", "http://localhost", "the address of the Nexus Repository Manager server to use.")
		c.PersistentFlags().IntVarP(&rmPort, "port", "", 8081, "the port which the Nexus Repository Manager server is listening on.")

		return c
	}()
)

func init() {
	RootCmd.AddCommand(RmCommand)
}

func newRMClient() nexusrm.RM {
	rmServer := viper.GetString("rmServer")
	rmPort := viper.GetInt("rmPort")
	rmUser := viper.GetString("rmUser")
	rmPassword := viper.GetString("rmPassword")

	rmHost := fmt.Sprintf("%s:%d", rmServer, rmPort)
	rm, err := nexusrm.New(rmHost, rmUser, rmPassword)
	if err != nil {
		panic(fmt.Errorf("could not create client to Nexus Repository Manager instance: %v", err))
	}
	return rm
}
