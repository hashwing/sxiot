package platform

import (
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/context"
	"github.com/hashwing/sxiot/sxiot-core/db"
)
func CreateGateway(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("gateway_name")
	if name == ""{
		w.WriteHeader(400)
		return
	}
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	id := uuid.NewV4().String()
	gateway :=&db.DeviceGateway{
		ID:id,
		AdminID:claims.UserID,
		Name:name,
	}
	err:=db.AddGateway(gateway)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func UpdateGateway(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("gateway_id")
	name := r.FormValue("gateway_name")
	if name == ""{
		w.WriteHeader(400)
		return
	}
	gateway :=&db.DeviceGateway{
		ID:id,
		Name:name,
	}
	err:=db.UpdateGateway(gateway)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindGateways(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	gateways,err:=db.FindGateways(claims.UserID)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(gateways))
}

func GetGateway(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("gateway_id")
	gateway,err:=db.GetGateway(id)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(gateway))
}

func DelGateway(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("gateway_id")
	err:=db.DelGateway(id)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}