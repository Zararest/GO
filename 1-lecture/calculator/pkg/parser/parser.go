package parser

import (
	"calculator/pkg/logger"
	"errors"
	"strconv" // for string convertion
	"strings" // for basic strings operations
)

type token struct {
	IsValue bool
	Value   float64
	Str     string
}

func getToken(str string, log logger.Logger) token {
	log.Print("Creating token from:"+str /*level*/, 3)
	val, err := strconv.ParseFloat(str /*bitsize*/, 64)
	if err != nil {
		return token{IsValue: false, Value: 0, Str: str}
	}
	log.Print("\tgot value: "+strconv.FormatFloat(val, 'f' /*prec*/, -1 /*bitSize*/, 64) /*level*/, 3)
	return token{IsValue: true, Value: val, Str: str}
}

func tokenize(str string, log logger.Logger) []token {
	log.Print("Tokenizing starts", 1)
	words := strings.Fields(str)
	result := make([]token, 0)
	for _, word := range words {
		result = append(result, getToken(word, log))
	}
	log.Print("Tokenizing ends", 1)
	return result
}

func calculate(tokens []token, log logger.Logger) (float64, error) {
	log.Print("Calculation starts", 1)
	_ = tokens
	return 0, errors.New("Not implemented")
}

// Main package function to parse mathematical expressions.
// supported operations:
//
//	-; +; *; /
func Parse(input string, log logger.Logger) float64 {
	log.Print("Parsing starts", 1)
	tokens := tokenize(input, log)
	res, err := calculate(tokens, log)
	if err != nil {
		panic(err)
	}
	return res
}
