package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	formatter "github.com/kwix/logrus-module-formatter"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.WithField("module", "main")

var (
	config  string
	version bool
	verbose bool
	logging string

	RootCmd = &cobra.Command{
		Use:   "gogo",
		Short: "not doing anything",
		Long:  "I'm a simple boilerplate. Use me and change me. I'm your whore.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			configLoaded := false

			if config != "" {
				// get the filepath
				abs, err := filepath.Abs(config)
				if err != nil {
					log.Error("Error reading filepath: ", err.Error())
				}

				// get the config name
				base := filepath.Base(abs)

				// get the path
				path := filepath.Dir(abs)

				//
				viper.SetConfigName(strings.Split(base, ".")[0])
				viper.AddConfigPath(path)
			}

			viper.AddConfigPath(".")

			// Find and read the config file; Handle errors reading the config file
			if err := viper.ReadInConfig(); err != nil {
				log.Info("Could not load config file. Falling back to args. Error: ", err)
			} else {
				configLoaded = true
			}

			if viper.GetString("db.connection-string") == "" && configLoaded {
				var user, pass string
				if !viper.IsSet("db.user") {
					user = viper.GetString("PG_USER")
				} else {
					user = viper.GetString("db.user")
				}

				if !viper.IsSet("db.password") {
					pass = viper.GetString("PG_PASSWORD")
				} else {
					pass = viper.GetString("db.password")
				}

				p := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s", viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.sslmode"), viper.GetString("db.dbname"), user, pass)

				viper.Set("db.connection-string", p)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {

			// fall back on default help if no args/flags are passed
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func initLogging() {
	if verbose && logging == "" {
		logging = "*=debug"
	}

	if logging == "" {
		logging = "*=info"
	}

	f, err := formatter.New(formatter.NewModulesMap(logging))
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(f)

	log.Debug("Debug mode")
}

func init() {
	// bind Viper to env variables
	viper.BindEnv("PG_USER")
	viper.BindEnv("PG_PASSWORD")

	// persistent flags
	RootCmd.PersistentFlags().StringVar(&config, "config", "", "/path/to/config.yml")

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display debug messages")
	RootCmd.PersistentFlags().StringVar(&logging, "logging", "", "Display debug messages")

	// local flags;
	RootCmd.Flags().BoolVar(&version, "version", false, "Display the current version of this CLI")

	// commands
	RootCmd.AddCommand(runCmd)

	cobra.OnInitialize(initLogging)
}
