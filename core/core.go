package core

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "core")

type Config struct {
}

type Core struct {
	config Config
}

func New(config Config) *Core {
	return &Core{
		config:config,
	}
}

func (c *Core) Run() {
	log.Warn("At least, we meet for the first time for the last time!")
}

func (c *Core) Close() {
	log.Info("Got stop signal")
}