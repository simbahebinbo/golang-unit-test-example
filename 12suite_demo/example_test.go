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

// 准备测试环境
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.VariableThatShouldStartAtFive = 5
}

// 恢复测试环境
func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
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
