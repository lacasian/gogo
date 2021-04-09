package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Say hello!",
	Long:  "Address a wonderful greeting to the majestic executioner of this CLI",
	Run: func(cmd *cobra.Command, args []string) {
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT)
		signal.Notify(stopChan, syscall.SIGTERM)

		logrus.Warn("At least, we meet for the first time for the last time!")

		select {
		case <-stopChan:
			log.Info("Got stop signal. Finishing work.")
			// close whatever there is to close
			log.Info("Work done. Goodbye!")
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	addDBFlags(runCmd)
}
