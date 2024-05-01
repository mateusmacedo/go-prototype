package service

import (
	"github.com/mateusmacedo/govibranium/prototype/internal/application/contract"
	"github.com/mateusmacedo/govibranium/prototype/internal/core/helper"
	"github.com/mateusmacedo/govibranium/prototype/internal/core/validation"
)

type OSFileSource struct {
	path string
	validators []validation.Validator
	adapter helper.Adapter
}

type OSFileSourceOption func(*OSFileSource) error

func OSFileSourceOptionsFunc(opts ...OSFileSourceOption) *OSFileSource {
	s := &OSFileSource{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithOSFileSourcePath(path string) OSFileSourceOption {
	return func(s *OSFileSource) error {
		s.path = path
		return nil
	}
}

func WithOSFileSourceValidators(validators ...validation.Validator) OSFileSourceOption {
	return func(s *OSFileSource) error {
		s.validators = validators
		return nil
	}
}

func WithOSFileSourceAdapter(adapter helper.Adapter) OSFileSourceOption {
	return func(s *OSFileSource) error {
		s.adapter = adapter
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

func (s *OSFileSource) Open() (interface{}, error) {
	file, err := s.adapter.Adapt(s.path)
    if err != nil {
        return nil, err
    }

    return file, nil
}

const (
	InvalidOSFileSourcePathErrMsg = "the path %s is invalid"
)
