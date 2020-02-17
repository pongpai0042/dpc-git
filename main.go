package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {
	url, username, token := os.Getenv("URL"), os.Getenv("USERNAME"), os.Getenv("TOKEN")
	opts := &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	if token != "" {
		opts.Auth = &http.BasicAuth{
			Username: username, // yes, this can be anything except an empty string
			Password: token,
		}
	}
	_, err := git.PlainClone("./share", false, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("clone done")
}
