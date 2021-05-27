package cmd

import (
	"log"
	"projectcreator/project"
	"projectcreator/template"
	"projectcreator/util"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)

	settings, err := util.GetSettings()

	if err != nil {
		createCmd.Flags().StringP("output", "o", ".", "Output directory")
	} else {
		defaults := settings.Defaults

		git := defaults.Git
		github := defaults.Github
		dir := defaults.ProjectDir

		if git {
			createCmd.Flags().BoolP("no-git", "", git, "Don't initalize git repository (override default configuration)")
		} else {
			createCmd.Flags().BoolP("git", "g", git, "Initalize git repository (override default configuration)")
		}

		if github {
			createCmd.Flags().BoolP("no-github", "", github, "Don't initalize GitHub repository (override default configuration)")
		} else {
			createCmd.Flags().BoolP("github", "", github, "Initialize GitHub repository (override default configuration)")
		}

		if dir != "" {
			createCmd.Flags().StringP("output", "o", dir, "Output directory")
		} else {
			createCmd.Flags().StringP("output", "o", ".", "Output directory")
		}
	}

	createCmd.Flags().StringP("template", "t", "default", "Template to create project")
}

var createCmd = &cobra.Command{
	Use:   "create [template]",
	Short: "Create a project using the specified template",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		templateFlag, _ := cmd.Flags().GetString("template")
		outputFlag, _ := cmd.Flags().GetString("output")
		gitFlag, _ := cmd.Flags().GetBool("git")

		templateData, err := template.LoadTemplate(templateFlag)

		if err != nil {
			log.Fatal(err)
		}

		createConfig := project.CreateConfig{
			Git:              gitFlag,
			GitHub:           false,
			PublicRepository: false,
		}

		project.Create(templateData, outputFlag, name, createConfig)
	},
}
