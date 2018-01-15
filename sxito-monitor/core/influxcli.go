package core

import (
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/hashwing/sxito/sxito-core/config"
)
var endpoint client.Client
var batchpoint client.BatchPoints

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

func EndPointAdd(table string,tags map[string]string,fields map[string]interface{})error{
	metric,err := client.NewPoint(table, tags, fields, time.Now())
	if err!=nil{
		return err
	}
	batchpoint.AddPoint(metric)
	return nil
}

func EndPointWrite() error {
	err:=endpoint.Write(batchpoint)
	return err
}