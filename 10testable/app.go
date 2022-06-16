package testable

import (
	"github.com/op/go-logging"
	"math/rand"
	"time"
)

type App struct {
	Logger
}

func (a *App) Start() {
	a.Logger.Info("app start ...")
}

// NewApp 构造函数，将依赖项注入
func NewApp(lg Logger) *App {
	return &App{
		Logger: lg, // 使用传入的依赖项完成初始化
	}
}

// Logger 将日志库抽象为接口类型
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(460000)
	if num%2 == 0 {
		var log = logging.MustGetLogger("example")
		app := NewApp(log)
		app.Start()
	} else {
		app := &App{}
		app.Start()
	}
}
