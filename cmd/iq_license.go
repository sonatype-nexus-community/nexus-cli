package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	iqLicenseCmd = &cobra.Command{
		Use:   "license",
		Short: "Manage Nexus IQ Licenses",
		Long:  `Install, Uninstall, and inspect the licenses used for Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("license called")
		},
	}

	iqLicenseInstall = &cobra.Command{
		Use:   "install",
		Short: "install a Nexus IQ license",
		Long:  `install a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("install license called")
			fmt.Println(args)
		},
	}

	iqLicenseUninstall = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall a Nexus IQ license",
		Long:  `uninstall a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("uninstall license called")
		},
	}

	iqLicenseInfo = &cobra.Command{
		Use:   "info",
		Short: "show the details of the installed Nexus IQ license",
		Long:  `show the details of the installed Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("uninstall license called")
		},
	}
)

func init() {
	iqCmd.AddCommand(iqLicenseCmd)
	iqLicenseCmd.AddCommand(iqLicenseInstall)
	iqLicenseCmd.AddCommand(iqLicenseUninstall)
	iqLicenseCmd.AddCommand(iqLicenseInfo)
}
