package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	licenseCmd = &cobra.Command{
		Use:   "license",
		Short: "Manage Nexus IQ Licenses",
		Long:  `Install, Uninstall, and inspect the licenses used for Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("license called")
		},
	}

	licenseInstall = &cobra.Command{
		Use:   "install",
		Short: "install a Nexus IQ license",
		Long:  `install a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("install license called")
			fmt.Println(args)
		},
	}

	licenseUninstall = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall a Nexus IQ license",
		Long:  `uninstall a Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("uninstall license called")
		},
	}

	licenseInfo = &cobra.Command{
		Use:   "info",
		Short: "show the details of the installed Nexus IQ license",
		Long:  `show the details of the installed Nexus IQ license`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("uninstall license called")
		},
	}
)

func init() {
	iqCmd.AddCommand(licenseCmd)
	licenseCmd.AddCommand(licenseInstall)
	licenseCmd.AddCommand(licenseUninstall)
	licenseCmd.AddCommand(licenseInfo)
}
