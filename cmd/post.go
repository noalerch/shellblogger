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

	out, err = edit.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
