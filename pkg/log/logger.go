package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Logger *logrus.Logger

func Init() error {
	Logger = logrus.New()
	fileName := "./log/" + time.Now().Format( "20060102" ) + ".log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create %v err : %v\n", fileName, err)
		return nil
	}
	writers := []io.Writer{
		file,
		os.Stdout}
	Logger.SetOutput(io.MultiWriter(writers...))
	Logger.SetFormatter(&logrus.JSONFormatter{})
	return nil
}

func Info(msg string, args ...interface{})  {
	Logger.Infof(msg, args...)
}

func Warn(msg string, args ...interface{})  {
	Logger.Warningf(msg, args...)
}

func Error(msg string, args ...interface{})  {
	Logger.Errorf(msg, args...)
}

func Panic (msg string, args ...interface{})  {
	Logger.Panicf(msg, args...)
}

