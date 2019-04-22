package main

import (
	"fmt"
	"os"

	"github.com/kwix/gogo/commands"
)

func main() {
	if err := commands.GogoCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
