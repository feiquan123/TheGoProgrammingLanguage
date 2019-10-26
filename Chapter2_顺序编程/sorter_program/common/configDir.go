package common

// 获取配置文件的路径

import (
	"runtime"
)

func GetConfig(appName string)(appConfigDir,appConfigFileName string){
	switch runtime.GOOS {
	case "windows":
		{
			appConfigDir      = "E:/GoProgram/src/myGolang/TheGoProgrammingLanguage/Chapter2_顺序编程/sorter_program/config/"
			appConfigFileName = "app_win"
		}
	default:
		appConfigDir      = "/deployment/apps/"+appName+"/current/"
		appConfigFileName = "app"
	}
	return appConfigDir,appConfigFileName
}
