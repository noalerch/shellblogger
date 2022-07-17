package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// TODO: user-specified editor (in config)
var content string = "blog/content"
var location string = "posts"
var editor string = "vim"

var deployNow bool

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a single blog post",
	Long: `post opens a text editor for writing a markdown blog post in markdown.
Use flag -D to deploy the post to remote server immediately after finishing.`,
	Run: func(cmd *cobra.Command, args []string) {
		post(args[0], content, location, editor)
		BuildSite("blog", "public")
	},
}

func post(postName string, content string, location string, editor string) {

	newPostLocation := location + "/" + postName

	newPostOutput := exec.Command("hugo", "new", "-c", content, newPostLocation+".md")

	// TODO: better error handling (keep to DRY)

	out, err := newPostOutput.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	editorPath, err := exec.LookPath(editor)
	if err != nil {
		fmt.Printf("Error %s while looking up path for %s", editorPath, editor)
	}

	edit := exec.Command(editorPath, content+"/"+newPostLocation+".md")
	// edit := exec.Command("vim") //editorPath)
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	err = edit.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	if deployNow {
		// TODO: less hardcoded, more configured
		BuildSite(sourceDir, outputDir)
		DeploySite(local, composeSSHDestination(user, server, destination))
		// FIXME: does not seem to deploy? maybe wrong dir?

	}

}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().BoolP("deploy", "d", false, "Immediately deploy post on exit")
}
