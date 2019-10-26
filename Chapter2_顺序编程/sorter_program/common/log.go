package common

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func setLogFromViper(){
	//接收配置文件中的参数
	leve := v.GetString("log.level")
	displayLine := v.GetBool("log.displayLine")
	isJson := v.GetBool("log.isJson")
	console := v.GetBool("log.console")
	path := v.GetString("log.path")
	file := v.GetString("log.file")

	//验证接收的参数
	if err := Validate(path, file, leve); err != nil {
		log.Error("Init logrus from config file, Error: %v", err)
	}

	// 设置等级
	switch strings.ToUpper(leve) {
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARNING":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	case "PANIC":
		log.SetLevel(logrus.PanicLevel)
	}

	// 打印行号
	if displayLine {
		// 会打印文件的绝对路径和行号	太长
		//log.SetReportCaller(v.GetBool("log.displayLine"))

		// 会打印文件的相对路径和行号
		filenameHook := filename.NewHook()
		filenameHook.Field = "file" //自定义字段
		log.AddHook(filenameHook)
	}

	// 打印格式
	if isJson {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	}

	// 定义输出的地方
	if console {
		log.SetOutput(os.Stdout)
	} else {
		// 创建日志文件夹,并指定输出
		os.MkdirAll(path, os.ModePerm)
		// 创建输出对象
		filepath := path + file
		logfile, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			log.Errorf("Create log file: %s ,Error: %s", filepath, err)
		}
		// 修改文件权限
		os.Chmod(filepath, 0666)
		log.Out = logfile
	}
}