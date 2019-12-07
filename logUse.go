package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

//DefaultFieldHook is test
type DefaultFieldHook struct {
}

// Fire s
func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "MyAppName"
	fmt.Println("this is fire")
	return nil
}

//
func (hook *DefaultFieldHook) Levels() []log.Level {
	fmt.Println("this is levels")
	return log.AllLevels
}

var logger = log.New()

func init() {
	//logger.SetFormatter(&log.JSONFormatter{})
	logger.Formatter = &log.JSONFormatter{}
	logger.SetOutput(os.Stdout)
	//logger.SetLevel(log.InfoLevel)
	logger.Level = log.InfoLevel

	sss := &DefaultFieldHook{}
	logger.AddHook(sss)
}

func main() {
	logger.WithFields(log.Fields{
		"a": "fdsafa",
	}).Info("FDSAfsaf")
	logger.WithFields(log.Fields{
		"a": "fdsafa",
	}).Warn("FDSAfsaf")
}
