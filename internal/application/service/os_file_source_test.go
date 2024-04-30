package service

import (
	"testing"
)

func TestOSFileSource(t *testing.T) {
	t.Run("Test NewOSFileSource", func(t *testing.T) {
		t.Run("Test NewOSFileSource with valid path", func(t *testing.T) {
			s, err := NewOSFileSource(WithOSFileSourcePath("test"))
			if err != nil {
				t.Errorf("Expected no error, got %s", err)
			}
			if s.path != "test" {
				t.Errorf("Expected path to be test, got %s", s.path)
			}
		})
	})
}