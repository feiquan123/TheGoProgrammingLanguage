package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 提示
var (
	H          *bool   // 帮助
	Infile     *string // 输入文件
	Outfile    *string // 输出文件
	Algorithm  *string // 算法
	Configfile *string // 配置文件
)

// 设置默认的提示
func usage() {
	fmt.Fprintf(os.Stderr, `appname: %s
version: 1.0
Usage: sorter_program [-h Help] [-i Infile] [ -o  Outfile ] [ -a Algorithm ( bubblesort | qsort )] [-c Configfile]

Options:
`, appname)
	flag.PrintDefaults()
}

//  提示
func tips() {
	H = flag.Bool("h", false, "help")
	Infile = flag.String("i", "infile", "File contains values for sorting")
	Outfile = flag.String("o", "outfile", "File to receive sorted values")
	Algorithm = flag.String("a", "qsort", "Sort algorithm")
	Configfile = flag.String("c", "", "Config file. must get")

	// 设置默认提示
	flag.Usage = usage

	// 参数解析
	setFalg()
}

// 参数解析
func setFalg() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *H {
		flag.Usage()
		os.Exit(1)
	}

	if Configfile == nil || *Configfile == "" {
		log.Error("config file not exist. Please add [ -c config_filename ] args.")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	if Infile == nil || *Infile == "" || Outfile == nil || *Outfile == "" {
		log.Errorf("infile=", *Infile, "outfile=", *Outfile, "algorithm=", *Algorithm)
		os.Exit(1)
	}
}

// 根据配置文件路径解析对应的参数
func getConfig(configFile string) (appConfigDir, appConfigFileName, appConfigFileType string) {
	// 相对路径转绝对路径
	configFile, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatal("config file path convert to abs path error. Error:", err)
	}

	// 将 // 转化为 \
	configFile = strings.ReplaceAll(configFile, "\\", "/")
	if _, err := os.Open(configFile); err != nil {
		log.Fatalf("config file path error, Error: %s", err)
	}
	arr := strings.Split(configFile, "/")

	// 获取配置文件的路径
	appConfigDir = strings.Join(arr[0:len(arr)-1], "/") + "/"

	// 获取配置文件名和文件类型
	configFile = arr[len(arr)-1]
	arr = strings.Split(configFile, ".")
	appConfigFileName = strings.Join(arr[0:len(arr)-1], ".")
	appConfigFileType = arr[len(arr)-1]

	if appConfigDir == "" || appConfigFileName == "" || appConfigFileType == "" {
		log.Fatal("config file path can not parse")
	}

	return appConfigDir, appConfigFileName, appConfigFileType
}
