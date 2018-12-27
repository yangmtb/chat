package logging

import (
	"chat/pkg/file"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type level int

var (
	f *os.File

	defaultPrefix      = ""
	defaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	debug level = iota
	info
	warn
	err
	fatal
)

// Setup log setup
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err = file.MustOpen(fileName, filePath)
	if nil != err {
		log.Fatalf("logging.setup err:%v", err)
	}
	logger = log.New(f, defaultPrefix, log.LstdFlags)
}

// Debug out
func Debug(v ...interface{}) {
	setPrefix(debug)
	logger.Println(v...)
}

// Info out
func Info(v ...interface{}) {
	setPrefix(info)
	logger.Println(v...)
}

// Warn out
func Warn(v ...interface{}) {
	setPrefix(warn)
	logger.Println(v...)
}

// Error out
func Error(v ...interface{}) {
	setPrefix(err)
	logger.Println(v...)
}

// Fatal out
func Fatal(v ...interface{}) {
	setPrefix(fatal)
	logger.Fatalln(v...)
}

func setPrefix(l level) {
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[l], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[l])
	}
	logger.SetPrefix(logPrefix)
}
