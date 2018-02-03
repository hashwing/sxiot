package main


import (
	"github.com/hashwing/sxiot/sxiot-dbagent/core"
	"github.com/hashwing/sxiot/sxiot-core/common"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/astaxie/beego/logs"
)


func run(){
	err:=config.SetLogConfig(config.DBAGENT_LOG_PATH)
	if err != nil {
		logs.Error(err)
	}
	err = config.NewCommonConfig()
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	core.NewMQClient()
}

func main(){
	common.BackGroundService(config.DBAGENT_SERVICE_NAME,config.DBAGENT_SERVICE_DESC,nil,run)
}