package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/smile-im/auth-service/program"
)

var (
	VERSION      string // 程序版本
	GIT_HASH     string // git hash
	COMPILE_DATE string // 编译时间
	GO_VERSION   string // go 版本
)

func main() {
	// 系统日志显示文件和行号
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetFlags(log.Llongfile | log.LstdFlags)

	// 判断是否是-v参数如果是，则输出版本信息
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "version" {
			fmt.Printf("Version: %s\nGit Hash: %s\nBuild Date: %s\nGO Version: %s\n", VERSION, GIT_HASH, COMPILE_DATE, GO_VERSION)
			return
		}
	}

	// 创建程序实例
	p := program.New()
	p.Run()

	// 监听退出信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	sig := <-ch
	// 停止服务
	p.Stop()
	log.Println("Exit")
	if i, ok := sig.(syscall.Signal); ok {
		os.Exit(int(i))
	} else {
		os.Exit(0)
	}
}
