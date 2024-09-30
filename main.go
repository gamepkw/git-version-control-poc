package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	// Step 1: Initialize a new Git repository in a temporary directory
	dir, err := ioutil.TempDir("", "git-poc")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // Clean up the temp directory after the example runs

	fmt.Println("Initializing new Git repository at:", dir)

	repo, err := git.PlainInit(dir, false)
	if err != nil {
		panic(err)
	}

	// Step 2: Create a new file and add it to the repository
	filePath := fmt.Sprintf("%s/hello.txt", dir)
	err = ioutil.WriteFile(filePath, []byte("Hello, Git!"), 0644)
	if err != nil {
		panic(err)
	}

	// Get the working directory for the repository
	w, err := repo.Worktree()
	if err != nil {
		panic(err)
	}

	// Step 3: Add the new file to the staging area
	_, err = w.Add("hello.txt")
	if err != nil {
		panic(err)
	}

	// Step 4: Commit the file
	commit, err := w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("New commit created with hash:", commit)

	// Step 5: Display the commit history
	fmt.Println("\nCommit History:")

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	// Get the commit object for the HEAD
	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Printf("Commit: %s\nAuthor: %s\nDate: %s\nMessage: %s\n\n",
			c.Hash, c.Author.Name, c.Author.When, c.Message)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
