package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func (c *Container) setupLogger(conf *LogConfig) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	lvl, err := log.ParseLevel(conf.Level)
	if err != nil {
		log.Warnf("log level \"%s\" unsupported", conf.Level)

		lvl = log.DebugLevel
	}

	log.SetLevel(lvl)
}
