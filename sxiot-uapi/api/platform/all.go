package platform


import (
	"io/ioutil"
	"net/http"
	log "github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/emqtt"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func EmqttCluster(w http.ResponseWriter, r *http.Request) {
	data,err:=emqtt.GetCluters()
	if err!=nil{
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func EmqttClients(w http.ResponseWriter, r *http.Request) {
	data,err:=emqtt.GetSessions()
	if err!=nil{
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func EmqttHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Error(err)
		w.WriteHeader(500)
		return
	}
	log.Info(string(body))
	w.WriteHeader(200)
}

func DeviceUser(w http.ResponseWriter, r *http.Request) {
	users,err:=db.CountUser()
	if err != nil {
        log.Error(err)
		w.WriteHeader(500)
		return
	}
	gateways,err:=db.CountGateway()
	if err != nil {
        log.Error(err)
		w.WriteHeader(500)
		return
	}
	deviceOn,userOn,err:=emqtt.CountDevice()
	if err != nil {
        log.Error(err)
		w.WriteHeader(500)
		return
	}
	devices,err:=db.CountDevice()
	msg:=map[string]int64{
		"user":users,
		"gateway":gateways,
		"device":devices,
		"gateway_online":deviceOn,
		"user_online":userOn,
	}
	w.Write(JsonMsg(msg))
}