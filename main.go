package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {
	url, username, token := os.Getenv("URL"), os.Getenv("USERNAME"), os.Getenv("TOKEN")
	_, err := git.PlainClone("./share", false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: username, // yes, this can be anything except an empty string
			Password: token,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("clone done")
}
