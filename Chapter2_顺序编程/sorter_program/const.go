package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 提示
var (
	H         *bool   // 帮助
	Infile    *string // 输入文件
	Outfile   *string // 输出文件
	Algorithm *string // 算法
)

// 全局的变量
var (
	appName           = "sorter_program"
	appConfigFileType = "yaml"
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
