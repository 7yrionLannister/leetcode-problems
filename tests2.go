package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	// flag arguments must be provided before any positional arguments
	args := os.Args
	stringFlag := flag.String("name", "default", "name of the user with flag")
	var stringFlagvar string
	flag.StringVar(&stringFlagvar, "namevar", "default", "name of the user with flag var")

	flag.Parse()

	// slog is a simple logging library
	// here we specify the output to be stdout and the log level to be debug
	mySlogInfoAndDebugLevel := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: false, Level: slog.LevelDebug}))
	// here we specify the output to be stderr and the log level to be error
	mySlogErrorLevel := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelError}))
	// two loggers are created, one for info and debug level and the other for error level
	// this is because we want to log info and debug level logs to stdout and error level logs to stderr
	// although mySlogInfoAndDebugLevel can and will log error level logs, it is not recommended to do so
	// as it will log error level logs to stdout

	mySlogInfoAndDebugLevel.Info(fmt.Sprintf("%+v", args))
	mySlogInfoAndDebugLevel.Debug(fmt.Sprint(*stringFlag))
	mySlogErrorLevel.Error(stringFlagvar)
	// example run go run tests2.go -name=pepe -namevar=asdasdasd a b c > out.log 2> error.log
}
