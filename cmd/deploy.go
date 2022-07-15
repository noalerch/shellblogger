package cmd

import (
	"github.com/spf13/cobra"
)

var cmdDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys built site to remote host",
	Long: `deploy transfers the content of a directory to a remote host over SSH (via SCP).
In this program it is primarily directed at the build contents of a site`,

	// TODO: catch error if site directory does not exist

}
