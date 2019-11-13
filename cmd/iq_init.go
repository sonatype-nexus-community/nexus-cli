package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var iqInitCmd = &cobra.Command{
	Use:     "init [-u|--user] username [-p|--password] password ",
	Short:   "initializes the nexus CLI.",
	Long:    `initializes the nexus CLI and creates a config file (default is $HOME/.nexus.yaml`,
	Example: `init --user Dave --password Password123!`,
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.WriteConfigAs(viper.ConfigFileUsed())
		if err != nil {
			fmt.Println(viper.ConfigFileUsed())
			fmt.Println(err)
		} else {
			fmt.Println()
		}
	},
}

func init() {
	iqInitCmd.PersistentFlags().StringVarP(&iqUser, "user", "u", "", "your Nexus IQ user name.")
	iqInitCmd.PersistentFlags().StringVarP(&iqPassword, "password", "p", "", "your Nexus IQ password.")
	iqInitCmd.PersistentFlags().StringVarP(&iqServer, "server", "s", "http://localhost", "the address of the Nexus IQ server to use.")
	iqInitCmd.PersistentFlags().IntVarP(&iqPort, "port", "", 8070, "the port which the Nexus IQ server is listening on.")

	viper.BindPFlag("iqUser", iqInitCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("iqPassword", iqInitCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("iqServer", iqInitCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("iqPort", iqInitCmd.PersistentFlags().Lookup("port"))

	iqInitCmd.MarkFlagRequired("user")
	iqInitCmd.MarkFlagRequired("password")

	iqCmd.AddCommand(iqInitCmd)
}
