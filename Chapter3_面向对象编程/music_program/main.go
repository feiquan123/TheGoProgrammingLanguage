package main

import (

	"fmt"
	"myGolang/TheGoProgrammingLanguage/Chapter3_面向对象编程/music_program/common"
	"myGolang/TheGoProgrammingLanguage/Chapter3_面向对象编程/music_program/kernel"
)

func init() {
	// 初始化配置文件
	if err := InitConst(); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化提示
	tips()

	// 判断系统，选择不同的配置文件,和配置路径
	appConfigDir, appConfigFileName := common.GetConfig(appName)
	// 传入app.yaml 的文件,设置Viper和logger
	common.SetVpierAndLogger(v, log, appConfigDir, appConfigFileName, appConfigFileType)

	// 初始化其它文件
	kernel.SetViperAndLog(v, log)
}

func main() {
	log.Info("go run start program")

}
