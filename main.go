package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/unkletayo/monkey/repl"
)

func main() {
	user, err := user.Current()
	fmt.Println(user)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
