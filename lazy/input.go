package lazy

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"strings"
)

// Reader struct with bufio reader
type Reader struct {
	reader 	*bufio.Reader
}

// NewReader returns lazy.Reader instance
func NewReader() *Reader {
	return &Reader {
		reader: bufio.NewReader(os.Stdin),
	}
}

// GetInput get input from user according to the type
func (r *Reader) GetInput(prompt string, inputType string, def string) (output string) {
	if len(prompt) > 0 {
		fmt.Print(prompt + " : ")
		if len(def) > 0 {
			fmt.Print("[" + def + "] ")
		}
	}
	var err error
	var input []byte
	if inputType == "password" {
		input, err = terminal.ReadPassword(syscall.Stdin)
	} else {
		input, err = r.reader.ReadBytes('\n')
	}


	if err != nil {
		log.Fatal(err)
	}

	output = strings.TrimSpace(string(input))

	if len(output) == 0 {
		output = def
	}

	return
}
