package common

// 获取配置文件的路径

import (
	"runtime"
)

func GetConfig(appName string)(appConfigDir,appConfigFileName string){
	switch runtime.GOOS {
	case "windows":
		{
			appConfigDir      = "E:/GoProgram/src/myGolang/TheGoProgrammingLanguage/Chapter3_面向对象编程/music_program/config/"
			appConfigFileName = "app_win"
		}
	default:
		appConfigDir      = "/apps/"+appName+"/config/"
		appConfigFileName = "app"
	}
	return appConfigDir,appConfigFileName
}
