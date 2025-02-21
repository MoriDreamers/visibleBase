package config

import (
	"JWT-TEST/utils/logs"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 //jwt过期时间，单位分钟
)

func initLogConfig(logLevel string) {
	//配置日志的输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: TimeFormat})
	fmt.Println("日志初始化完成！")
}

func init() {
	logs.Info(nil, "开始加载程序配置")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JWT_SIGN_KEY", "MORI")
	viper.SetDefault("JWT_EXPIRE_TIME", "120")
	viper.AutomaticEnv()
	Port = ":" + viper.GetString("PORT")           //获取端口的配置
	logLevel := viper.GetString("LOG_LEVEL")       //获取日志输出的配置
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")   //获取JWT签名密钥的配置
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME") //获取JWT过期时间的配置
	//加载日志输出格式
	initLogConfig(logLevel)

}
