package app

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func FindTopNews(w http.ResponseWriter, r *http.Request) {
	news,err:=db.FindTopNews()
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(news))
}