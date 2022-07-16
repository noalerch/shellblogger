/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

	// editorOutput := exec.Command(editor, content+"/"+newPostLocation)
	edit := exec.Command(editorPath)
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	edit.Stderr = os.Stderr
	err = edit.Start()

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
