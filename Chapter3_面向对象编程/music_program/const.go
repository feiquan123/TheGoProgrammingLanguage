package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"myGolang/TheGoProgrammingLanguage/Chapter3_面向对象编程/music_program/kernel/musicManager"
)

// 提示
var (
	H *bool // 帮助
)

// 全局的变量
var (
	appName           = "sorter_program"
	appConfigFileType = "yaml"
	v                 *viper.Viper
	log               *logrus.Logger
	lib               *musicManager.MusicManager // 全局的音乐库
	ctrl, singal      chan int                   // 全局控制信号
	id                = 1						 // 全局音乐库的id
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


	lib = musicManager.NewMusicManager()
	return nil
}
