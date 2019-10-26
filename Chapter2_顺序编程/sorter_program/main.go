package main

import (
	"fmt"
	"myGolang/TheGoProgrammingLanguage/Chapter2_顺序编程/sorter_program/algorithms/bubblesort"
	"myGolang/TheGoProgrammingLanguage/Chapter2_顺序编程/sorter_program/algorithms/qsort"
	"myGolang/TheGoProgrammingLanguage/Chapter2_顺序编程/sorter_program/common"
	"myGolang/TheGoProgrammingLanguage/Chapter2_顺序编程/sorter_program/sorter"
	"time"
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
	sorter.SetViperAndLog(v, log)
}

func main() {
	values, err := sorter.ReadValues(*Infile)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	t1 := time.Now()
	switch *Algorithm {
	case "qsort":
		qsort.QucikSort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		log.Infof("Sorting algorithm ", *Algorithm, " is either unknown or unsupported.")
	}
	t2 := time.Now()
	log.Infof("The sorting process costs %ds to complete",t2.Sub(t1))

	if err:=sorter.WriteValues(values, *Outfile);err !=nil{
		fmt.Errorf(err.Error())
	}
}
