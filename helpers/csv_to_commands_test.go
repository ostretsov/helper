package helpers

import "testing"

func Test_countPlaceholders(t *testing.T) {
	t.Run("0placeholders", func(t *testing.T) {
		result := placeholders("")
		if len(result) != 0 {
			t.Errorf("Expected 0 placeholders, got %d", result)
		}
	})
	t.Run("1placeholder", func(t *testing.T) {
		result := placeholders("/cmd arg1 ${1}")
		if result[0] != 1 {
			t.Errorf("Expected 1 placeholder, got %d", result)
		}
	})
	t.Run("2placeholder", func(t *testing.T) {
		result := placeholders("/cmd arg1 arg2 ${1} ${2}")
		if len(result) != 2 {
			t.Errorf("Expected 2 placeholders, got %d", result)
		}
	})
	t.Run("placeholdersGap", func(t *testing.T) {
		result := placeholders("/cmd arg1 arg2 ${1} ${3}")
		if len(result) != 2 {
			t.Errorf("Expected 1 placeholders, got %d", result)
		}
	})
}
