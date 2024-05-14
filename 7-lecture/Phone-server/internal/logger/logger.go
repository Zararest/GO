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

// FIXME: make logger async

// 0 - always true
// 1 - main steps of the server
// 2 - internal organization
// 3 - values dump
func (log Logger) Print(msg string, level uint) {
	if level <= log.Level {
		for i := 0; i < int(level); i++ {
			fmt.Printf("%c", '-')
		}
		fmt.Println(msg)
	}
}

func Create(cfg config.LoggerConfig) (Logger, error) {
	return Logger{cfg.Level}, nil
}
