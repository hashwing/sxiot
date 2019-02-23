package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/hashwing/sxiot/sxiot-core/common"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/hashwing/sxiot/sxiot-core/db"
	"github.com/astaxie/beego/logs"
)

type loginReply struct {
	Token string `json:"token"`
}

func JsonMsg(msg interface{}) []byte {
	data, _ := json.Marshal(msg)
	return data
}

type MyCustomClaims struct {
//	Username           string `json:"username"`
	UserID             string    //`json:"-"`
//	IsAdmin            bool   `json:"admin"`
	jwt.StandardClaims //`json:"-"`
}

//login
func Login(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	password := r.FormValue("password")
	if account == "" || password == "" {
		w.WriteHeader(400)
		return
	}
	res,user:= db.AuthAdmin(account, password)
	if !res {
		w.WriteHeader(400)
		return
	}
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	claims := MyCustomClaims{
	//	user.UserAlias,
		user.UserID,
	//	true,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "sxiot",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.CommonConfig.Platform.JwtSecret))
	msg := loginReply{
		Token: signedToken,
	}
	w.Write(JsonMsg(msg))
}

//check token
func Auth(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		return errors.New("非法操作")
	}
	if cookie.String()=="Auth="{
		return errors.New("非法操作")
	}
	splitCookie := strings.Split(cookie.String(), "Auth=")
	token, err := jwt.ParseWithClaims(splitCookie[1], &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.CommonConfig.Platform.JwtSecret), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		context.Set(r, "Claims", claims)
	} else {
		return errors.New("非法操作")
	}
	return nil
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	user,err:=db.GetAdmin(claims.UserID)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(user))
	context.Clear(r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	account := r.FormValue("account")
	password := r.FormValue("pwd")
	uid := r.FormValue("uid")
	if uid!=""{
		user:=db.AdminUser{
			UserAccount:account,
			UserPassword:password,
			UserID:uid,
			UserAlias:name,
		}
		err:=db.UpdateAdmin(&user)
		if err!=nil{
			logs.Error(err)
			w.WriteHeader(500)
			return
		}
	}else{
		user:=db.AdminUser{
			UserAccount:account,
			UserPassword:password,
			UserID:common.NewUUID(),
			UserAlias:name,
		}
		err:=db.AddAdmin(&user)
		if err!=nil{
			logs.Error(err)
			w.WriteHeader(500)
			return
		}
	}
	w.WriteHeader(204)
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	err:=db.DelAdmin(uid)
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	users,err:=db.FindAdmins()
	if err!=nil{
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(users))
}