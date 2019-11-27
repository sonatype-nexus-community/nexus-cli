package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	u "os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TemplateJSONPretty is intended to be used as a text/template function that returns a json-encoded version of the given interface
var TemplateJSONPretty = func(v interface{}) string {
	a, _ := json.MarshalIndent(v, "", "  ")
	return string(a)
}

var (
	cfgFile string

	// RootCmd is the primary command that will handle the nouns
	RootCmd = &cobra.Command{
		Use:   "nexus",
		Short: `A CLI to interact with Sonatype Nexus IQ and Sonatype Nexus Repository Manager`,
		Long:  `A Command Line Interface to interact with Sonatype Nexus IQ and Sonatype Nexus Repository Manager`,
	}
)

// Execute builds the command tree for the CLI. Exits if an error if found.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	usr, err := u.Current()
	if err != nil {
		log.Fatal(err)
	}

	configPath := fmt.Sprintf("%s/.nexus.json", usr.HomeDir)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "cfgFile", configPath, "configuration file for the Sonatype Nexus platform")
}

func initConfig() {
	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		usr, err := u.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(usr.HomeDir)
		viper.SetConfigName(".nexus.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//Used for Debug
		//fmt.Printf(viper.GetString("user"))
	}
}
