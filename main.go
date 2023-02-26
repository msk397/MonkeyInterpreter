package main

import (
	"MonkeyInterpreter/REPL"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user1, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey playground!", user1.Username)
	fmt.Printf("Feel free to type in commands\n")
	REPL.Start(os.Stdin, os.Stdout)
}
