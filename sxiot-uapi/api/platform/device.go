package platform

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/common"
	"github.com/gorilla/context"
	"github.com/hashwing/sxiot/sxiot-core/db"
	"github.com/hashwing/sxiot/sxiot-core/etcd"
)
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	brandID :=r.FormValue("brand_id")
	name := r.FormValue("device_name")
	unit := r.FormValue("device_unit")
	if name == ""{
		w.WriteHeader(400)
		return
	}
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	id :=common.NewUUID()
	device :=&db.Device{
		ID:id,
		AdminID:claims.UserID,
		BrandID:brandID,
		Name:name,
		Unit:unit,
	}
	err:=db.AddDevice(device)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	ec,err:=etcd.NewEtcd()
	if err!=nil{
		logs.Error(err)
	}
	err=ec.PutDataDevice(id)
	if err!=nil{
		logs.Error(err)
	}
	w.WriteHeader(204)
}

func UpdateDevice(w http.ResponseWriter, r *http.Request){
	id :=r.FormValue("device_id")
	brandID :=r.FormValue("brand_id")
	name := r.FormValue("device_name")
	unit := r.FormValue("device_unit")
	if name == ""|| id==""{
		w.WriteHeader(400)
		return
	}
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	device :=&db.Device{
		ID:id,
		AdminID:claims.UserID,
		BrandID:brandID,
		Name:name,
		Unit:unit,
	}
	err:=db.UpdateDevice(device)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindDevcies(w http.ResponseWriter, r *http.Request){
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	devices,err:=db.FindDevices(claims.UserID)
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
	err:=db.DelDevice(id)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	ec,err:=etcd.NewEtcd()
	if err!=nil{
		logs.Error(err)
	}
	err=ec.DeleteDataDevice(id)
	if err!=nil{
		logs.Error(err)
	}
	w.WriteHeader(204)
}