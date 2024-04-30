package validation

type Validator interface {
	Validate(target interface{}) error
}
