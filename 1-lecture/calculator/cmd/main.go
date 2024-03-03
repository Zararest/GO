package main

import (
	"bufio"
	"calculator/pkg/logger"
	"calculator/pkg/parser"
	"flag"
	"fmt"
	"os"
)

func main() {
	var logLevel uint
	flag.UintVar(&logLevel, "log-level", 0, "The level of logging details")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter expression:")
	expression, _ := reader.ReadString('\n')
	ans := parser.Parse(expression, logger.Logger{Level: logLevel})
	fmt.Println("Answer: ", ans)
}
