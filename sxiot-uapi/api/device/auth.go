package device

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/astaxie/beego/logs"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hashwing/sxiot/sxiot-core/common"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/hashwing/sxiot/sxiot-core/db"
	"github.com/hashwing/sxiot/sxiot-core/emqtt"
)

func JsonMsg(msg interface{}) []byte {
	data, _ := json.Marshal(msg)
	return data
}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	rurl := r.FormValue("redirect_uri")
	clientID := r.FormValue("client_id")
	state := r.FormValue("state")
	v := url.Values{}
	v.Add("code", clientID)
	v.Add("state", state)
	body := v.Encode()
	target := fmt.Sprintf("%s&%s", rurl, body)
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	token := common.GetToken(config.CommonConfig.Platform.JwtSecret, clientID)
	data := map[string]interface{}{
		"access_token":  token,
		"refresh_token": token,
		"expires_in":    17600000,
	}
	w.Write(JsonMsg(data))
}

type requestBody struct {
	Header  header     `json:"header"`
	Payload payloadReq `json:"payload"`
}

type header struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	MessageID      string `json:"messageId"`
	PayLoadVersion int    `json:"payLoadVersion"`
}

type payloadReq struct {
	AccessToken string `json:"accessToken"`
	DeviceId    string `json:"deviceId"`
	Attribute   string `json:"attribute"`
	Value       string `json:"value"`
}

type devicesResp struct {
	Header  header         `json:"header"`
	Payload devicesPayload `json:"payload"`
}

type devicesPayload struct {
	Devices []device `json:"devices"`
}

type device struct {
	DeviceId   string      `json:"deviceId"`
	DeviceType string      `json:"deviceType"`
	DeviceName string      `json:"deviceName"`
	Brand      string      `json:"brand"`
	Model      string      `json:"model"`
	Icon       string      `json:"icon"`
	Properties []propertie `json:"properties"`
	Actions    []string    `json:"actions"`
}

type propertie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Device(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error("read body err, %v", err)
		w.WriteHeader(400)
		return
	}
	println("json:", string(body))

	var rData requestBody
	if err = json.Unmarshal(body, &rData); err != nil {
		logs.Error("Unmarshal err, %v", err)
		w.WriteHeader(400)
		return
	}
	uid, err := common.GetUID(rData.Payload.AccessToken, config.CommonConfig.Platform.JwtSecret)
	if err != nil {
		logs.Error("auth token err, %v", err)
		w.WriteHeader(400)
		return
	}
	switch rData.Header.Namespace {
	case "AliGenie.Iot.Device.Discovery":
		ds, err := findDevices(uid)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		rData.Header.Name = "DiscoveryDevicesResponse"
		dresp := devicesResp{
			Header:  rData.Header,
			Payload: devicesPayload{Devices: ds},
		}
		logs.Debug(string(JsonMsg(dresp)))
		w.Write(JsonMsg(dresp))
		break
	case "AliGenie.Iot.Device.Control":
		if rData.Payload.Attribute != "powerstate" {
			rData.Header.Name += "Response"
			errMsg := errMessage{
				Header: rData.Header,
				Payload: errPayload{
					DeviceId:  rData.Payload.DeviceId,
					ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
					Message:   "device not support",
				},
			}
			w.WriteHeader(500)
			w.Write(JsonMsg(errMsg))
			return
		}
		err := controlDevice(rData.Payload.DeviceId, rData.Payload.Value)
		if err != nil {
			logs.Error(err)
			errMsg := errMessage{
				Header: rData.Header,
				Payload: errPayload{
					DeviceId:  rData.Payload.DeviceId,
					ErrorCode: "SERVICE_ERROR",
					Message:   "服务器出错啦",
				},
			}
			w.WriteHeader(500)
			w.Write(JsonMsg(errMsg))
			return
		}
		msg := successMessage{
			Header: rData.Header,
			Payload: successPayload{
				DeviceId: rData.Payload.DeviceId,
			},
		}
		w.Write(JsonMsg(msg))
		break
	case "AliGenie.Iot.Device.Query":
		errMsg := errMessage{
			Header: rData.Header,
			Payload: errPayload{
				DeviceId:  rData.Payload.DeviceId,
				ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
				Message:   "device not support",
			},
		}
		w.WriteHeader(500)
		w.Write(JsonMsg(errMsg))
		break
	default:
		w.WriteHeader(400)
	}

}

func findDevices(uid string) ([]device, error) {
	res := make([]device, 0)
	gateways, err := db.FindPersonDevices(uid)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	logs.Debug(uid, gateways)
	for _, g := range gateways {
		data, err := emqtt.GetSubsByClinetID(g.DeviceID)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		for _, d := range data {
			dv := device{
				DeviceId:   d.ID,
				DeviceType: "light",
				DeviceName: g.Alias + d.Name,
				Brand:      "随心物联",
				Model:      g.Alias + d.Name,
				Icon:       "https://raw.githubusercontent.com/hashwing/sxiot-h5-app/master/img/logo.png",
				Properties: []propertie{propertie{Name: "powerstate", Value: "on"}},
				Actions:    []string{"TurnOn", "TurnOff"},
			}
			res = append(res, dv)
		}
	}
	return res, nil
}

func controlDevice(deviceID string, value string) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.CommonConfig.MQTT.MqttURL)
	opts.SetClientID("agent_" + common.NewUUID())
	opts.SetUsername(config.CommonConfig.MQTT.CUser)
	opts.SetPassword(config.CommonConfig.MQTT.CPassword)
	logs.Debug(config.CommonConfig.MQTT.CUser)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	defer c.Disconnect(250)
	text := "0"
	if value == "on" {
		text = "100"
	}
	if token := c.Publish(deviceID, 0, false, text); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

type errMessage struct {
	Header  header     `json:"header"`
	Payload errPayload `json:"payload"`
}

type errPayload struct {
	DeviceId  string `json:"deviceId"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

type successMessage struct {
	Header  header         `json:"header"`
	Payload successPayload `json:"payload"`
}

type successPayload struct {
	DeviceId string `json:"deviceId"`
}
