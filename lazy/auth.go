package lazy

import (
	"github.com/google/go-github/github"
	"strings"
)

// AuthUser stores username and password and returns a github Auth Client
func AuthUser(reader *Reader) *github.BasicAuthTransport {

	// lets take the input
	var username string = reader.GetInput("Enter GitHub Username or Email", "text", "")
	var password string = reader.GetInput("Enter GitHub Password", "password", "")
	return &github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
}
