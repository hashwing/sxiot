package emqtt

import (
	"io"
	"net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/hashwing/sxiot/sxiot-core/db"
	"github.com/hashwing/sxiot/sxiot-core/config"    
)

func BasicAuth(method, uri string, body io.Reader) ([]byte,error) {
    username := config.CommonConfig.MQTT.User
	passwd := config.CommonConfig.MQTT.Password
	api :=config.CommonConfig.MQTT.URL
    client := &http.Client{}
    req, err := http.NewRequest(method, api+uri, body)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        return nil,err
    }
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil{
        return nil,err
    }
    return bodyText,nil
}


func GetSubsByClinetID(id string)([]*db.Device,error){
    devices :=make([]*db.Device,0)
    data,err:=BasicAuth("GET","/api/v2/subscriptions/"+id,nil)
	if err!=nil{
		return nil,err
    }
    var res ClusterSubLists
    err=json.Unmarshal(data,&res)
    for _,v:=range res.Result.Objectds{
        if v.Topic==id{
            continue
        }
       d,err:=db.GetDevice(v.Topic)
       if err!=nil{
		    return nil,err
        }
       b,err:=db.GetBrand(d.BrandID)
       if err!=nil{
            return nil,err
        }
        d.BrandID=b.Type
       devices=append(devices,d)
    }
    return devices,nil
}

func GetClientStatus(id string)(bool,error){
    data,err:=BasicAuth("GET","/api/v2/clients/"+id,nil)
	if err!=nil{
		return false,err
    }
    if string(data)==`{"code":0,"result":{"objects":[]}}`{
        return false,nil
    }
    return true,nil
}

