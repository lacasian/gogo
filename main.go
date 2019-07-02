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
	commands.RootCmd.Version = buildVersion

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
