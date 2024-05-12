package logger

import (
	"fmt"
	"phone-numbers-server/internal/config"
)

// 0 - standart output
// 1 or higher - logging level
type Logger struct {
	Level uint
}

func (log Logger) Print(msg string, level uint) {
	if level <= log.Level {
		fmt.Println(msg)
	}
}

func Create(cfg config.LoggerConfig) (Logger, error) {
	return Logger{cfg.Level}, nil
}