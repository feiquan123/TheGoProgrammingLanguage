//读取配置文件,返回修改后的viper

package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 后期可以使用数组
func ReadConfigFile(configdir, configfilename, configfiletype string) {
	//读取app的配置文件
	v.AddConfigPath(configdir)
	v.SetConfigName(configfilename)
	v.SetConfigType(configfiletype)

	if err := v.ReadInConfig(); err != nil {
		log.SetFormatter(&logrus.JSONFormatter{})
		log.Error("Can not read app config file from dir : ", configdir, "/", configfilename, ".", configfiletype)
		os.Exit(1)
	}
}
