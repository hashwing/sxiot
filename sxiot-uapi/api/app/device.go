package app

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/emqtt"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func GetSonDevices(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("gateway_id")
	if id ==""{
		w.WriteHeader(400)
		return
	}
	data,err:=emqtt.GetSubsByClinetID(id)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(data))
}

func UpdateSonDevice(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("device_id")
	name:=r.FormValue("device_name")
	device :=&db.Device{
		ID:id,
		Name:name,
	}
	err:=db.UpdateDevice(device)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func GetAllDevices(w http.ResponseWriter, r *http.Request){
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	devices,err:=db.FindPersonDevices(claims.UserID)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}

	datas:=make([]map[string]interface{},0)
	for _,v:=range devices{
		clients,err:=emqtt.GetSubsByClinetID(v.DeviceID)
		if err!=nil{
			logs.Error(err)
			w.WriteHeader(500)
			return
		}
		data :=map[string]interface{}{
			"gateway_id":v.DeviceID,
			"son_device":clients,
		}
		datas=append(datas,data)
	}
	w.Write(JsonMsg(datas))
}