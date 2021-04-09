package cmd

import (
	"fmt"
	formatter "github.com/lacasian/logrus-module-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func initLogging() {
	logging := viper.GetString("logging")

	if verbose {
		logging = "*=debug"
	}

	if vverbose {
		logging = "*=trace"
	}

	if logging == "" {
		logging = "*=info"
	}

	viper.Set("logging", logging)

	f, err := formatter.New(formatter.NewModulesMap(logging))
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(f)

	log.Debug("Debug mode")
}

func buildDBConnectionString() {
	if viper.GetString("db.connection-string") == "" {
		user := viper.GetString("db.user")
		pass := viper.GetString("db.password")

		p := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s", viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.sslmode"), viper.GetString("db.dbname"), user, pass)
		viper.Set("db.connection-string", p)
	}
}

func mustGetSubconfig(v *viper.Viper, key string, out interface{}) {
	err := unmarshalSubconfig(v, key, out)
	if err != nil {
		log.Fatal(err)
	}
}

func unmarshalSubconfig(v *viper.Viper, key string, out interface{}) error {
	vc := subtree(v, key)
	if vc == nil {
		return errors.Errorf("key '%s' not found", key)
	}
	err := vc.Unmarshal(out)
	return err
}

func subtree(v *viper.Viper, name string) *viper.Viper {
	r := viper.New()
	for _, key := range v.AllKeys() {
		if strings.Index(key, name+".") == 0 {
			r.Set(key[len(name)+1:], v.Get(key))
		}
	}
	return r
}
