package config

import (
	"context"
	"github.com/ctl5563096/base/library"
)

var CommandConfigs []*library.CommandConfig

func LoadCommandConfigs() []*library.CommandConfig {
	CommandConfigs = []*library.CommandConfig{
		{
			Signature:   "hello:test {--say=:说了什么，必填项} {--lang=:用啥语言，必填项}",
			Description: "脚本功能测试 执行./eframe_demo script hello:test --say=xxxx --lang=xxxx可以查看结果",
			IsEndless:   false, //直接完脚本后是否退出
			HandleFunc: func(ctx context.Context, cmd *library.ExecCommand) error {
				return nil
			},
		},
	}
	return CommandConfigs
}
