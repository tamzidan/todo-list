package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Setup will execute default setup initialization related to logger configuration.
// This will be applicable to the global logger
func Setup(level log.Level) {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
