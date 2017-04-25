package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func ReadIt(r *bufio.Reader) string {
	rec, _ := r.ReadString('\n')
	temp := strings.TrimSpace(rec)
	return temp

}

func main() {
	r := bufio.NewReader(os.Stdin)
	ctx := context.Background()
	fmt.Print("GitHub Username : ")
	username, _ := r.ReadString('\n')
	fmt.Print("Github Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	client := github.NewClient(tp.Client())
	fmt.Println("\nEnter the details for the repository you want to create\n")
	fmt.Print("Name of the repository : ")
	name := ReadIt(r)
	//fmt.Println(reflect.TypeOf(r))
	fmt.Print("Do you want it to be a private repository? [Y or N] : ")
	choice := ReadIt(r)
	var private bool
	if choice == "Y" || choice == "y" {
		private = true
	} else if choice == "N" || choice == "n" {
		private = false
	} else {
		return
	}
	fmt.Print("Give a one line description of your repository : ")
	description := ReadIt(r)

	repo := &github.Repository{
		Name:        github.String(name),
		Private:     &private,
		Description: github.String(description),
	}
	client.Repositories.Create(ctx, "", repo)
	var buffer bytes.Buffer
	buffer.WriteString("https://github.com/")
	buffer.WriteString(username)
	buffer.WriteString("/")
	buffer.WriteString(name)
	buffer.WriteString(".git")
	repoUrl := strings.Split(buffer.String(), "\n")
	repoUrl2 := repoUrl[0] + repoUrl[1]
	//fmt.Println(repoUrl2)
	out, err := exec.Command("git", "clone", repoUrl2).Output()
	if err != nil {
		fmt.Println("Repository created but ,error while cloning")
	} else {

		fmt.Println("You can cd to you repository now", string(out))
		fmt.Println("Go to the following link to access your repository online :- ", repoUrl2)
	}

}
