package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type App struct {
	AppName      string
	AppDebug     bool
	AppAesSecret string
	JwtSecret    string

	// 日志文件
	AppAccessLogName string
	AppErrorLogName  string
	AppGrpcLogName   string
	LogSavePath      string

	// Database
	DbConnection string
	DbHost       string
	DbPort       string
	DbDatabase   string
	DbUsername   string
	DbPassword   string
	DbCharset    string
	TablePrefix  string

	// Server
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TsGap        int64 // 接口时效性 3000ms

	//Redis
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// Setup initialize the configuration instance
func Setup() {
	var AppConfig App
	err := envconfig.Process("myapp", &AppConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Debug: %s\nPort: %s\nUser: %s\nRate: %s\n"
	_, err = fmt.Printf(format, AppConfig.DbDatabase, AppConfig.DbHost, AppConfig.DbUsername, AppConfig.DbConnection)
	if err != nil {
		log.Fatal(err.Error())
	}
}
