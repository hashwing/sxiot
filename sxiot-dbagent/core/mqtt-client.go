package core

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/satori/go.uuid"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/hashwing/sxiot/sxiot-core/config"
	"github.com/hashwing/sxiot/sxiot-core/etcd"
)

var endpoint client.Client
var batchpoint client.BatchPoints


var f mqtt.MessageHandler = func(cli mqtt.Client, msg mqtt.Message) {	
	if endpoint != nil {
		tag := map[string]string{
			"__name__":"device",
			"device_id":msg.Topic(),
		}
		field := map[string]interface{}{
			"f64":msg.Payload(),
		}
		logs.Info(msg.Topic())
		metric, err:= client.NewPoint("_", tag, field, time.Now())
		if err!=nil{
			logs.Error(err)
			return
		}
		batchpoint.AddPoint(metric)
		err=endpoint.Write(batchpoint)
		if err!=nil{
			logs.Error(err)
		}
	}
}

func NewInflux()error{
	var err error
	endpoint, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.CommonConfig.Influxdb.URL,
		Username: config.CommonConfig.Influxdb.UserName,
		Password: config.CommonConfig.Influxdb.Password,
	})
	if err != nil {
		return err
	}
	batchpoint, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.CommonConfig.Influxdb.DbName,
		Precision: "s",
	})
	return err
}

func NewMQClient()error{
	err:=NewInflux()
	if err!=nil{
		return err
	}
	go CountTimer()
	
	opts :=mqtt.NewClientOptions()
	opts.AddBroker(config.CommonConfig.MQTT.MqttURL)
	opts.SetClientID("agent_"+uuid.NewV4().String())
	opts.SetKeepAlive(10 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(9 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetUsername(config.CommonConfig.MQTT.CUser)
	opts.SetPassword(config.CommonConfig.MQTT.CPassword)
	logs.Debug(config.CommonConfig.MQTT.CUser)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	ec,err:=etcd.NewEtcd()
	if err!=nil{
		return err
	}

	go func(){
		err:=ec.WatchDataDevice(func(k string,p bool){
			logs.Info(k,true)
			if p{
				if token := c.Subscribe(k, 0, nil); token.Wait() && token.Error() != nil {
					logs.Error(token.Error())
				}
			}else{
				if token := c.Unsubscribe(k); token.Wait() && token.Error() != nil {
					logs.Error(token.Error())
				}
			}
			
		})
		if err!=nil{
			logs.Error(err)
			panic(err)
		}
	}()

	topics,err:=ec.GetDataDevice()
	if err!=nil{
		return err
	}
	for _,t:=range topics{
		if token := c.Subscribe(t, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return token.Error()
		}
	}
	return nil
}