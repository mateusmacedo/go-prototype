package service

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mateusmacedo/govibranium/prototype/internal/core/err"
	"github.com/mateusmacedo/govibranium/prototype/test/mocks"
)

func TestOSFileSource(t *testing.T) {
	t.Run("Test NewOSFileSource", func(t *testing.T) {
		t.Run("Test NewOSFileSource with valid path", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
    		defer ctrl.Finish()
			expValidate := interface{}(nil)
			mockValidator := mocks.NewMockValidator(ctrl)
			mockValidator.EXPECT().Validate("test").Return(expValidate)
			mockAdapter := mocks.NewMockSource(ctrl)
			_, err := NewOSFileSource(
				WithOSFileSourcePath("test"),
				WithOSFileSourceValidators(mockValidator),
				WithOSFileSourceAdapter(mockAdapter),
			)
			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}
		})
		t.Run("Test NewOSFileSource with invalid path", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
    		defer ctrl.Finish()
			expValidate := err.ErrorFactory(InvalidOSFileSourcePathErrMsg, "test")
			mockValidator := mocks.NewMockValidator(ctrl)
			mockValidator.EXPECT().Validate("test").Return(expValidate)
			mockAdapter := mocks.NewMockSource(ctrl)
			_, err := NewOSFileSource(
				WithOSFileSourcePath("test"),
				WithOSFileSourceValidators(mockValidator),
				WithOSFileSourceAdapter(mockAdapter),
			)
			if err == nil {
				tr.Error("Expected error, got nil")
			}
			if err.Error() != expValidate.Error() {
				tr.Errorf("Expected error to be %s, got %s", expValidate, err)
			}
		})
	})
	t.Run("Test OSFileSource Open", func(t *testing.T) {
		t.Run("Test OSFileSource Open with valid adapter", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
    		defer ctrl.Finish()
			expValidate := interface{}(nil)
			mockValidator := mocks.NewMockValidator(ctrl)
			mockValidator.EXPECT().Validate("test").Return(expValidate)
			expOpen := interface{}(nil)
			mockAdapter := mocks.NewMockSource(ctrl)
			mockAdapter.EXPECT().Open().Return(nil, expOpen)
			s, err := NewOSFileSource(
				WithOSFileSourcePath("test"),
				WithOSFileSourceValidators(mockValidator),
				WithOSFileSourceAdapter(mockAdapter),
			)
			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}

			p, err := s.Open()
			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}

			if p != expOpen {
				tr.Errorf("Expected %v, got %v", expOpen, p)
			}
		})
		t.Run("Test OSFileSource Open with invalid adapter", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
			defer ctrl.Finish()
			expValidate := interface{}(nil)
			mockValidator := mocks.NewMockValidator(ctrl)
			mockValidator.EXPECT().Validate("test").Return(expValidate)
			expOpen := err.ErrorFactory("error %s", "test")
			mockAdapter := mocks.NewMockSource(ctrl)
			mockAdapter.EXPECT().Open().Return(nil, expOpen)
			s, err := NewOSFileSource(
				WithOSFileSourcePath("test"),
				WithOSFileSourceValidators(mockValidator),
				WithOSFileSourceAdapter(mockAdapter),
			)
			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}

			p, err := s.Open()
			if err == nil {
				tr.Error("Expected error, got nil")
			}
			if err.Error() != expOpen.Error() {
				tr.Errorf("Expected error to be %s, got %s", expOpen, err)
			}
			if p != nil {
				tr.Errorf("Expected nil, got %v", p)
			}
		})
	})
}