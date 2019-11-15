package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	privateiq "github.com/hokiegeek/gonexus-private/iq"
)

var (
	iqLicenseCmd = &cobra.Command{
		Use:   "license",
		Short: "Manage Nexus IQ Licenses",
		Long:  `Install, Uninstall, and inspect the licenses used for Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("license called")
		},
	}

	iqLicenseInstall = &cobra.Command{
		Use:   "install",
		Short: "install a Nexus IQ license",
		Long:  `install a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			// log.Println("install license called")
			iqInstallLicense(args[0])
		},
	}

	iqLicenseUninstall = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall a Nexus IQ license",
		Long:  `uninstall a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("uninstall license called")
		},
	}

	iqLicenseInfo = &cobra.Command{
		Use:   "info",
		Short: "show the details of the installed Nexus IQ license",
		Long:  `show the details of the installed Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("uninstall license called")
		},
	}
)

func init() {
	iqCmd.AddCommand(iqLicenseCmd)
	iqLicenseCmd.AddCommand(iqLicenseInstall)
	// iqLicenseCmd.AddCommand(iqLicenseUninstall)
	// iqLicenseCmd.AddCommand(iqLicenseInfo)
}

func iqInstallLicense(licensePath string) {
	license, err := os.Open(licensePath)
	if err != nil {
		panic(err)
	}

	iq := newIQClient()
	if err = privateiq.InstallLicense(iq, license); err != nil {
		panic(err)
	}

	fmt.Println("Installed license")
}
