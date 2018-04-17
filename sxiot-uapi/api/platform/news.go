package platform

import (
	"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

func CreateNews(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("news_title")
	content := r.FormValue("news_content")
	if title == ""{
		w.WriteHeader(400)
		return
	}
	id := uuid.NewV4().String()
	news :=&db.News{
		ID:id,
		Title:title,
		Content:content,
	}
	err:=db.AddNews(news)
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func DelNews(w http.ResponseWriter, r *http.Request) {
	id :=r.FormValue("news_id")
	err:=db.DelNews(id)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindNews(w http.ResponseWriter, r *http.Request) {
	news,err:=db.FindNews()
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(news))
}