package log

import (
	"log"
)

var debug bool

func SetDebug() {
	debug = true
}

func Info(msg string) {
	if debug {
		log.Print(msg)
	}
}

func Infof(fmt string, args ...interface{}) {
	if debug {
		log.Printf(fmt, args...)
	}
}

func Fatal(msg string) {
	log.Fatal(msg)
}

func Fatalf(fmt string, args ...interface{}) {
	log.Fatalf(fmt, args...)
}
