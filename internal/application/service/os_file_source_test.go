package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mateusmacedo/govibranium/prototype/test/mocks"
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
		t.Run("Test NewOSFileSource with invalid path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
    		defer ctrl.Finish()
			mockValidator := mocks.NewMockValidator(ctrl)
			mockValidator.EXPECT().Validate("test").Return(fmt.Errorf("error"))
			_, err := NewOSFileSource(
				WithOSFileSourcePath("test"),
				WithOSFileSourceValidators(mockValidator),
			)
			if err == nil {
				t.Error("Expected error, got nil")
			}
			if err.Error() != "error" {
				t.Errorf("Expected error to be error, got %s", err.Error())
			}
		})
	})
}