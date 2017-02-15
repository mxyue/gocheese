package config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func logDefine() {
	log.SetFormatter(&log.JSONFormatter{})
	logfile, err := os.OpenFile("cheese.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)

	log.SetLevel(log.DebugLevel)
}
