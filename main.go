package main

import (
	"fmt"
	"kong/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Kong's land\n", user.Username)
	fmt.Printf("Smash commands to see wonders!\n")
	repl.Start(os.Stdin, os.Stdout)
}
