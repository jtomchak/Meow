package main

import (
	"fmt"
	"meow/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Meow, the programming language!\n", user.Username)
	fmt.Printf("Use the repl to test out meow code...\n")
	repl.Start(os.Stdin, os.Stdout)
}
