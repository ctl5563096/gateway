package main

import (
	"context"
	"fmt"
	"gateway/bootstrap"
	"gateway/job/monitor"
	"github.com/ctl5563096/base/library/closable"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		isEndless bool
		err       error
	)
	// 初始化脚手架context
	ctx, cancel := context.WithCancel(context.Background())

	cliConf, cliCmd, err := bootstrap.GetInitOptions()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(101)
	}
	// 初始化所有组件
	if successInit, err := bootstrap.Init(ctx, cliCmd, cliConf); successInit != true {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(102)
	}

	// 开启协程阻塞监控
	go monitor.MonitorWaitGo()

	//监听退出信号
	go func() {
		if isEndless {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
			<-c
		}
		cancel()
	}()
	<-ctx.Done()
	//time.Sleep(5 * time.Second)
	closable.Done()
}
