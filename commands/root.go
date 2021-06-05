package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "graphy",
	Short: "Create diagrams from yaml",
	Long: `CLI tool to create diagrams 
            from yaml descriptions.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
