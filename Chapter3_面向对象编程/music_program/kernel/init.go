package kernel

import (
"github.com/sirupsen/logrus"
"github.com/spf13/viper"
)

var (
	v   *viper.Viper
	log *logrus.Logger
)

// 初始化viper和日志
func SetViperAndLog(viper *viper.Viper, logger *logrus.Logger) {
	v = viper
	log = logger
}

