package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mateusmacedo/govibranium/prototype/test/mocks"
)

func TestOSFileReader(t *testing.T) {
	t.Run("Test NewOSFileReader", func(t *testing.T) {
		r := NewOSFileReader()

		if r == nil {
			t.Error("Expected OSFileReader, got nil")
		}
	})
	t.Run("Test OSFileReader Read", func(t *testing.T) {
		t.Run("Test OSFileReader Read with valid source", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
			defer ctrl.Finish()
			mockSource := mocks.NewMockSource(ctrl)
			mockSource.EXPECT().Open().Return(nil, nil)
			r := NewOSFileReader()

			_, err := r.Read(mockSource)
			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}
		})
		t.Run("Test OSFileReader Read with invalid source", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
			defer ctrl.Finish()
			mockSource := mocks.NewMockSource(ctrl)
			expectedErr := fmt.Errorf("error")
			mockSource.EXPECT().Open().Return(nil, expectedErr)
			r := NewOSFileReader()

			_, err := r.Read(mockSource)

			if err == nil {
				tr.Errorf("Expected error to be %s, got nil", expectedErr)
			}
		})
	})
}
