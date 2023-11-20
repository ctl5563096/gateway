package config

import (
	"gateway/providers"
	"github.com/ctl5563096/base/library"
	"os"
)

var EnvConfigs []*library.EnvConfig

func InitEnvConfigs(cliConfig map[string]interface{}) {
	fileName, _ := cliConfig["config"].(string)
	EnvConfigs = []*library.EnvConfig{
		{
			Receivers:       &providers.Env,
			FileType:        "env",
			FileName:        fileName,
			FilePath:        GetExecDirectory(),
			EnableReReading: false, //配置热加载支持
		},
	}
}

// GetExecDirectory 获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err == nil {
		return file + "/"
	}
	return ""
}
