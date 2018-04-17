package etcd

import (
	"strings"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/hashwing/sxiot/sxiot-core/db"
)

// GetDataDevice get data device id
func (ec *Client)GetDataDevice()(ids []string,err error){
	grsp,err :=ec.Get("data/",clientv3.WithPrefix())
	ids = make([]string,0)
	for _,v:=range grsp.Kvs{
		ids=append(ids,string(v.Value))
	}
	return
}

func (ec *Client)WatchDataDevice(call func(k string,p bool))(err error){
	wch := ec.Watch("data/", clientv3.WithPrefix())
	for {
		select {
		case c := <-wch:
			for _, e := range c.Events {
				fmt.Println(string(e.Kv.Key))
				if e.Type == clientv3.EventTypeDelete {
					call(strings.Replace(string(e.Kv.Key),"data/","",-1),false)
				}
				if e.Type == clientv3.EventTypePut {
					call(string(e.Kv.Value),true)
				}
			}
		}
	}
}

func (ec *Client)PutDataDevices()error{
	devices,err:=db.FindDevicesByB("data")
	if err!=nil{
		return err
	}
	for _,v:=range devices{
		_,err:=ec.Put("data/"+v.ID,v.ID)
		if err!=nil{
			return err
		}
	}
	return nil
}

func (ec *Client)PutDataDevice(key string)error{
		_,err:=ec.Put("data/"+key,key)
		return err
}

func (ec *Client)DeleteDataDevice(key string)error{
	_,err:=ec.Delete("data/"+key)
	return err
}