package service

import "github.com/mateusmacedo/govibranium/prototype/internal/application/contract"

type OSFileReader struct {
}

type OSFileReaderOption func(*OSFileReader) error

func OSFileReaderOptionsFunc(opts ...OSFileReaderOption) *OSFileReader {
	s := &OSFileReader{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func NewOSFileReader(opts ...OSFileReaderOption) contract.SourceReader {
	return OSFileReaderOptionsFunc(opts...)
}

func (s *OSFileReader) Read(source contract.Source) (interface{}, error) {
	_, err := source.Open()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
