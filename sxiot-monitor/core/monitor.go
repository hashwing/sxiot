package core

import (
	"net"
	"strconv"
	"strings"
	"time"
	"errors"

	gpcpu "github.com/shirou/gopsutil/cpu"
	gpdisk "github.com/shirou/gopsutil/disk"
	gpmem "github.com/shirou/gopsutil/mem"
	gpnet "github.com/shirou/gopsutil/net"
	"github.com/astaxie/beego/logs"
	"github.com/hashwing/sxiot/sxiot-core/config"
)

var host string
var preNet PreNetData

func GetLocalIPStr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("not found ip")
}

func GetLocalIP() error {
	ip, err := GetLocalIPStr()
	if err != nil {
		return err
	}
	host = ip
	return nil
}

//获取CPU信息
func GetCPUInfo() error {
	data, err := gpcpu.Percent(0, false)
	if err != nil {
		return err
	}
	tag := make(map[string]string)
	tag["host"] = host
	for i, value := range data {
		field := make(map[string]interface{})
		field["UsePercentage"] = value
		tag["core"] = "cpu-" + strconv.Itoa(i)
		err=EndPointAdd("cpu", tag, field)
		return err
	}
	return nil
}

type PreNetData struct {
	Data []gpnet.IOCountersStat
	Time int64
}

//获取网络信息
func GetNetInfo() error {
	data, err := gpnet.IOCounters(true)
	if err != nil {
		return err
	}
	t := time.Now().Unix()
	logs.Info(len(data))
	if preNet.Time != 0 {
		for i, value := range data {
			iface, err := net.InterfaceByName(value.Name)
			if err != nil {
				continue
			}
			if iface.Flags&net.FlagLoopback == net.FlagLoopback {
				continue
			}
			if iface.Flags&net.FlagUp == 0 {
				continue
			}
			tag := make(map[string]string)
			tag["host"] = host
			field := make(map[string]interface{})
			field["ReceivedPerSec"] = int64(value.BytesRecv-preNet.Data[i].BytesRecv) / (t - preNet.Time)
			field["SentPerSec"] = int64(value.BytesSent-preNet.Data[i].BytesSent) / (t - preNet.Time)
			tag["core"] = strings.Replace(value.Name, " ", "", -1)
			err=EndPointAdd("network", tag, field)
			if err!=nil{
				return err
			}
		}
	}
	preNet.Data = data
	preNet.Time = t
	return nil
}

//逻辑磁盘信息
func GetDiskInfo() error {
	part, err := gpdisk.Partitions(false)
	if err != nil {
		return err
	}

	for _, value := range part {
		tag := make(map[string]string)
		tag["host"] = host
		data, err := gpdisk.Usage(value.Mountpoint)
		if err != nil {
			return err
		}
		field := make(map[string]interface{})
		field["UsePercentage"] = data.UsedPercent
		field["Size"] = int64(data.Total)
		field["FreeSpace"] = int64(data.Free)
		tag["core"] = strings.Replace(value.Mountpoint, " ", "", -1)
		err=EndPointAdd("disk", tag, field)
		if err != nil {
			return err
		}
	}
	return nil
}

//获取磁盘读写速率信息
func GetDiskRWInfo() error {
	part, err := gpdisk.Partitions(true)
	if err != nil {
		return err
	}
	tag := make(map[string]string)
	tag["host"] = host
	var DiskReadBytesPerSec, DiskWriteBytesPerSec uint64
	for _, value := range part {
		data, err := gpdisk.IOCounters(value.Device)
		if err != nil {
			return err
		}
		DiskReadBytesPerSec += data[value.Device].ReadBytes
		DiskWriteBytesPerSec += data[value.Device].WriteBytes
	}
	field := make(map[string]interface{})
	field["DiskReadPerSec"] = int64(DiskReadBytesPerSec)
	field["DiskWritePerSec"] = int64(DiskWriteBytesPerSec)
	tag["core"] = "diskrw"
	err=EndPointAdd("diskrw", tag, field)
	return err
}

//获取内存信息
func GetMemoryInfo() error {
	data, err := gpmem.VirtualMemory()
	if err != nil {
		return err
	}
	tag := make(map[string]string)
	tag["host"] = host
	field := make(map[string]interface{})
	field["UsePercentage"] = data.UsedPercent
	tag["core"] = "memory"
	err=EndPointAdd("memory", tag, field)
	return err
}

func RunMonitor(){
	err:=NewInflux()
	if err!=nil{
		logs.Error(err)
		panic(err)
	}
	err=GetLocalIP()
	if err!=nil{
		logs.Error(err)
		panic(err)
	}
	for {
		err := GetCPUInfo()
		if err != nil {
			logs.Error(err)
		}
		err = GetMemoryInfo()
		if err != nil {
			logs.Error(err)
		}
		err = GetDiskInfo()
		if err != nil {
			logs.Error(err)
		}
		err = GetNetInfo()
		if err != nil {
			logs.Error(err)
		}
		err = GetDiskRWInfo()
		if err != nil {
			logs.Error(err)
		}
		err = EndPointWrite()
		if err != nil {
			logs.Error(err)
		}
		time.Sleep(time.Duration(config.CommonConfig.Monitor.Interval) * time.Second)
	}
}

