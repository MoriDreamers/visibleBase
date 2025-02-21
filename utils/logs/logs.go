package logs

import "github.com/sirupsen/logrus"

//打印debug类型的日志

func Debug(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Debug(msg)
}
func Info(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Info(msg)
}
func Warning(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Warning(msg)
}
func Erorr(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Error(msg)
}
