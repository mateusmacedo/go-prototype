package service

import (
	"testing"
)

func TestOSFileSource(t *testing.T) {
	t.Run("Test NewOSFileSource", func(t *testing.T) {
		t.Run("Test NewOSFileSource with valid path", func(t *testing.T) {
			s := NewOSFileSource("valid_path")
			if s == nil {
				t.Error("Expected NewOSFileSource to return a non-nil value")
			}
		})
		t.Run("Test NewOSFileSource with empty path", func(t *testing.T) {
			s := NewOSFileSource("")
			if s != nil {
				t.Error("Expected NewOSFileSource to return a nil value")
			}
		})
		t.Run("Test NewOSFileSource with invalid path", func(t *testing.T) {
			s := NewOSFileSource("invalid_path")
			if s != nil {
				t.Error("Expected NewOSFileSource to return a nil value")
			}
		})
	})
}