package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lacasian/gogo/examplehttp"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Example HTTP request",
	Long:  "Execute a HTTP request and print the result to the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		// uncomment the following lines for long-running process

		// stopChan := make(chan os.Signal, 1)
		// signal.Notify(stopChan, syscall.SIGINT)
		// signal.Notify(stopChan, syscall.SIGTERM)

		x := examplehttp.NewWorkerLib(examplehttp.Config{
			TargetURL: viper.GetString("target-url"),
		})

		err := x.Run()
		if err != nil {
			log.Fatal(err)
		}


		// uncomment the following lines for long-running process
		// select {
		// case <-stopChan:
		// 	log.Info("Got stop signal. Finishing work.")
		//  // close whatever there is to close
		// 	log.Info("Work done. Goodbye!")
		// }
	},
}

func init() {
	RootCmd.AddCommand(exampleCmd)

	exampleCmd.Flags().String("target-url", "", "The target of the example HTTP request")
}
