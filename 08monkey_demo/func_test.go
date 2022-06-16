package monkey_demo

import (
	"golang-unit-test-example/varys"
	"strings"
	"testing"

	"bou.ke/monkey"
)

//  TestMyFunc 为函数打桩
func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入uid是多少,都返回 &varys.UserInfo{Name: "RandySun"}, nil

	monkey.Patch(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "RandySun"}, nil
	})

	res := MyFunc(123)

	if !strings.Contains(res, "RandySun") {
		t.Fatal()
	}
	t.Logf("name: %s", res)
}

func TestFooFunc(t *testing.T) {
	t.Log(Foo1("apple")) // 输出：i am apple, calling function foo1
	t.Log(Foo2("pear"))  // 输出：i am pear, calling function foo2

	monkey.Patch(Foo1, Foo2)
	t.Log(Foo1("apple")) //输出：i am apple, calling function foo2

	t.Log(Foo3()) // 输出：true
	t.Log(Foo4()) // 输出：false

	monkey.Patch(Foo3, Foo4)
	t.Log(Foo3()) //实际输出：true
}
