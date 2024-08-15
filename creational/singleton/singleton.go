package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	logLevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

// the logger instance
var logger *MyLogger

// to enforce the goroutine safety
var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating the logger instance")
		logger = &MyLogger{}
	})

	fmt.Println("Returning logger instance")
	return logger
}
