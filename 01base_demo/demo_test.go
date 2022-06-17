package base_demo

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func square(op int) int {
	return op * op
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}

// 利用+=连接
func TestConcatStringByAdd(t *testing.T) {
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal(t, "12345", ret)
}

// 利用buffer连接
func TestConcatStringBytesBuffer(t *testing.T) {
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal(t, "12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
}

func TestSomething(t *testing.T) {
	// 断言两者相等
	assert.Equal(t, 123, 123)

	// 断言两者不相等
	assert.NotEqual(t, 123, 456)
}
