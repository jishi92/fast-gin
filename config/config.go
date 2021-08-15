package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server    server
	Database  database
	App       app
	LogSeting logSeting
}

type server struct {
	Mode         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type database struct {
	Dialect      string
	User         string
	Password     string
	Host         string
	Name         string
	Protocol     string
	Charset      string
	ParseTime    string
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
	TablePrefix  string
}

type app struct {
	JwtSecret       int
	PrefixUrl       string
	ImageSavePath   string
	ImageAllowExts  []string
	ExportSavePath  string
	QrCodeSavePath  string
	RuntimeRootPath string
	ImageMaxSize    int
}

type logSeting struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  int
}

var Cfg config

// init 初始化 Cfg 全局变量。
func init() {
	envFile := "config/config.toml"
	if _, err := toml.DecodeFile(envFile, &Cfg); err != nil {
		log.Fatal(err)
	}
}
