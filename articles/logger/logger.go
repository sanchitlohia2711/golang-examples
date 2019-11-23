package main

import (
	"fmt"
	"log"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

//Fields logrus fields
type Fields logrus.Fields

func init() {
	refresh()
	go timer(time.Second * 1)
}

//initialize logger
func refresh() {
	logger = logrus.New()
	path := "/var/log/service/old.UTC."
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName("/var/log/service/current"),
		rotatelogs.WithMaxAge(time.Second*10),
		rotatelogs.WithRotationTime(time.Second*1),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	writer.Write([]byte(""))

	logger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  writer.CurrentFileName(),
		logrus.WarnLevel:  writer.CurrentFileName(),
		logrus.ErrorLevel: writer.CurrentFileName(),
		logrus.FatalLevel: writer.CurrentFileName(),
	}, &logrus.JSONFormatter{}))

	return
}

func timer(t time.Duration) {
	for _ = range time.Tick(t) {
		refresh()
	}
}

//GetLogger : get the logger
func GetLogger() *logrus.Logger {
	if logger == nil {
		refresh()
	}
	return logger
}
