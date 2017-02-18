package config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

const (
	logfile = "product.log"
)

func logDefine() {
	log.SetFormatter(&log.JSONFormatter{})
	err := os.Remove(logfile)
	logfile, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)

	log.SetLevel(log.DebugLevel)
}
