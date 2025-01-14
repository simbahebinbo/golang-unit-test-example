package gomock_demo

import (
	"golang-unit-test-example/06gomock_demo/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	// 创建gomock控制器， 用来记录后续操作信息
	ctrl := gomock.NewController(t)

	// 调用mockgen生成代码中的NewMockDB方法
	//  mockgen -source="db.go" -destination="mocks/db_mock.go" -package="mocks"
	// 这里mocks是我们生成代码时指定的package包名称
	m := mocks.NewMockDB(ctrl)
	// 打桩(stub)
	//  当传入Get函数的参数为RandySun时返回1和nil
	m.
		EXPECT().
		Get(gomock.Eq("RandySun")). // 参数
		Return(1, nil). // 返回值
		Times(1) // 调用次数

	// 调用GetFromDB函数时传入上面的mock对象m, v = Return(1, nil)返回的值
	if v := GetFromDB(m, "RandySun"); v != 1 {
		t.Fatal()
	}
}
