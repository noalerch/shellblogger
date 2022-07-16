package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config [SETTING] [VALUE]",
	Short: "Configure local project",
	Long: `config lets you configure the project.
Use shblog config default for default settings.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
