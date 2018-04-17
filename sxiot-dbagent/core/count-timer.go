package core

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/hashwing/sxiot/sxiot-core/emqtt"
)

var device int64
var user	int64

func CountTimer(){
	mustSave()
	t1:=time.NewTimer(time.Duration(10)*time.Second)
	t2:=time.NewTimer(time.Duration(3600)*time.Second)
    for {
        select {
			case <-t1.C:
				saveCount()
				t1=time.NewTimer(time.Duration(10)*time.Second)
			case <-t2.C:
				mustSave()
				t2=time.NewTimer(time.Duration(3600)*time.Second)
		}
    }
}

func saveCount(){
	devicen,usern,err:=emqtt.CountDevice()
	if err!=nil{
		return
	}
	if usern!=user{
		saveData("user",usern)
		user=usern
	}

	if device!=devicen{
		saveData("device",devicen)
		device=devicen
	}
}

func mustSave(){
	devicen,usern,err:=emqtt.CountDevice()
	if err!=nil{
		logs.Error(err)
		return
	}
	saveData("user",usern)
	saveData("device",devicen)
}


func saveData(id string,data interface{})error{
	if endpoint != nil {
		tag := map[string]string{
			"id":id,
		}
		field := map[string]interface{}{
			"data":data,
		}
		metric, err:= client.NewPoint("trends", tag, field, time.Now())
		if err!=nil{
			logs.Error(err)
			return err
		}
		batchpoint.AddPoint(metric)
		err=endpoint.Write(batchpoint)
		if err!=nil{
			logs.Error(err)
			return err
		}
	}
	return nil	
}