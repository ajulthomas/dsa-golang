package reverse_string

import (
	"fmt"
	"testing"

	"github.com/ajulthomas/dsa-golang/model"
)

type TestCases []model.TestCase

var testCases = TestCases{
	{TestInput: "apple", ExpectedOutput: "elppa"},
	{TestInput: "hello", ExpectedOutput: "olleh"},
}

func TestReverse(t *testing.T) {
	fmt.Println("Running Tests \n================= ")
	for index, testCase := range testCases {
		var output string = Reverse(testCase.TestInput)
		if output != testCase.ExpectedOutput {
			t.Errorf("Output %q not equal to expected %q", output, testCase.ExpectedOutput)
		}
		fmt.Printf("CASE %v ====> PASSED: Testcase: %v, Expected_Output: %v, Output: %v \n \n",
			index,
			testCase.TestInput,
			testCase.ExpectedOutput,
			output)
	}
}
