package commands

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Link struct {
	Name string
}

type Application struct {
	Name   string
	Type   string
	Group  string
	Groups []string
	Tags   []string
	Links  []Link
}

type System struct {
	Applications []Application `applications`
}

var templateFilename string
var filterFilename string
var outputFilename string

var cmdConvert = &cobra.Command{
	Use:   "convert [file to convert]",
	Short: "Convert yaml to diagram",
	Long: `convert is for converting architecture description in yaml
    to different diagram formats.
    `,
	Run: convertRun,
}

func convertRun(cmd *cobra.Command, args []string) {
	filename := args[0]
	var system System
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &system)
	if err != nil {
		panic(err)
	}

	fmap := template.FuncMap{
		"replace": func(s1 string, s2 string) string {
			return strings.Replace(s2, s1, "", -1)
		},
	}

	tmpl := template.Must(template.New(templateFilename).Funcs(fmap).ParseFiles(templateFilename))

	err1 := tmpl.Execute(os.Stdout, system)
	if err1 != nil {
		panic(err1)
	}
}

func init() {
	RootCmd.AddCommand(cmdConvert)
	RootCmd.PersistentFlags().StringVar(&templateFilename, "template", "dot.tmpl", "template filename")
	RootCmd.PersistentFlags().StringVar(&filterFilename, "filters", "", "filters filename")
	RootCmd.PersistentFlags().StringVar(&outputFilename, "output", "", "output filename")
}
