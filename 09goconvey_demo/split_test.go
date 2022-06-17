package goconvey_demo

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

//  TestSplit
func TestSplit(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = ":"
			expect = []string{"a", "b", "c"}
		)
		got := Split(s, sep)
		convey.So(got, convey.ShouldResemble, expect) // 断言
	})

	convey.Convey("不包含分隔符用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = "|"
			expect = []string{"a:b:c"}
		)
		got := Split(s, sep)
		convey.So(got, convey.ShouldResemble, expect) // 断言
	})
}

//  TestChildrenSplit 嵌套调用
func TestChildrenSplit(t *testing.T) {
	// 只需要在顶层的Convey调用时传入t
	convey.Convey("分隔符在开头或者结尾用例", t, func() {
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
			convey.Convey(tc.name, func() {
				// 嵌套调用Convery
				got := Split(tc.s, tc.sep)
				convey.So(got, convey.ShouldResemble, tc.expect)
			})
		}
	})
}

func TestStringSliceEqual(t *testing.T) {
	convey.Convey("TestStringSliceEqual", t, func() {
		convey.Convey("true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			convey.So(StringSliceEqual(a, b), convey.ShouldBeTrue)
		})

		convey.Convey("true when a ＝= nil  && b ＝= nil", func() {
			convey.So(StringSliceEqual(nil, nil), convey.ShouldBeTrue)
		})
	})
}
