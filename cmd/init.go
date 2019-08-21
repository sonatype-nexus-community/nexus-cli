package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
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
	initCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "your Nexus IQ user name.")
	initCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "your Nexus IQ password.")
	initCmd.PersistentFlags().StringVarP(&server, "server", "s", "localhost", "the address of the Nexus IQ server to use.")
	initCmd.PersistentFlags().IntVarP(&port, "port", "", 8070, "the port which the Nexus IQ server is listening on.")

	viper.BindPFlag("user", initCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", initCmd.PersistentFlags().Lookup("password"))
	//TODO: We need to use a token not store creds locally - viper.BindPFlag("token")
	viper.BindPFlag("host", initCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("port", initCmd.PersistentFlags().Lookup("port"))

	initCmd.MarkFlagRequired("user")
	initCmd.MarkFlagRequired("password")

	iqCmd.AddCommand(initCmd)
}
