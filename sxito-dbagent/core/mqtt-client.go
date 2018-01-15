package core

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/hashwing/sxito/sxito-core/config"
)

var endpoint client.Client
var batchpoint client.BatchPoints

type DBData struct{
	GatewayID string `json:"gateway_id"`
	DeviceID string `json:"device_id"`
	Data interface{} `json:"data"`
}

var f mqtt.MessageHandler = func(cli mqtt.Client, msg mqtt.Message) {
	dbData :=&DBData{}
	err:=json.Unmarshal(msg.Payload(),dbData)
	if err!=nil{
		logs.Error(err)
		return
	}
	if endpoint != nil {
		tag := map[string]string{
			"client_id":dbData.DeviceID,
		}
		field := map[string]interface{}{
			"data":dbData.Data,
		}
		metric, err:= client.NewPoint(dbData.GatewayID, tag, field, time.Now())
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
	opts :=mqtt.NewClientOptions().AddBroker(config.CommonConfig.DBAgent.MqttURL).SetClientID(config.CommonConfig.DBAgent.ClientID)
	opts.SetKeepAlive(10 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(9 * time.Second)
	opts.SetAutoReconnect(true)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	if token := c.Subscribe(config.CommonConfig.DBAgent.Topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return token.Error()
	}

	return nil
}