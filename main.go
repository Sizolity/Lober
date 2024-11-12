package main

import (
	"fmt"

	"github.com/rot/greetings"
)

func main() {
	message, err := greetings.Hello("Deva")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("This is the entrance of rot's home..")
	fmt.Println(message)
}
