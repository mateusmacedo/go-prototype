package service

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mateusmacedo/govibranium/prototype/internal/core/err"
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
		t.Run("Test OSFileReader Read from valid source", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
			defer ctrl.Finish()
			mockSource := mocks.NewMockSource(ctrl)
			expRead := "test"
			mockReader := mocks.NewMockReader(ctrl)
			mockReader.EXPECT().Read(mockSource).Return(expRead, nil)

			r := NewOSFileReader(
				WithOSFileReaderAdapter(mockReader),
			)

			d, err := r.Read(mockSource)

			if err != nil {
				tr.Errorf("Expected no error, got %s", err)
			}

			if d != expRead {
				tr.Errorf("Expected %s, got %s", expRead, d)
			}
		})
		t.Run("Test OSFileReader Read from invalid source", func(tr *testing.T) {
			ctrl := gomock.NewController(tr)
			defer ctrl.Finish()
			mockSource := mocks.NewMockSource(ctrl)
			expRead := interface{}(nil)
			expReadErr := err.ErrorFactory("error %s", "test")
			mockReader := mocks.NewMockReader(ctrl)
			mockReader.EXPECT().Read(mockSource).Return(expRead, expReadErr)

			r := NewOSFileReader(
				WithOSFileReaderAdapter(mockReader),
			)

			d, err := r.Read(mockSource)

			if err == nil {
				tr.Error("Expected error, got nil")
			}

			if err.Error() != expReadErr.Error() {
				tr.Errorf("Expected %s, got %s", expReadErr, err)
			}

			if d != expRead {
				tr.Errorf("Expected %v, got %v", expRead, d)
			}
		})
	})
}
