package base_demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	// "detartrated"是一个回文字符串，因此IsPalindrome("detartrated")的返回值应该为true
	// 如果返回false，则表示其实现有问题，需要使用t.Error进行错误报告
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}

	// 断言IsPalindrome方法的返回值为True
	assert.True(t, IsPalindrome("detartrated"))
	assert.True(t, IsPalindrome("kayak"))
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}

	// 断言IsPalindrome方法的返回值为False
	assert.False(t, IsPalindrome("palindrome"))
}
