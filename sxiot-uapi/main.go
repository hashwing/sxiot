package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-uapi/api"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/hashwing/sxiot/sxiot-core/common"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func run(){
	err := config.NewCommonConfig()
	if err != nil {
		panic(err)
	}
	config.SetLogConfig(config.UAPI_LOG_PATH)
	err=db.NewDB()
	if err!=nil{
		logs.Error(err)
		panic(err)
	}
	root := mux.NewRouter()
	api.NewRouter(root)
	http.ListenAndServe(":"+strconv.Itoa(config.CommonConfig.Platform.WebPort), root)
}

func main() {
	common.BackGroundService(config.UAPI_SERVICE_NAME,config.UAPI_SERVICE_DESC,nil,run)
}
