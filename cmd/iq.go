package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqCmd = &cobra.Command{
		Use:   "iq",
		Short: "command for managing functionality of Nexus IQ",
		Long:  `command for managing functionality of Nexus IQ`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("iq called")
		},
	}
)

func init() {
	rootCmd.AddCommand(iqCmd)
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
