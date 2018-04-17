package emqtt

import (
	"strings"
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
    data,err:=BasicAuth("GET","/api/v2/subscriptions/device_"+id,nil)
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
    data,err:=BasicAuth("GET","/api/v2/clients/device_"+id,nil)
	if err!=nil{
		return false,err
    }
    if string(data)==`{"code":0,"result":{"objects":[]}}`{
        return false,nil
    }
    return true,nil
}

func GetNodes()(*ClusterNodes,error){
    data,err:=BasicAuth("GET","/api/v2/monitoring/nodes",nil)
    if err!=nil{
		return nil,err
    }
    var res ClusterNodes
    err=json.Unmarshal(data,&res)
    return &res,err
}

func CountDevice()(int64,int64,error){
    data,err:=BasicAuth("GET","/api/v2/monitoring/nodes",nil)
    if err!=nil{
		return 0,0,err
    }
    var res ClusterNodes
    err=json.Unmarshal(data,&res)
    if err!=nil{
		return 0,0,err
    }
    var DeviceSum int64 = 0
    var UserSum    int64 = 0
    for _,v:=range res.Result{
        clientsbyte,err:=BasicAuth("GET","/api/v2/nodes/"+v.Name+"/clients",nil)
        if err!=nil{
            return 0,0,err
        }
        var clients Clients
        err=json.Unmarshal(clientsbyte,&clients)
        if err!=nil{
		    return 0,0,err
        }
        for _,c:=range clients.Result.Objects{
            if strings.HasPrefix(c.ID,"device_"){
                DeviceSum++
            }
            if strings.HasPrefix(c.ID,"app_"){
                UserSum++
            }
        }
        
    }
    return DeviceSum,UserSum,err
}

func GetCluters()([]byte,error){
    return BasicAuth("GET","/api/v2/management/nodes",nil)
}

func GetSessions()([]byte,error){
    return BasicAuth("GET","/api/v2/monitoring/nodes",nil)
}
