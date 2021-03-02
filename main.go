package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"yuque-webhook-wecom/config"
	"yuque-webhook-wecom/server"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `yuque-webhook-wecom
Usage: yuque-webhook-wecom [-h help] [-c ./config.yaml]
Options:
`)
	flag.PrintDefaults()
}

func main() {

	flag.StringVar(&filePath, "c", "./config.yaml", "配置文件所在")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		usage()
		return
	}

	// 设置配置文件和静态变量
	config.SetConfig(filePath)
	// Echo instance
	ins, err := server.CreateEngine()
	if err != nil {
		log.Printf("create engine err is %v\n", err)
		return
	}

	err = ins.Start(":8887")
	if err != nil {
		log.Printf("start server err is %v\n", err)
		return
	}

}
