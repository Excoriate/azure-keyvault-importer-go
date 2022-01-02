package config

import (
	"github.com/withmandala/go-log"
	"os"
)

func GetLogger() *log.Logger {
	logger := log.New(os.Stderr).WithColor()
	return logger
}