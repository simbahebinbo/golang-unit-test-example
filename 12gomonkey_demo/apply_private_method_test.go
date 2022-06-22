package gomonkey_demo

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/fake"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestApplyPrivateMethod(t *testing.T) {
	convey.Convey("TestApplyPrivateMethod", t, func() {
		convey.Convey("patch private pointer method in the different package", func() {
			f := new(fake.PrivateMethodStruct)
			var s *fake.PrivateMethodStruct
			patches := gomonkey.ApplyPrivateMethod(s, "ok", func(_ *fake.PrivateMethodStruct) bool {
				return false
			})
			defer patches.Reset()
			result := f.Happy()
			convey.So(result, convey.ShouldEqual, "unhappy")
		})

		convey.Convey("patch private value method in the different package", func() {
			s := fake.PrivateMethodStruct{}
			patches := gomonkey.ApplyPrivateMethod(s, "haveEaten", func(_ fake.PrivateMethodStruct) bool {
				return false
			})
			defer patches.Reset()
			result := s.AreYouHungry()
			convey.So(result, convey.ShouldEqual, "I am hungry")
		})
	})
}
