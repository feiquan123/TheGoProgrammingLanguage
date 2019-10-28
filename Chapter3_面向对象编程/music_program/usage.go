package main

import (
	"flag"
	"fmt"
	"myGolang/TheGoProgrammingLanguage/Chapter3_面向对象编程/music_program/kernel/musicManager"
	"os"
	"strconv"
)

// 设置默认的提示
func usage() {
	fmt.Fprintf(os.Stderr, `appname: %s 
version: 1.0
Usage: sorter_program [-h help]

Options:
	- lib list : 
			View the existing music lib
	- lib add <name> <artist> <style> <source> <type> : 
			Add a music to the music lib 
	- lib remove <name>:
			Remove the spectified music from the lib
	- play <name>:
			Play the spectified music
`,appname)
	flag.PrintDefaults()
}

//  提示
func tips() {
	H = flag.Bool("h", false, "help")

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
}

// 音库管理命令
func handleLibCommands(tokens []string){
	switch tokens[1] {
	case "list":
		for i :=0;i<lib.Len();i++{
			e,_ := lib.Get(i)
			fmt.Println(i+1,":",e.Name,e.Artist,e.Style,e.Source,e.Type)
		}
	case "add":
		if len(tokens) == 7{
			id ++
			lib.Add(&musicManager.MusicEntry{strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens[5],tokens[6]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <style> <source> <type>")
		}
	case "remove":
		if len(tokens) ==3 {
			lib.RemoveByName()
		}

	}
}

