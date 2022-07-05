package main

import (
	"errors"
	stdlog "log"
	"os"
)

type PosLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

var log PosLogger = stdlog.New(os.Stderr, "", stdlog.LstdFlags)

func SetLogger(logger PosLogger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}
	log = logger
	return nil
}
