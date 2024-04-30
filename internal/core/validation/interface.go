package validation

import "fmt"

type Validator interface {
	Validate(target interface{}) error
}

var (
	ErrInvalidValidator = fmt.Errorf("invalid validator")
)