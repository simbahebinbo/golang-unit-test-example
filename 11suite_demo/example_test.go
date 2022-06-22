package suite_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// 在测试套件启动前执行一次
func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("SetupSuite 在测试套件启动前执行一次")
}

// 在测试套件用例都执行完成
func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite 在测试套件用例都执行完成")
}

// 在每个用例执行前执行一次
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest 在每个用例执行前执行一次")
	suite.VariableThatShouldStartAtFive = 5
}

// 在每个用例执行后执行一次
func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest 在每个用例执行后执行一次")
}

// 测试案例
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// 执行ExampleTestSuite结构体里面的全部测试函数
func TestExampleTestSuite(t *testing.T) {
	// 运行ExampleTestSuite中Test开始的函数
	suite.Run(t, new(ExampleTestSuite))
}
