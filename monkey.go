package main

import (
	"fmt"
	"os"
	"os/user"

	"mkint/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wh311c0m3 %s\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
