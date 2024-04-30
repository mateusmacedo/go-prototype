package service

import (
	"github.com/mateusmacedo/govibranium/prototype/internal/application/contract"
)

type OSFileReader struct {
	adapter contract.Reader
}

type OSFileReaderOption func(*OSFileReader) error


func WithOSFileReaderAdapter(adapter contract.Reader) OSFileReaderOption {
	return func(s *OSFileReader) error {
		s.adapter = adapter
		return nil
	}
}

func OSFileReaderOptionsFunc(opts ...OSFileReaderOption) *OSFileReader {
	s := &OSFileReader{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func NewOSFileReader(opts ...OSFileReaderOption) contract.Reader {
	return OSFileReaderOptionsFunc(opts...)
}

func (s *OSFileReader) Read(source contract.Source) (interface{}, error) {
	return s.adapter.Read(source)
}
