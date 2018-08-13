package lazy

import (
	"os/exec"
	"fmt"
	"log"
)

// Clone clones repo to the local file system
func (lc *LazyClient) Clone(cloneURL string) {
	out, err := exec.Command("git", "clone", cloneURL).Output()

	if err != nil {
		log.Fatal("Error while cloning")
	}

	fmt.Println("Yay, you may visit the directory", string(out))
}
