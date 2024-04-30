package contract

type Source interface {
	Open() (interface{}, error)
}

type Reader interface {
	Read(s Source) (interface{}, error)
}