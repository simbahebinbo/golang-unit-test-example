package monkey_demo

import (
	"fmt"
	"golang-unit-test-example/varys"
)

// MyFunc 获取用户名
func MyFunc(uid int64) string {
	u, err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里处理逻辑代码
	return fmt.Sprintf("hello %s\n", u.Name)
}

func Foo1(s string) string {
	return fmt.Sprintf("i am %s, calling function foo1", s)
}
func Foo2(s string) string {
	return fmt.Sprintf("i am %s, calling function foo2", s)
}

func Foo3() bool {
	return true
}

func Foo4() bool {
	return false
}
