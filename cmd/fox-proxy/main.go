package main

import (
	"fmt"
	"os"

	"github.com/NpoolPlatform/fox-proxy/pkg/servicename"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	cli "github.com/urfave/cli/v2"
)

const serviceName = servicename.ServiceName

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	err := app.Init(
		serviceName,
		description,
		"",
		"",
		"./",
		nil,
		commands,
		mysqlconst.MysqlServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		logger.Sugar().Errorw("main", "ServiceName", serviceName, "Error", err)
		return
	}
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorw("main", "Msg", "end of running", "ServiceName", serviceName, "Error", err)
	}
}
