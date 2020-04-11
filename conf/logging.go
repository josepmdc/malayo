package conf

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LoggingConfig struct {
	Level string
	File  string
}

func ConfigureLogging(config *LoggingConfig) (*logrus.Entry, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	if config.File != "" {
		f, errOpen := os.OpenFile(config.File, os.O_RDWR|os.O_APPEND, 0660)
		if errOpen != nil {
			return nil, errOpen
		}
		logrus.SetOutput(bufio.NewWriter(f))
	}

	level, err := logrus.ParseLevel(strings.ToUpper(config.Level))
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(level)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:true,
		DisableTimestamp:false,
	})

	return logrus.StandardLogger().WithField("hostname", hostname), nil
}
