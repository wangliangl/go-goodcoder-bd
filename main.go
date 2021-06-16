package main

import (
	"flag"
	"fmt"
	"icode.baidu.com/baidu/go-lib/log"
	"icode.baidu.com/baidu/go-lib/log/log4go"
	"mini_spider/util"
)

const (
	//about script project
	scriptVersion = "1.0.0"
	projectName   = "mini_spider"

	//about log of baidu
	// when:
	//       "M", minute
	//       "H", hour
	//       "D", day
	//       "MIDNIGHT", roll over at midnight
	logWhen = "D"
	//       backupCount: If backupCount is > 0, when rollover is done, no more than
	logBackUp = 0
)

var (
	confPath *string = flag.String("c", "conf/spider.yaml", "mini spider configure path")
	help     *bool   = flag.Bool("h", false, "show help")
	logPath  *string = flag.String("l", "log", "dir path of log")
	stdout   *bool   = flag.Bool("s", false, "show log in stdout")
	version  *bool   = flag.Bool("v", false, "show version")
	debug    *bool   = flag.Bool("d", false, "show debug log")
)

func main() {

	flag.Parse()
	if *version {
		fmt.Printf("version is :%s\n", scriptVersion)
		return
	}

	if *help {
		flag.PrintDefaults()
		return
	}

	var logLevel string

	if *debug {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}

	// init baidu log
	log4go.SetLogFormat(log4go.FORMAT_DEFAULT_WITH_PID)
	err := log.Init(projectName, logLevel, *logPath, *stdout, logWhen, logBackUp)
	if err != nil {
		fmt.Printf("baidu log init err:%s\n", err.Error())
		//todo 优雅退出
	}

	// init script config
	conf, err := util.LoadConfig(*confPath)
	if err != nil {
		_ = log.Logger.Error("load config err:%s", err.Error())
		return
	}
	fmt.Println(conf.Spider)

	// log url data


}
