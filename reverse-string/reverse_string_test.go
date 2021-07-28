package main

import "testing"

type TestCase struct {
	testInput      string
	expectedOutput string
}

var testCases = []TestCase{
	{"apple", "elppa"}, {"hello", "olleh"},
}

func TestReverse(t *testing.T) {
	for _, testCase := range testCases {
		if output := Reverse(testCase.testInput); output != testCase.expectedOutput {
			t.Errorf("Output %q not equal to expected %q", output, testCase.expectedOutput)
		}
	}
}
