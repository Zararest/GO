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

type parsingState struct {
	Tokens   []token
	Position int
}

func (state parsingState) isEndOfTokens() bool {
	return state.Position >= len(state.Tokens) 
}

func (state parsingState) getToken() token {
	return state.Tokens[state.Position]
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
	log.Print("Tokenizing starts" /*level*/, 1)
	words := strings.Fields(str)
	result := make([]token, 0)
	for _, word := range words {
		result = append(result, getToken(word, log))
	}
	log.Print("Tokenizing ends" /*level*/, 1)
	return result
}

func parseVal(state *parsingState, log logger.Logger) (float64, error) {
	log.Print("parsing value ["+state.getToken().Str+"]", /*level*/2)
	if !state.getToken().IsValue {
		return 0, errors.New("Invalid expression: "+state.getToken().Str)
	}
	ans := state.getToken().Value
	state.Position++
	return ans, nil
}

func parseSub(state *parsingState, log logger.Logger) (float64, error) {
	log.Print("parsing sub on token["+strconv.FormatInt(int64(state.Position), /*base*/10), /*level*/2)
	lhs, err := parseVal(state, log)
	if err != nil {
		return 0, err
	}
	ans := lhs

	for !state.isEndOfTokens() {
		if state.getToken().Str != "-" {
			return ans, nil
		}
		state.Position++

		rhs, err := parseVal(state, log)
		if err != nil {
			return 0, nil
		}
		ans -= rhs
	}

	return ans, nil

}

func parseSum(state *parsingState, log logger.Logger) (float64, error) {
	log.Print("parsing sum on token["+strconv.FormatInt(int64(state.Position), /*base*/10), /*level*/2)
	lhs, err := parseSub(state, log)
	if err != nil {
		return 0, err
	}
	ans := lhs

	for !state.isEndOfTokens() {
		if state.getToken().Str != "+" {
			return ans, nil
		}
		state.Position++

		rhs, err := parseSub(state, log)
		if err != nil {
			return 0, nil
		}
		ans += rhs
	}

	return ans, nil
}

func calculate(tokens []token, log logger.Logger) (float64, error) {
	log.Print("Calculation starts", /*level*/1)
	state := parsingState{Tokens: tokens, Position: 0}
	ans, err := parseSum(&state, log)
	return ans, err
}

// Main package function to parse mathematical expressions.
// supported operations:
//
//	-; +; *; /
func Parse(input string, log logger.Logger) float64 {
	log.Print("Parsing starts on a line:" , /*level*/1)
	log.Print(input, /*level*/1)
	tokens := tokenize(input, log)
	res, err := calculate(tokens, log)
	if err != nil {
		panic(err)
	}
	return res
}
