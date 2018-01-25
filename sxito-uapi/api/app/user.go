package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/hashwing/sxito/sxito-core/config"
	"github.com/hashwing/sxito/sxito-core/db"
)

type loginReply struct {
	Token string `json:"token"`
}

func JsonMsg(msg interface{}) []byte {
	data, _ := json.Marshal(msg)
	return data
}

type MyCustomClaims struct {
	UserID             string
	jwt.StandardClaims 
}

//login
func Login(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	password := r.FormValue("password")
	if account == "" || password == "" {
		w.WriteHeader(400)
		return
	}
	res,user := db.AuthUser(account, password)
	if !res {
		w.WriteHeader(400)
		return
	}
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	claims := MyCustomClaims{
		user.UserID,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "sxito",
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
	header:=r.Header.Get("Auth")
	if header==""{
		return errors.New("无权限")
	}
	token, err := jwt.ParseWithClaims(header, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.CommonConfig.Platform.JwtSecret), nil
	})
	if err!=nil{
		return err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		context.Set(r, "Claims", claims)
	} else {
		return errors.New("非法操作")
	}
	return nil
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	w.Write(JsonMsg(claims))
	context.Clear(r)
}
