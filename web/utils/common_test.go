package utils

import (
	"testing"
)


func TestValidOrDefault(t *testing.T) {
	var testCases = []struct{
		inputValue interface{}
		defaultValue interface{}
		expectResult interface{}
	}{
		{ nil, 1, 1 },
		{ 0, 1, 0 },
		{ 1, 2, 1 },
		{ true, false, true },
		{ false, true, false },
		{ "", "a", "a" },
		{ "a", "b", "a" },
		{ "", 999,  999 },
	}

	for _, testCase := range testCases {
		actualResult := ValidOrDefault(testCase.inputValue, testCase.defaultValue)
		if testCase.expectResult != actualResult {
			t.Errorf("test failed: except %v, got %v", testCase.expectResult, actualResult)
		}

	}
}
