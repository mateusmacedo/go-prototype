package helper

type Adapter interface {
	Adapt(t interface{}) (interface{}, error)
}