package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()
	logger.Info("This is info level log")
	logger.Warning("This is warning level log")
	logger.Error("This is error level log")
	logger.Debug("This is debug level log")

	logger1 := logrus.New()
	logger1.WithFields(logrus.Fields{
		"User":  "ram",
		"Event": "log in",
	}).Info("User logged in")
}
