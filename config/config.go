package config

import (
	"fmt"
	"visibleBase/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

// 返回给前端的数据结构
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数初始化此结构体并赋值默认值
// 我们敲定tokne有问题那么data的status==401
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	returnData.Message = "你还没有配置返回信息"
	return returnData
}

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 //jwt过期时间，单位分钟
	Username   string
	Password   string

	//Incluster 相关配置
	MetaDataNameSpace       string                //元数据存储namespace
	InclusterKubeConfigPath string                //incluster kubeconfig路径
	InClusterClinetSet      *kubernetes.Clientset //clientgo客户端工具
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

// 只要是以init开头的函数都会在程序启动时自动执行，调包的时候也是
func init() {
	logs.Info(nil, "开始加载程序配置")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JWT_SIGN_KEY", "MoriDreamer")
	//过期时间120000分钟，即2天
	viper.SetDefault("JWT_EXPIRE_TIME", "120000")
	viper.SetDefault("USERNAME", "Mori")
	viper.SetDefault("PASSWORD", "10086")
	viper.SetDefault("METADATA_NAMESPACE", "visible-k8s")
	viper.AutomaticEnv()
	Port = ":" + viper.GetString("PORT")           //获取端口的配置
	logLevel := viper.GetString("LOG_LEVEL")       //获取日志输出的配置
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")   //获取JWT签名密钥的配置
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME") //获取JWT过期时间的配置
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	MetaDataNameSpace = viper.GetString("METADATA_NAMESPACE")
	//加载日志输出格式
	initLogConfig(logLevel)

}
