package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/cobra"

	nexusiq "github.com/sonatype-nexus-community/gonexus/iq"
)

var (
	iqComponentsCmd = &cobra.Command{
		Use:   "components",
		Short: "Manage Nexus IQ components",
		Long:  `List and evaluate components with your Nexus IQ Server`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("components called")
		},
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

	/*
		iqComponentsEvaluate = &cobra.Command{
			Use:   "evaluate",
			Short: "Evaluate a named component",
			Long:  `Evaluate a given component for a given application with your Nexus IQ Server`,
			Run: func(cmd *cobra.Command, args []string) {
				iqComponentEval()
			},
		}
	*/
)

func init() {
	iqCmd.AddCommand(iqComponentsCmd)
	iqComponentsCmd.AddCommand(iqComponentsDetails())
	// iqComponentsCmd.AddCommand(iqComponentsEvaluate)
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

/*
func iqComponentEval() {
	// iq := newIQClient()
}
*/
