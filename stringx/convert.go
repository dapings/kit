package stringx

import (
	"strings"
	"unicode"
)

// CamelToSnake converts a camelCase string to snake_case.
// if the input string is not camelCase, it will return the lower original string.
// UserAgent -> user_agent
// User -> user
func CamelToSnake(camelCase string) string {
	var result strings.Builder
	for i, char := range camelCase {
		if unicode.IsUpper(char) && i > 0 {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(char))
	}
	return result.String()
}

// BigCamelToSmallCamel converts a BigCamelCase string to smallCamelCase.
// If not BigCamelCase, the lower first letter of original string.
// UserAgent -> userAgent
func BigCamelToSmallCamel(bigCamelCase string) string {
	if len(bigCamelCase) == 0 {
		return bigCamelCase
	}

	return strings.ToLower(bigCamelCase[:1]) + bigCamelCase[1:]
}

// CapitalizeFirstLetter capitalizes the first letter of the input string.
// userAgent -> UserAgent
func CapitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}

	firstChar := []rune(input)[:1]
	firstCharUpper := string(unicode.ToUpper(firstChar[0]))

	return firstCharUpper + input[1:]
}
