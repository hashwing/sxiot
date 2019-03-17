package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

type loginReply struct {
	Token string `json:"token"`
}

func JsonMsg(msg interface{}) []byte {
	data, _ := json.Marshal(msg)
	return data
}

type MyCustomClaims struct {
	UserID string
	jwt.StandardClaims
}

//login
func Login(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	password := r.FormValue("password")
	code := r.FormValue("code")
	if code == "" && (account == "" || password == "") {
		w.WriteHeader(400)
		w.Write([]byte("请求参数错误"))
		return
	}
	var user *db.PersonUser
	var res bool
	var err error
	openID := ""
	if code != "" {
		openID, err = getOpenID(code)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("code 验证失败"))
			return
		}
	}
	if account == "" {
		res, user = db.AuthOpenID(openID)
		if !res {
			w.WriteHeader(401)
			w.Write([]byte("用户未注册"))
			return
		}
	} else {
		res, user = db.AuthUser(account, password)
		if !res {
			w.WriteHeader(401)
			w.Write([]byte("账号密码错误"))
			return
		}
		err = db.UpdateUserOpenID(openID, account)
		if err != nil {
			logs.Error(err)
		}
	}

	expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()
	claims := MyCustomClaims{
		user.UserID,
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
	header := r.Header.Get("Auth")
	if header == "" {
		return errors.New("无权限")
	}
	token, err := jwt.ParseWithClaims(header, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.CommonConfig.Platform.JwtSecret), nil
	})
	if err != nil {
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

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "Claims").(*MyCustomClaims)
	if claims.UserID == "" {
		logs.Error("uid is null")
		w.WriteHeader(500)
		return
	}
	user, err := db.GetUser(claims.UserID)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(JsonMsg(user))
}
