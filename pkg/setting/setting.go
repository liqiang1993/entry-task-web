package setting

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	DeadlineSecond  int
	RuntimeRootPath string
	TokenExpireHour int
	MaxHeaderBytes  int

	ImageMaxSize  int
	ImageAllowExt []string

	LogLevel      string // 日志文件展示级别
	LogFileName   string // 日志文件存放路径与名称
	LogMaxSize    int    // 日志文件大小，单位是 MB
	LogMaxBackups int    // 最大过期日志保留个数
	LogMaxAgeDay  int    // 保留过期文件最大时间，单位 天
	LogCompress   bool   // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type DaoServer struct {
	GrpcPort int
}

var DaoServerSetting = &DaoServer{}

var cfg *ini.File

func InitConfig() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("setting.Init, fail to parse app.ini: %s", err)
		panic("read config file failed")
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("dao-server", DaoServerSetting)

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("Cfg.MapTo %s err: %s", section, err)
		panic("get config failed")
	}
}
