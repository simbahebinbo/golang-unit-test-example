package goconvey_demo

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

//  TestSplit
func TestSplit(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = ":"
			expect = []string{"a", "b", "c"}
		)
		got := Split(s, sep)
		c.So(got, c.ShouldResemble, expect) // 断言
	})

	c.Convey("不包含分隔符用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = "|"
			expect = []string{"a:b:c"}
		)
		got := Split(s, sep)
		c.So(got, c.ShouldResemble, expect) // 断言
	})
}

//  TestChildrenSplit 嵌套调用
func TestChildrenSplit(t *testing.T) {
	// 只需要在顶层的Convey调用时传入t
	c.Convey("分隔符在开头或者结尾用例", t, func() {
		tt := []struct {
			name   string
			s      string
			sep    string
			expect []string
		}{
			{"分隔符在开头", "*1*2*3", "*", []string{"", "1", "2", "3"}},
			{"分隔符在结尾", "1+2+3+", "+", []string{"1", "2", "3", ""}},
		}
		for _, tc := range tt {
			c.Convey(tc.name, func() {
				// 嵌套调用Convery
				got := Split(tc.s, tc.sep)
				c.So(got, c.ShouldResemble, tc.expect)
			})
		}
	})
}

func TestStringSliceEqual(t *testing.T) {
	c.Convey("TestStringSliceEqual", t, func() {
		c.Convey("true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			c.So(StringSliceEqual(a, b), c.ShouldBeTrue)
		})

		c.Convey("true when a ＝= nil  && b ＝= nil", func() {
			c.So(StringSliceEqual(nil, nil), c.ShouldBeTrue)
		})
	})
}
