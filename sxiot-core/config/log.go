package config

import (
	"os"
	"github.com/astaxie/beego/logs"
)
func SetLogConfig(path string)error{
	os.MkdirAll(path,0664)
	return logs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+path+`/log.log","separate":["error"]}`)
}