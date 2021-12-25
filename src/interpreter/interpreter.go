package chefgi

import (
	"regexp"
	"strconv"
	"strings"
)

func Tokenize(source string) []int {
	rawTokens := strings.Split(source, "\n")
	tokens := make([]int, 0)

	matchNumberPattern, _ := regexp.Compile(`(\d)+`)

	for _, token := range rawTokens {
		matches := matchNumberPattern.MatchString(token)

		if matches {
			parsedToken := strings.ReplaceAll(token, "\t", "")
			item := strings.Split(parsedToken, " ")[0]
			parsed, _ := strconv.Atoi(item)
			tokens = append(tokens, parsed)
		}
	}

	return tokens
}

func ParseTokensToString(source []int) string {
	var parsed string
	for _, item := range source {
		parsed += string(rune(item))
	}
	return parsed
}
