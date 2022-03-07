package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	/******** Levels ********/
	// Info messages
	log.Info("This is some info")
	log.Info("This is some more info")

	// Warning Messages
	log.Warn("This is a warning message")

	// Error Messages
	log.Error("This is an Error")

	// Panic Messages
	// log.Panic("This is a Panic Message")

	/******** Fields ********/
	// Log messages with Fields
	log.WithFields(
		log.Fields{
			"Day":  "Foo",
			"Time": "bar",
		},
	).Info("This is a JSON message")

	/******** Formatters ********/
	// JSON Formatter With custom Timestamp format
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Info("This message is logged in JSON format")
	log.WithFields(log.Fields{
		"Type":  "JSON",
		"Value": 1,
	}).Info("This message is logged in JSON format with fields")

	// TEXT Formatter With custom Timestamp format
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	log.Info("This message is logged in TEXT format")
	log.WithFields(log.Fields{
		"Type":  "Text",
		"Value": 1,
	}).Info("This message is logged in TEXT format with fields")

}
