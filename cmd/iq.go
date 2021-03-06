package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqClient nexusiq.IQ

	// IqCommand is the noun which handles any Nexus IQ actions
	IqCommand = func() *cobra.Command {
		c := &cobra.Command{
			Use:     "iq",
			Aliases: []string{"q"},
			Short:   "Subcommand for managing functionality of Nexus IQ",
			Long:    `Subcommand for managing functionality of Nexus IQ`,
			PersistentPreRun: func(cmd *cobra.Command, args []string) {
				iqClient = newIQClient()
			},
		}

		c.PersistentFlags().StringVarP(&iqUser, "user", "u", "", "your Nexus IQ user name.")
		c.PersistentFlags().StringVarP(&iqPassword, "password", "p", "", "your Nexus IQ password.")
		c.PersistentFlags().StringVarP(&iqServer, "server", "s", "http://localhost", "the address of the Nexus IQ Server to use.")
		c.PersistentFlags().IntVarP(&iqPort, "port", "", 8070, "the port which the Nexus IQ Server is listening on.")

		return c
	}()
)

func init() {
	RootCmd.AddCommand(IqCommand)
}

func newIQClient() nexusiq.IQ {
	iqServer := viper.GetString("iqServer")
	iqPort := viper.GetInt("iqPort")
	iqUser := viper.GetString("iqUser")
	iqPassword := viper.GetString("iqPassword")

	iqHost := fmt.Sprintf("%s:%d", iqServer, iqPort)
	iq, err := nexusiq.New(iqHost, iqUser, iqPassword)
	if err != nil {
		panic(fmt.Errorf("could not create client to Nexus IQ instance: %v", err))
	}
	return iq
}

func iqOrgsIDMap(iq nexusiq.IQ) (id2name map[string]string, name2id map[string]string, err error) {
	orgs, err := nexusiq.GetAllOrganizations(iq)
	if err != nil {
		return id2name, name2id, err
	}

	id2name = make(map[string]string)
	name2id = make(map[string]string)
	for _, o := range orgs {
		id2name[o.ID] = o.Name
		id2name[o.Name] = o.ID
	}

	return id2name, name2id, nil
}
