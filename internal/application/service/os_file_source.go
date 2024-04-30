package service

type OSFileSource struct {
	path string
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

func NewOSFileSource(opts ...OSFileSourceOption) (*OSFileSource, error) {
	s := OSFileSourceOptionsFunc(opts...)
	return s, nil
}