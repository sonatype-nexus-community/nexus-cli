package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
	privateiq "github.com/sonatype/gonexus-private/iq"
)

var (
	iqPoliciesCmd = &cobra.Command{
		Use:     "policies",
		Short:   "(beta) Do stuff with policies",
		Aliases: []string{"pol", "p"},
	}

	iqPoliciesImport = &cobra.Command{
		Use:     "import",
		Short:   "Import the indicated policies",
		Aliases: []string{"i"},
		Run: func(cmd *cobra.Command, args []string) {
			importPolicies(args[0])
		},
	}

	iqPoliciesExport = &cobra.Command{
		Use:     "export",
		Aliases: []string{"a"},
		Short:   "exports the policies of the indicated IQ",
		Run: func(cmd *cobra.Command, args []string) {
			exportPolicies()
		},
	}

	iqPoliciesList = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls, l"},
		Short:   "Lists all policies configured on the instance",
		Run: func(cmd *cobra.Command, args []string) {
			listPolicies()
		},
	}
)

func init() {
	IqCommand.AddCommand(iqPoliciesCmd)
	iqPoliciesCmd.AddCommand(iqPoliciesImport)
	iqPoliciesCmd.AddCommand(iqPoliciesExport)
	iqPoliciesCmd.AddCommand(iqPoliciesList)
}

func importPolicies(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = privateiq.ImportPolicies(iqClient, file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Policies imported")
}

func exportPolicies() {
	policies, err := privateiq.ExportPolicies(iqClient)
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.MarshalIndent(policies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}

func listPolicies() {
	w := csv.NewWriter(os.Stdout)

	w.Write([]string{"Name", "PolicyType", "ThreatLevel", "OwnerID", "OwnerType", "ID"})
	if policies, err := nexusiq.GetPolicies(iqClient); err == nil {
		for _, p := range policies {
			w.Write([]string{p.Name, p.PolicyType, strconv.Itoa(p.ThreatLevel), p.OwnerID, p.OwnerType, p.ID})
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}
}
