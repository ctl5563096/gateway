package bootstrap

import (
	"context"
	"gateway/config"
	"github.com/ctl5563096/base/library"
	"github.com/urfave/cli/v2"
)

func GetInitOptions() (map[string]interface{}, *library.CliCommand, error) {
	//cli 启动初始化
	cliCmd := library.NewCliCommand("base-service", "脚手架", "基础脚手架命令")
	cliCmd.AddConfig(config.LoadCommandConfigs()...)
	cliCmd.CustomInitCommand(InitCommand())
	err := cliCmd.Setup()
	if err != nil {
		return nil, nil, err
	}
	cliConf := cliCmd.ExecCommand().Options()
	return cliConf, cliCmd, nil
}

func Init(ctx context.Context, cliCmd *library.CliCommand, cliConf map[string]interface{}) (isEndless bool, err error) {

}

func InitCommand() (cliCommand *cli.Command, execCommand *library.ExecCommand) {
	execCommand = library.NewExecCommand()
	cliCommand = &cli.Command{
		Name:  "run",
		Usage: "run http service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   ".env.local",
				Aliases: []string{"c"},
			},
			&cli.IntFlag{
				Name:    "port",
				Value:   2345,
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:    "script",
				Value:   "",
				Aliases: []string{"s"},
			},
		},
		Action: func(context *cli.Context) error {
			execCommand.SetCurrentCmdName(context.Command.Name)
			execCommand.SetIsEndless(true)
			execCommand.SetOption("config", context.String("config"))
			execCommand.SetOption("port", context.Int("port"))
			execCommand.SetOption("script", context.String("script"))
			execCommand.SetOption("BRANCH", context.String("BRANCH"))
			return nil
		},
	}
	return cliCommand, execCommand
}
