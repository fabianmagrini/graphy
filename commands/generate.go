package commands

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type SystemApplication struct {
	App Application `application`
}

var cmdGenerate = &cobra.Command{
	Use:   "generate [file to generate]",
	Short: "Generate diagram from yaml manifest",
	Long: `generate is for generating architecture diagram from description in yaml
    to different diagram formats.
    `,
	Run: generateRun,
}

func generateRun(cmd *cobra.Command, args []string) {
	var system System
	system.Applications = []Application{}

	var filters []Filter
	if filterFilename != "" {
		filters = getFilters(filterFilename)
	}

	for _, filename := range resolveArgs(args) {
		var r io.Reader
		var err error
		r, err = os.Open(filename)
		if err != nil {
			panic(err)
		}

		dec := yaml.NewDecoder(r)
		for {
			var application SystemApplication
			err = dec.Decode(&application)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			if filters != nil {
				if matchFilter(application.App, filters) {
					system.Applications = append(system.Applications, application.App)
				}
			} else {
				system.Applications = append(system.Applications, application.App)
			}
		}
	}

	fmap := template.FuncMap{
		"replace": func(s1 string, s2 string) string {
			return strings.Replace(s2, s1, "", -1)
		},
	}

	templateBase := path.Base(templateFilename)
	tmpl := template.Must(template.New(templateBase).Funcs(fmap).ParseFiles(templateFilename))

	if outputFilename != "" {
		outputFile, err1 := os.OpenFile(
			outputFilename,
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err1 != nil {
			panic(err1)
		}
		defer outputFile.Close()

		err2 := tmpl.Execute(outputFile, system)
		if err2 != nil {
			panic(err2)
		}
	} else {
		err1 := tmpl.Execute(os.Stdout, system)
		if err1 != nil {
			panic(err1)
		}
	}

}

func init() {
	RootCmd.AddCommand(cmdGenerate)
}

type Filter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func getFilters(filename string) []Filter {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c []Filter
	json.Unmarshal(raw, &c)
	return c
}

func matchFilter(application Application, filters []Filter) bool {
	for _, filter := range filters {
		if strings.EqualFold("tags", filter.Name) {
			for _, value := range filter.Values {
				if application.Tags != nil {
					for _, tag := range application.Tags {
						if strings.EqualFold(value, tag) {
							return true
						}
					}
				}
			}
		} else if strings.EqualFold("groups", filter.Name) {
			for _, value := range filter.Values {
				if application.Groups != nil {
					for _, group := range application.Groups {
						if strings.EqualFold(value, group) {
							return true
						}
					}
				}
			}
		} else if strings.EqualFold("group", filter.Name) {
			for _, value := range filter.Values {
				if application.Group != "" {
					if strings.EqualFold(value, application.Group) {
						return true
					}
				}
			}
		}
	}
	return false
}

func resolveArgs(args []string) []string {
	var result = []string{}
	for _, arg := range args {
		matches, err := filepath.Glob(arg)
		if err != nil {
			panic(err)
		}
		result = append(result, matches...)
	}
	return result
}
