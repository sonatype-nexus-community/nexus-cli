package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/cobra"

	privateiq "github.com/hokiegeek/gonexus-private/iq"
	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqComponentsCmd = &cobra.Command{
		Use:   "components",
		Short: "Work with the information in Nexus IQ about third-party components",
		Long:  `List and evaluate components with your Nexus IQ Server`,
		/*
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentsAll("")
			},
		*/
	}

	iqComponentsDetails = func() *cobra.Command {
		var format string

		c := &cobra.Command{
			Use:   "details",
			Short: "Get Component Details",
			Long:  `Get details of the given component from your Nexus IQ Server`,
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentDeets(format, args...)
			},
		}

		c.Flags().StringVarP(&format, "format", "f", "", "Golang template format to apply to output")

		return c
	}

	iqComponentsList = func() *cobra.Command {
		var format string

		c := &cobra.Command{
			Use:   "list",
			Short: "List all Nexus IQ components",
			Long:  `List all of the components in your Nexus IQ Server`,
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentsAll(format)
			},
		}

		c.Flags().StringVarP(&format, "format", "f", "", "Golang template format to apply to output")

		return c
	}

	iqComponentsEvaluate = func() *cobra.Command {
		var format, application string

		c := &cobra.Command{
			Use:   "evaluate",
			Short: "Evaluate a named component",
			Long: `Evaluate a given component for a given application with your Nexus IQ Server. Returns JSON output.
			
Each component to be evaluated can be passed in as a PURL identifier (https://github.com/package-url/purl-spec)
			`,
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentEval(format, application, args)
			},
			Args: cobra.MinimumNArgs(1),
			Example: `nexus iq components evaluate 'pkg:maven/org.bouncycastle/bcprov-jdk15on@1.55?type=jar'
nexus iq components evaluate 'pkg:pypi/django@1.11.1' 'pkg:gem/doorkeeper@4.3?platform=ruby'
nexus iq components evaluate --application AwesomeApp 'pkg:maven/axis/axis@1.2.1?type=jar'`,
		}

		c.Flags().StringVarP(&format, "format", "f", "", "Go template format to apply to output")
		c.Flags().StringVarP(&application, "application", "a", "", "The name or identifier of an application to evaluate the component against.")

		return c
	}

	iqComponentsRemediation = func() *cobra.Command {
		var format, application, organization, stage string

		c := &cobra.Command{
			Use:   "remediation",
			Short: "Determine remediation version of given components",
			Long: `Given component and a Nexus IQ application returns JSON output with the next version of the component that does not violate or fail Nexus IQ policies.
			
Each component to be evaluated can be passed in as a PURL identifier (https://github.com/package-url/purl-spec)
			`,
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentRemediation(format, application, organization, stage, args)
			},
			Args: cobra.MinimumNArgs(1),
			Example: `nexus iq components remediation 'pkg:maven/org.bouncycastle/bcprov-jdk15on@1.55?type=jar'
nexus iq components remediation 'pkg:pypi/django@1.11.1' 'pkg:gem/doorkeeper@4.3?platform=ruby'
nexus iq components remediation --application AwesomeApp 'pkg:maven/axis/axis@1.2.1?type=jar'`,
		}

		c.Flags().StringVarP(&format, "format", "f", "", "Go template format to apply to output")
		c.Flags().StringVarP(&application, "application", "a", "", "The name or identifier of an application whose policies will be used")
		c.Flags().StringVarP(&organization, "organization", "o", "", "The name or identifier of an organization whose policies will be used")
		c.Flags().StringVarP(&stage, "stage", "s", "build", "The stage to use when identifying non-failing versions")

		return c
	}
)

func init() {
	iqCmd.AddCommand(iqComponentsCmd)
	iqComponentsCmd.AddCommand(iqComponentsDetails())
	iqComponentsCmd.AddCommand(iqComponentsList())
	iqComponentsCmd.AddCommand(iqComponentsEvaluate())
	// iqComponentsCmd.AddCommand(iqComponentsRemediation())
}

func iqComponentDeets(format string, ids ...string) {
	iq := newIQClient()

	type catcher struct {
		id  string
		err error
	}

	errs := make([]catcher, 0)
	for _, id := range ids {
		c, err := nexusiq.NewComponentFromString(id)
		var components []nexusiq.ComponentDetail
		if err == nil {
			components, err = nexusiq.GetComponents(iq, []nexusiq.Component{*c})
		}
		if err != nil {
			errs = append(errs, catcher{id, err})
			continue
		}

		if format != "" {
			tmpl := template.Must(template.New("deets").Funcs(template.FuncMap{"json": templateJSONPretty}).Parse(format))
			tmpl.Execute(os.Stdout, components)
		} else {
			buf, err := json.MarshalIndent(components, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(buf))
		}
	}

	for _, e := range errs {
		log.Printf("error with %s: %v\n", e.id, e.err)
	}
}

func iqComponentsAll(format string) {
	iq := newIQClient()

	components, err := nexusiq.GetAllComponents(iq)
	if err != nil {
		log.Printf("error listing components: %v\n", err)
		return
	}

	if format != "" {
		tmpl := template.Must(template.New("deets").Funcs(template.FuncMap{"json": templateJSONPretty}).Parse(format))
		tmpl.Execute(os.Stdout, components)
	} else {
		buf, err := json.MarshalIndent(components, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf))
	}
}

func iqComponentEval(format, app string, components []string) {
	iq := newIQClient()

	comps := make([]nexusiq.Component, len(components))
	for i, c := range components {
		comps[i] = nexusiq.Component{PackageURL: c}
	}

	var report *nexusiq.Evaluation
	var err error
	if app != "" {
		report, err = nexusiq.EvaluateComponents(iq, comps, app)
	} else {
		report, err = privateiq.EvaluateComponentsWithRootOrg(iq, comps)
	}
	if err != nil {
		log.Fatal(err)
	}

	if format != "" {
		tmpl := template.Must(template.New("report").Funcs(template.FuncMap{"json": templateJSONPretty}).Parse(format))
		tmpl.Execute(os.Stdout, report)
	} else {
		json, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
	}
}

func getRemediationByPURL(iq nexusiq.IQ, application, organization, stage, component string) (nexusiq.Remediation, error) {
	c := &nexusiq.Component{PackageURL: component}

	switch {
	case application != "":
		return nexusiq.GetRemediationByApp(iq, *c, stage, application)
	case organization != "":
		return nexusiq.GetRemediationByOrg(iq, *c, stage, organization)
	}

	return nexusiq.Remediation{}, fmt.Errorf("could not get enough information")
}

func iqComponentRemediation(format, application, organization, stage string, components []string) {
	iq := newIQClient()

	type catcher struct {
		component string
		err       error
	}

	var remediations []nexusiq.Remediation

	errs := make([]catcher, 0)
	for _, c := range components {

		remediation, err := getRemediationByPURL(iq, application, organization, stage, c)
		if err != nil {
			errs = append(errs, catcher{c, err})
			continue
		}

		remediations = append(remediations, remediation)
	}

	if format != "" {
		tmpl := template.Must(template.New("remediation").Funcs(template.FuncMap{"json": templateJSONPretty}).Parse(format))
		tmpl.Execute(os.Stdout, remediations)
	} else {
		buf, err := json.MarshalIndent(remediations, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf))
	}

	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "error finding remediation for '%s': %v\n", e.component, e.err)
	}
}
