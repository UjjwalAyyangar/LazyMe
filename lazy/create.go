package lazy

import (
	"fmt"
	"log"
	"strings"
	"github.com/google/go-github/github"
	"context"
	"strconv"
)

func chooser(input string) bool {
	if input == "n" {
		return false
	} else if input == "y" {
		return true
	} else {
		log.Fatal("Invalid Input, try [y or n]")
	}
	return true
}

// CreateRepo will attempt on creating a repo for the profile
func (lc *LazyClient) CreateRepo() {

	fmt.Println("\nEnter the details for the repository you want to create\n")
	repo := lc.reader.GetInput("Name of the repository", "text", "")
	choice := strings.ToLower(lc.reader.GetInput("Do you want it to be private?", "text", "n"))
	
	var private = chooser(choice)

	description := lc.reader.GetInput("Give a one line description for your Repository", "text", "")

	readme := strings.ToLower(lc.reader.GetInput("Do you want to initialize this repository with a README?", "text", "n"))

	var autoinit = chooser(readme)

	repository := &github.Repository{
		Name:        github.String(repo),
		Private:     github.Bool(private),
		Description: github.String(description),
		AutoInit:	github.Bool(autoinit),
	}

	repository, response, err := lc.client.Repositories.Create(context.Background(), "", repository)

	if err != nil {
		log.Println(response.Status)
		log.Fatal(strconv.Itoa(response.Remaining) + " un-authorised requests left for this hour")
	}

	fmt.Println("Yohoo! Repo created, visit", repository.GetHTMLURL())

	clone := strings.ToLower(lc.reader.GetInput("\nDo you want to clone the repository?", "text", "y"))
	if chooser(clone) {
		lc.Clone(repository.GetCloneURL())
	}

	fmt.Println("Bye!")
}
