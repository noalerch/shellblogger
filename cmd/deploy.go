package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

////////// temporary placeholders before i fix args and config
// deploys whole public dir right now
var local = "blog/public/"
var user = "nole3668"
var server = "polhem.it.uu.se"
var destination = "public_html/blog/"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys built site to remote host",
	Long: `deploy transfers the content of a directory to a remote host over SSH (via SCP).
In this program it is primarily directed at the build contents of a site`,

	// TODO: catch error if site directory does not exist
	Run: func(cmd *cobra.Command, args []string) {
		remote := composeSSHDestination(user, server, destination)
		println(local)
		println(remote)
		deploy(local, remote)
	},
}

func deploy(local string, remote string) {
	cmd := exec.Command("scp", "-r", local, remote)
	// out, err := exec.Command("scp", "-r", "blog/public/*", "nole3668@polhem.it.uu.se:public_html/blog/").Output()

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + " " + stderr.String())
		// log.Fatal(err)
	}

	fmt.Println(out.String())
}

func composeSSHDestination(user string, server string, outDir string) string {
	var output string = fmt.Sprintf("%s@%s:%s", user, server, outDir)

	return output
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
