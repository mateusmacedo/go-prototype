package err

import (
	"fmt"
	"strings"
)

func ErrorFactory(msg string, tokens ...interface{}) error {
	numberOfExpectedTokens := strings.Count(msg, "%")
	numberOfTokens := len(tokens)
	if numberOfExpectedTokens != numberOfTokens {
		return fmt.Errorf("invalid number of tokens, expected %d, got %d", numberOfExpectedTokens, numberOfTokens)
	}

	return fmt.Errorf(msg, tokens...)
}