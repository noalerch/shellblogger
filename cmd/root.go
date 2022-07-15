package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shbl",
	Short: "Shellblogger is a command-line blogging utility",
	Long: `Command-line tool for blog posts, generated statically via Hugo,
			and transferred to a user-specified server over SSH.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
