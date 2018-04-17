package platform

import (
	"strings"
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/db"
	"github.com/hashwing/sxiot/sxiot-core/config"
)

func AuthGateway(w http.ResponseWriter, r *http.Request){
	clientID :=r.FormValue("clientid")
	username :=r.FormValue("username")
	password :=r.FormValue("password")
	defaultU:=config.CommonConfig.MQTT.CUser
	defaultP:=config.CommonConfig.MQTT.CPassword
	logs.Debug(clientID,username)
	if strings.HasPrefix(clientID,"agent_")&&username==defaultU&&password==defaultP{
		w.WriteHeader(200)
		return
	}
	if !strings.HasPrefix(clientID,"app_")&&!strings.HasPrefix(clientID,"device_"){
		w.WriteHeader(401)
		return
	}
	clientID =strings.Replace(clientID,"app_","",-1)
	clientID =strings.Replace(clientID,"device_","",-1)
	device,err:=db.GetGateway(clientID)
	if err!=nil||device==nil{
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(200)
}

func AuthSuper(w http.ResponseWriter, r *http.Request){
	clientID :=r.FormValue("clientid")
	username :=r.FormValue("username")
	logs.Debug(clientID,username)
	defaultU:=config.CommonConfig.MQTT.CUser
	if strings.HasPrefix(clientID,"agent_")&&username==defaultU{
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(401)
}

func AuthDevice(w http.ResponseWriter, r *http.Request){
	clientID :=r.FormValue("clientid")
	username :=r.FormValue("username")
	topic :=r.FormValue("topic")
	defaultU:=config.CommonConfig.MQTT.CUser
	if strings.HasPrefix(clientID,"agent_")&&username==defaultU{
		w.WriteHeader(200)
		return
	}
	if username==""||topic==""{
		w.WriteHeader(401)
		return
	}
	res:=db.AuthDevice(topic,username)
	if !res{
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(200)
}