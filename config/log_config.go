package config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	logfile, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)

	log.SetLevel(log.WarnLevel)
}
