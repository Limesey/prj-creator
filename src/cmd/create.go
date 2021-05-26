package cmd

import (
	"log"
	"projectcreator/project"
	"projectcreator/template"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("template", "t", "default", "Template to create project")
	createCmd.Flags().StringP("output", "o", ".", "Output directory")
}

var createCmd = &cobra.Command{
	Use:   "create [template]",
	Short: "Create a project using the specified template",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		templateFlag, _ := cmd.Flags().GetString("template")
		outputFlag, _ := cmd.Flags().GetString("output")

		templateData, err := template.LoadTemplate(templateFlag)

		if err != nil {
			log.Fatal(err)
		}

		project.Create(templateData, outputFlag, name)
	},
}
