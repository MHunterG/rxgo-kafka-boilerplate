package logs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Setup() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}
