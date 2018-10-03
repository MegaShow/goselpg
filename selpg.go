package main

import (
	"fmt"
	"github.com/MegaShow/goselpg/command"
	"github.com/MegaShow/goselpg/print"
	"os"
)

func main() {
	settings, err := command.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: " + err.Error())
		os.Exit(2)
	}
	err = print.Execute(settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(2)
	}
}
