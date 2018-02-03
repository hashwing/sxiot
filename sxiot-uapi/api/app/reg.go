package app

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	logs.Error(r.Header.Get("Auth"))
	account := r.FormValue("user_account")
	password := r.FormValue("user_password")
	if account == ""||password == ""{
		w.WriteHeader(400)
		return
	}
	id := uuid.NewV4().String()
	user :=&db.PersonUser{
		UserID:id,
		UserAccount:account,
		UserPassword:password,
	}
	err:=db.AddUser(user)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}