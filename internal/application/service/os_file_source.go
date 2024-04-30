package service

import (
	"fmt"

	"github.com/mateusmacedo/govibranium/prototype/internal/application/contract"
)

type OSFileSource struct {
	path string
}

func NewOSFileSource(path string) contract.Source {
	return &OSFileSource{path: path}
}

func (s *OSFileSource) Open() (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}