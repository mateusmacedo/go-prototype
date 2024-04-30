package service

import (
	"github.com/mateusmacedo/govibranium/prototype/internal/application/contract"
	"github.com/mateusmacedo/govibranium/prototype/internal/core/validation"
)

type OSFilePathSource struct {
	path string
	validators []validation.Validator
}

type OSFileSourceOption func(*OSFilePathSource) error

func OSFileSourceOptionsFunc(opts ...OSFileSourceOption) *OSFilePathSource {
	s := &OSFilePathSource{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithOSFileSourcePath(path string) OSFileSourceOption {
	return func(s *OSFilePathSource) error {
		s.path = path
		return nil
	}
}

func WithOSFileSourceValidators(validators ...validation.Validator) OSFileSourceOption {
	return func(s *OSFilePathSource) error {
		s.validators = validators
		return nil
	}
}

func NewOSFileSource(opts ...OSFileSourceOption) (contract.Source, error) {
	s := OSFileSourceOptionsFunc(opts...)
	for _, v := range s.validators {
		if err := v.Validate(s.path); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *OSFilePathSource) Open() (interface{}, error) {
	return s.path, nil
}

const (
	InvalidOSFileSourcePathErrMsg = "the path %s is invalid"
)
