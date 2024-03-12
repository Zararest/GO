package parser

import (
	"calculator/pkg/logger"
	"errors"
	"regexp"
	"strconv" // for string convertion
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

func (state parsingState) getToken() token { // FIXME add checker
	return state.Tokens[state.Position]
}

func getToken(str string, log logger.Logger) token {
	log.Print("Creating token from:"+str /*level*/, 3)
	val, err := strconv.ParseFloat(str /*bitsize*/, 64)
	if err != nil {
		return token{IsValue: false, Value: 0, Str: str}
	}
	floatString := strconv.FormatFloat(val, 'f' /*prec*/, -1 /*bitSize*/, 64)
	log.Print("\tgot value: "+floatString /*level*/, 3)
	return token{IsValue: true, Value: val, Str: str}
}

func tokenize(str string, log logger.Logger) []token {
	log.Print("Tokenizing starts" /*level*/, 1)
	re := regexp.MustCompilePOSIX("(([0-9]*[.])?[0-9]+)|[-,+,/,*]|[(,)]") // FIXME
	words := re.FindAllString(str /*n*/, -1)
	result := make([]token, 0)
	for _, word := range words {
		result = append(result, getToken(word, log))
	}
	log.Print("Tokenizing ends" /*level*/, 1)
	return result
}

func parseVal(state *parsingState, log logger.Logger) (float64, error) {
	log.Print("parsing value ["+state.getToken().Str+"]" /*level*/, 2)
	if !state.getToken().IsValue {
		log.Print("Invalid token: "+state.getToken().Str /*level*/, 2)
		return 0, errors.New("invalid expression: " + state.getToken().Str)
	}
	ans := state.getToken().Value
	state.Position++
	return ans, nil
}

func parseExpr(state *parsingState, log logger.Logger) (float64, error) {
	log.Print("parsing expression on token: "+state.getToken().Str /*level*/, 2)
	if state.getToken().Str != "(" {
		return parseVal(state, log)
	}
	state.Position++
	ans, err := parseSum(state, log)

	if state.getToken().Str != ")" {
		return 0, errors.New("there is no )")
	}
	state.Position++
	return ans, err
}

func parseDiv(state *parsingState, log logger.Logger) (float64, error) {
	positionString := strconv.FormatInt(int64(state.Position) /*base*/, 10)
	log.Print("parsing div on token["+positionString+"]" /*level*/, 2)
	lhs, err := parseExpr(state, log)
	if err != nil {
		return 0, err
	}
	ans := lhs

	for !state.isEndOfTokens() {
		if state.getToken().Str != "/" {
			return ans, nil
		}
		state.Position++

		rhs, err := parseExpr(state, log)
		if err != nil {
			return 0, err
		}

		if rhs == 0 {
			return 0, errors.New("dividing by zero")
		}

		ans /= rhs
	}

	return ans, nil
}

func parseMul(state *parsingState, log logger.Logger) (float64, error) {
	positionString := strconv.FormatInt(int64(state.Position) /*base*/, 10)
	log.Print("parsing mul on token["+positionString+"]" /*level*/, 2)
	lhs, err := parseDiv(state, log)
	if err != nil {
		return 0, err
	}
	ans := lhs

	for !state.isEndOfTokens() {
		if state.getToken().Str != "*" {
			return ans, nil
		}
		state.Position++

		rhs, err := parseDiv(state, log)
		if err != nil {
			return 0, err
		}
		ans *= rhs
	}

	return ans, nil
}

func parseSub(state *parsingState, log logger.Logger) (float64, error) {
	positionString := strconv.FormatInt(int64(state.Position) /*base*/, 10)
	log.Print("parsing sub on token["+positionString+"]" /*level*/, 2)
	lhs, err := parseMul(state, log)
	if err != nil {
		return 0, err
	}
	ans := lhs

	log.Print("got sub" /*level*/, 2)
	for !state.isEndOfTokens() {
		if state.getToken().Str != "-" {
			return ans, nil
		}
		state.Position++

		rhs, err := parseMul(state, log)
		if err != nil {
			return 0, err
		}
		ans -= rhs
	}

	return ans, nil
}

func parseSum(state *parsingState, log logger.Logger) (float64, error) {
	positionString := strconv.FormatInt(int64(state.Position) /*base*/, 10)
	log.Print("parsing sum on token["+positionString+"]" /*level*/, 2)
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
			return 0, err
		}
		ans += rhs
	}

	return ans, nil
}

func calculate(tokens []token, log logger.Logger) (float64, error) {
	log.Print("Calculation starts" /*level*/, 1)
	state := parsingState{Tokens: tokens, Position: 0}
	ans, err := parseSum(&state, log)
	if err == nil && !state.isEndOfTokens() {
		return 0, errors.New("thre is extra symbols in the input")
	}
	return ans, err
}

// Main package function to parse mathematical expressions.
// supported operations:
//
//	-; +; *; /
func Parse(input string, log logger.Logger) float64 {
	log.Print("Parsing starts on a line:" /*level*/, 1)
	log.Print(input /*level*/, 1)
	tokens := tokenize(input, log)
	res, err := calculate(tokens, log)
	if err != nil {
		panic(err)
	}
	return res
}
