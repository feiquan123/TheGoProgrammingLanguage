package main

import (
	"flag"
	"fmt"
	"os"
)

// 设置默认的提示
func usage() {
	fmt.Fprintf(os.Stderr, `appname: sorter_program 
version: 1.0
Usage: sorter_program [-h help] [-i Infile] [ -o  Outfile ] [ -a Algorithm ( bubblesort | qsort )]

Options:
`)
	flag.PrintDefaults()
}

//  提示
func tips() {
	H = flag.Bool("h", false, "help")
	Infile = flag.String("i", "infile", "File contains values for sorting")
	Outfile = flag.String("o", "outfile", "File to receive sorted values")
	Algorithm = flag.String("a", "qsort", "Sort algorithm")

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

	if Infile == nil || *Infile==""|| Outfile ==nil || *Outfile ==""{
		log.Errorf("infile=", *Infile, "outfile=", *Outfile, "algorithm=", *Algorithm)
		os.Exit(1)
	}
}
