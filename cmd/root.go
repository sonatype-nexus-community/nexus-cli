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

var templateJSONPretty = func(v interface{}) string {
	a, _ := json.MarshalIndent(v, "", "  ")
	return string(a)
}

var (
	cfgFile    string
	iqUser     string
	iqPassword string
	iqServer   string
	iqPort     int

	rootCmd = &cobra.Command{
		Use:     "nexus",
		Short:   `A CLI to interact with Sonatype Nexus IQ and Sonatype Nexus Repository Manager`,
		Long:    `A Command Line Interface to interact with Sonatype Nexus IQ and Sonatype Nexus Repository Manager`,
		Version: "0.2.0",
	}
)

// Execute builds the command tree for the CLI. Exits if an error if found.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfgFile", configPath, fmt.Sprintf("config file (default is %s)", configPath))
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
