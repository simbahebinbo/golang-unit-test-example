package base_demo

import "testing"

type testData struct {
	input          interface{}
	expectedOutput interface{}
}

func myTestFail(
	t *testing.T,
	testCase testData,
	actualOutput interface{},
	index int) {

	if actualOutput != testCase.expectedOutput.(bool) {
		t.Errorf("\n\ncase %+v:", index)
		t.Errorf("input = %+v", testCase.input)
		t.Errorf("expected output = %+v", testCase.expectedOutput)
		t.Errorf("actual output = %+v", actualOutput)
	}
}

func TestIsValidHostName(t *testing.T) {
	testCaseList := []testData{
		// 正向case,每行是一个case
		{"hostaa", true},
		{"hostbb", true},
		{"hostcc", true},

		//负向case,每行是一个case
		{"Hostaa", false},
		{"host123", false},
		{"host!", false},

		// 边界case，每行是一个case
		{"host", true},
		{"hostabcd", true},
		{"hos", false},
		{"hostabcde", false},
	}

	for index, testCase := range testCaseList {
		actualOutput := IsValidHostName(testCase.input.(string))
		myTestFail(t, testCase, actualOutput, index)
	}
}
