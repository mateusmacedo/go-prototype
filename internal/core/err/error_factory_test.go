package err

import "testing"

func TestErrorFactory(t *testing.T){
	t.Run("Test ErrorFactory", func(t *testing.T){
		t.Run("Test ErrorFactory with valid tokens", func(t *testing.T){
			expectedErr := "this is a test"
			err := ErrorFactory("this is a %s", "test")
			if err.Error() != expectedErr {
				t.Errorf("Expected error to be %s, got %s", expectedErr, err.Error())
			}
		})
		t.Run("Test ErrorFactory with invalid tokens", func(t *testing.T){
			expectedErr := "invalid number of tokens, expected 1, got 0"
			emptyTokens := []interface{}{}
			err := ErrorFactory("this is a %s", emptyTokens...)
			if err.Error() != expectedErr {
				t.Errorf("Expected error to be %s, got %s", expectedErr, err.Error())
			}
		})
	})
}