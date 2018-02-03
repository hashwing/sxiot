package app

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"github.com/gorilla/context"
	"github.com/hashwing/sxiot/sxiot-core/db"
)
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	gid :=r.FormValue("gateway_id")
	alias := r.FormValue("device_alias")
	if alias == ""|| gid ==""{
		w.WriteHeader(400)
		return
	}
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	id := uuid.NewV4().String()
	device :=&db.PersonDevice{
		ID:id,
		UserID:claims.UserID,
		DeviceID:gid,
		Alias:alias,
	}
	logs.Info(claims.UserID)
	err:=db.AddPersonDevice(device)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func UpdateDevice(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("id")
	gid :=r.FormValue("gateway_id")
	alias := r.FormValue("device_alias")
	if alias == ""|| id==""|| gid==""{
		w.WriteHeader(400)
		return
	}
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	device :=&db.PersonDevice{
		ID:id,
		UserID:claims.UserID,
		DeviceID:gid,
		Alias:alias,
	}
	err:=db.UpdatePersonDevice(device)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindDevices(w http.ResponseWriter, r *http.Request){
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	devices,err:=db.FindPersonDevices(claims.UserID)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(devices))
}

func GetDevcie(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("device_id")
	if id==""{
		w.WriteHeader(400)
		return
	}
	device,err:=db.GetDevice(id)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(device))
}

func DelDevice(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("device_id")
	if id==""{
		w.WriteHeader(400)
		return
	}
	err:=db.DelPersonDevice(id)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}