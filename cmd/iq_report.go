package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	privateiq "github.com/sonatype/gonexus-private/iq"
)

var (
	iqReportCmd = &cobra.Command{
		Use:   "report",
		Short: "Manage Nexus IQ reports",
		Long:  `View application reports stored in your Nexus IQ Server`,
	}

	iqReportReevaluate = &cobra.Command{
		Use:   "reevaluate",
		Short: "Reevaluate Nexus IQ application reports",
		Long:  `Reevaluate application reportss in your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			iqReevaluate(args)
		},
	}
)

func init() {
	IqCommand.AddCommand(iqReportCmd)
	iqReportCmd.AddCommand(iqReportReevaluate)
}

func iqReevaluate(apps[] string) {
	if len(apps) == 0 {
		if err := privateiq.ReevaluateAllReports(iqClient); err != nil {
			panic(fmt.Sprintf("could not re-evaluate reports: %v", err))
		}
		return
	}

	for _, app := range apps {
		splitPos := strings.LastIndex(app, ":")
		appID := app[:splitPos]
		stage := app[splitPos+1:]

		if err := privateiq.ReevaluateReportByApp(iqClient, appID, stage); err != nil {
			fmt.Printf("warn: could not re-evaluate report for '%s' at '%s' stage: %v", appID, stage, err)
		}
	}
}
