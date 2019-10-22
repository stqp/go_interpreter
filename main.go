package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/stqp/go_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming language!\n",
		user.Username)
	fmt.Println("Feel free to type in command!\n")
	repl.Start(os.Stdin, os.Stdout)

}
