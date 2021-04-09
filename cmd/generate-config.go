package cmd

import (
	"io/ioutil"

	"github.com/lacasian/gogo/confgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ignore = []string{"verbose", "v", "vv", "version", "help", "config", "connection-string"}
)
var (
	generateConfigCmd = &cobra.Command{
		Use:   "generate-config",
		Short: "generate a sample config file",
		Long:  "generates a sample config file named config-generated.yml",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

			if !viper.GetBool("with-defaults") {
				RootCmd.PersistentPreRun(cmd, args)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("starting config")
			c := viper.AllSettings()

			ba, err := confgen.Viper(c, cmd, ignore)
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile("config-generated.yml", ba, 0644)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("done writing config")
		},
	}
)

func init() {
	RootCmd.AddCommand(generateConfigCmd)

	generateConfigCmd.Flags().Bool("with-defaults", true, "Generate the config using the default values. If set to false and a config.yml is loaded, it will take the params from the config")

	addDBFlags(generateConfigCmd)
}
