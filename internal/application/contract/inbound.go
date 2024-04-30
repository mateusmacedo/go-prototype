package contract

type Source interface {
	Open() (interface{}, error)
}