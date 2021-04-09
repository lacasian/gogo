package cmd

import (
	"io/ioutil"

	"github.com/lacasian/gogo/confgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ignore = []string{"verbose", "v", "vv", "version", "help", "config", "connection-string", "password"}
)
var (
	generateConfigCmd = &cobra.Command{
		Use:   "generate-config",
		Short: "generate a sample config file",
		Long:  "generates a sample config file named config-generated.yml",
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

	addDBFlags(generateConfigCmd)
}
