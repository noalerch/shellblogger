package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

////////// temporary placeholders before i fix args and config
var local = "blog/public/*"
var user = "nole3668"
var server = "polhem.it.uu.se"
var destination = "public_html/blog"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys built site to remote host",
	Long: `deploy transfers the content of a directory to a remote host over SSH (via SCP).
In this program it is primarily directed at the build contents of a site`,

	// TODO: catch error if site directory does not exist
	Run: func(cmd *cobra.Command, args []string) {
		remote := composeSSHDestination(user, server, destination)
		deploy(local, remote)
	},
}

func deploy(content string, destination string) {
	out, err := exec.Command("scp", "-r", content, destination).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func composeSSHDestination(user string, server string, outDir string) string {
	var output string = fmt.Sprintf("%s@%s:%s", user, server, outDir)

	return output
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
