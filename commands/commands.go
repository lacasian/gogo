package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config  string
	version bool

	GogoCmd = &cobra.Command{
		Use:   "gogo",
		Short: "not doing anything",
		Long:  "I'm a simple boilerplate. Use me and change me. I'm your whore.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if config != "" {
				// get the filepath
				abs, err := filepath.Abs(config)
				if err != nil {
					logrus.Error("Error reading filepath: ", err.Error())
				}

				// get the config name
				base := filepath.Base(abs)

				// get the path
				path := filepath.Dir(abs)

				//
				viper.SetConfigName(strings.Split(base, ".")[0])
				viper.AddConfigPath(path)

				// Find and read the config file; Handle errors reading the config file
				if err := viper.ReadInConfig(); err != nil {
					logrus.Fatal("Failed to read config file: ", err.Error())
					os.Exit(1)
				}
			}
		},

		Run: func(cmd *cobra.Command, args []string) {

			// fall back on default help if no args/flags are passed
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func init() {

	// set config defaults
	// viper.SetDefault("some-flag", false)

	// persistent flags
	GogoCmd.PersistentFlags().StringP("backend", "b", "file://", "Hoarder backend driver")

	//
	viper.BindPFlag("backend", GogoCmd.PersistentFlags().Lookup("backend"))

	// local flags;
	GogoCmd.Flags().StringVar(&config, "config", "", "/path/to/config.yml")
	GogoCmd.Flags().BoolVarP(&version, "version", "v", false, "Display the current version of this CLI")

	// commands
	GogoCmd.AddCommand(helloCmd)
}
