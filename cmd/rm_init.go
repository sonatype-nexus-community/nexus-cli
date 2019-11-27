package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rmUser     string
	rmPassword string
	rmServer   string
	rmPort     int

	rmInitCmd = &cobra.Command{
		Use:     "init [-u|--user] username [-p|--password] password ",
		Short:   "initializes the nexus CLI",
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
)

func init() {
	rmInitCmd.PersistentFlags().StringVarP(&rmUser, "user", "u", "", "your Nexus Repository Manager user name.")
	rmInitCmd.PersistentFlags().StringVarP(&rmPassword, "password", "p", "", "your Nexus Repository Manager password.")
	rmInitCmd.PersistentFlags().StringVarP(&rmServer, "server", "s", "http://localhost", "the address of the Nexus Repository Manager server to use.")
	rmInitCmd.PersistentFlags().IntVarP(&rmPort, "port", "", 8081, "the port which the Nexus Repository Manager server is listening on.")

	viper.BindPFlag("rmUser", rmInitCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("rmPassword", rmInitCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("rmServer", rmInitCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("rmPort", rmInitCmd.PersistentFlags().Lookup("port"))

	rmInitCmd.MarkFlagRequired("user")
	rmInitCmd.MarkFlagRequired("password")

	RmCommand.AddCommand(rmInitCmd)
}
