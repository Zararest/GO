package main

import (
	"calculator/pkg/parser"
	"calculator/pkg/logger"
)

func main() {
	parser.Parse("hello", logger.Logger{Level:1})
}
