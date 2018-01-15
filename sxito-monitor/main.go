package main

import (

	"github.com/hashwing/sxito/sxito-core/config"
	"github.com/hashwing/sxito/sxito-core/common"
	"github.com/hashwing/sxito/sxito-monitor/core"
)

func run(){
	err := config.NewCommonConfig()
	if err != nil {
		panic(err)
	}
	config.SetLogConfig(config.MONITOR_LOG_PATH)
	core.RunMonitor()
}

func main() {
	common.BackGroundService(config.MONITOR_SERVICE_NAME,config.MONITOR_SERVICE_DESC,nil,run)
}
