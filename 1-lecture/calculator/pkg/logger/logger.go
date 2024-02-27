package logger

import "fmt"

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
