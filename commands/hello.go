package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello!",
	Long:  "Address a wonderful greeting to the majestic executioner of this CLI",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Warn("At least, we meet for the first time for the last time!")
	},
}