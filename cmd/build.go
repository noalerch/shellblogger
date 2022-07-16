package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// TODO: read source and output directories from config
// 		 do this with viper?
var sourceDir string = "blog"
var outputDir string = "public"

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds site locally with Hugo",
	Long: `build generates the static website locally with Hugo,
accessed by index.html in your directory.
At the moment it is equivalent to just typing hugo.`,

	// arguments: paths to site
	// default (no args): default site as defined by config
	// TODO: flags to decide input and output dirs?
	// TODO: build multiple sites concurrently?? could be useful for deploying many sites from a central place

	// FIXME: nothing happens when using command
	Run: func(cmd *cobra.Command, args []string) {
		buildSite(sourceDir, outputDir)

	},
}

func buildSite(source string, output string) {
	out, err := exec.Command("hugo", "-s", sourceDir, "-d", outputDir).Output() // -s %s -d %s", source, output)

	// err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
