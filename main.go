package main

import (
	"fmt"
	"os"

	"github.com/kwix/gogo/commands"
)

var (
	buildVersion string
)

func main() {
	commands.GogoCmd.Version = buildVersion

	if err := commands.GogoCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
