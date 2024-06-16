package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/K-Saikrishnan/go-monkey/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Monkey programming language!\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
