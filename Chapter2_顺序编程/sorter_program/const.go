package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 全局的变量
var (
	appname           = "sorter_program"
	v                 *viper.Viper
	log               *logrus.Logger
)

func InitConst() error {
	v = viper.New()
	log = logrus.New()

	if v == nil {
		return fmt.Errorf("创建全局 viper 对象失败")
	}

	if log == nil {
		return fmt.Errorf("创建全局 logger 对象失败")
	}

	return nil
}
