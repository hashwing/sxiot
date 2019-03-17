package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/logs"
)

type wxReturn struct {
	Openid     string `json:"openid"`
	Errcode    int    `json:"errcode"`
	SessionKey string `json:"session_key"`
	ErrMsg     string `json:"errmsg"`
	Unionid    string `json:"unionid"`
}

func getOpenID(code string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", "wx4ec2de710b34f252", "57d1f3559d979e6580160a8619671733", code)
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("Couldn't parse response body. %+v", err)
		return "", err
	}
	var r wxReturn
	err = json.Unmarshal(body, &r)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	if r.Errcode != 0 {
		logs.Error("auth code error: %s", r.ErrMsg)
		return "", fmt.Errorf("auth code error: %s", r.ErrMsg)
	}
	return r.Openid, nil
}
