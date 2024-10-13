package main

import (
	"fmt"
	"groupie-tracker/server"
	"os"
)

var (
	Port   = "8080"
	Adress = "0.0.0.0"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	err := server.Run(Adress, Port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
