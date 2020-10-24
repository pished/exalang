package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/pished/exalang/repl"
)

// Source Code -> Tokens -> Abstract Syntax Tree (ABS)
// Source Code -> Tokens is called lexical analysis

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the EXAPUNKS language!\n", user.Username)
	fmt.Printf("Feel free to experiment with commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
