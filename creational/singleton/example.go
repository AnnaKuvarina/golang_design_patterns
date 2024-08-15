package main

import "fmt"

func main() {
	log := getLoggerInstance()

	log.SetLogLevel(1)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(2)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(3)
	log.Log("This is a log message")

	// Logger instance concurrently
	for i := 1; i < 10; i++ {
		go getLoggerInstance()
	}

	fmt.Scanln()
}
