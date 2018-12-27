package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type app struct {
	JwtSecret string
	PageSize  int

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type server struct {
	RunMode      string
	HTTPPort     int `ini:"HTTPPort"`
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
}

type redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type captcha struct {
	Height        int
	Width         int
	Len           int
	Mode          int
	UseSimpleFont bool
	DotNoise      bool
	SlimeLine     bool
}

var (
	// AppSetting is app setting
	AppSetting = &app{}
	// ServerSetting is server setting
	ServerSetting = &server{}
	// DatabaseSetting is database setting
	DatabaseSetting = &database{}
	// RedisSetting is redis setting
	RedisSetting = &redis{}
	// CaptchaSetting is captcha setting
	CaptchaSetting = &captcha{}

	cfg *ini.File
	err error
)

// Setup is init function
func Setup() {
	cfg, err = ini.Load("conf/app.ini")
	if nil != err {
		log.Fatalln("setting.Setup, fail to parse 'conf/app.ini'", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
	mapTo("captcha", CaptchaSetting)
}

func mapTo(section string, v interface{}) {
	err = cfg.Section(section).MapTo(v)
	if nil != err {
		log.Fatalln("cfg.mapto err", err)
	}
}
