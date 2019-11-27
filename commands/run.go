package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Say hello!",
	Long:  "Address a wonderful greeting to the majestic executioner of this CLI",
	PreRun: func(cmd *cobra.Command, args []string) {
		bindViperToDBFlags(cmd)
	},
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
	addDBFlags(runCmd)
}