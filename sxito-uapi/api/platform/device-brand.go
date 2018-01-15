package platform

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"github.com/hashwing/sxito/sxito-core/db"
)

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("brand_name")
	metadata := r.FormValue("brand_metadata")
	if name == ""{
		w.WriteHeader(400)
		return
	}
	id := uuid.NewV4().String()
	brand :=&db.DeviceBrand{
		ID:id,
		Name:name,
		Metadata:metadata,
	}
	err:=db.AddBrand(brand)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("brand_id")
	name := r.FormValue("brand_name")
	metadata := r.FormValue("brand_metadata")
	if name == ""{
		w.WriteHeader(400)
		return
	}
	brand :=&db.DeviceBrand{
		ID:id,
		Name:name,
		Metadata:metadata,
	}
	err:=db.UpdateBrand(brand)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindBrands(w http.ResponseWriter, r *http.Request) {
	brands,err:=db.FindBrands()
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(brands))
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("brand_id")
	brand,err:=db.GetBrand(id)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(brand))
}

func DelBrand(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("brand_id")
	err:=db.DelBrand(id)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}