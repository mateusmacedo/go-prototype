package contract

type Source interface {
	Open() (interface{}, error)
}

type SourceReader interface {
	Read(s Source) (interface{}, error)
}