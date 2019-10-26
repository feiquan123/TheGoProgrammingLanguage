package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//viper和logger初始化，只进行修改不进行创建
var (
	v   *viper.Viper
	log *logrus.Logger
)

// 设置Viper
func SetViper(viper *viper.Viper) {
	v = viper
}

// 设置logger
func SetLog(logger *logrus.Logger) {
	log = logger
}

// 传入app.yaml 的文件,设置Viper和logger
func SetVpierAndLogger(viper *viper.Viper, logger *logrus.Logger, configdir, configfilename, configfiletype string) {
	if viper == nil {
		fmt.Println("viper is nil, can not common init")
		return
	}
	if logger == nil {
		fmt.Println("logger is nil, can not common init")
		return
	}
	if configdir == "" {
		fmt.Println("coonfig path is null, can not read config file")
	}
	if configfilename == "" {
		fmt.Println("config filename is null , can not read config file")
	}
	if configfiletype == "" {
		fmt.Println("config type is null, can not read config file")
	}

	// 设置Viper
	SetViper(viper)
	// 读取配置文件
	ReadConfigFile(configdir, configfilename, configfiletype)
	// 设置logger
	SetLog(logger)
	// 通过viper设置logger
	setLogFromViper()
}
