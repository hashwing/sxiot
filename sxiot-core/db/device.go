package db

func AddDevice(device *Device)error{
	_,err:=MysqlDB.Table("sxiot_device").Insert(device)
	return err
}

func FindDevices(adminID string)([]Device,error){
	var devcies []Device
	err:=MysqlDB.Table("sxiot_device").Where("admin_id=?",adminID).Find(&devcies)
	return devcies,err
}

func GetDevice(deviceID string)(*Device,error){
	var device Device
	res,err:=MysqlDB.Table("sxiot_device").Where("device_id=?",deviceID).Get(&device)
	if !res{
		return nil,nil
	}
	return &device,err
}


func UpdateDevice(device *Device)error{
	_,err:=MysqlDB.Table("sxiot_device").Where("device_id=?",device.ID).Update(device)
	return err
}

func DelDevice(deviceID string)error{
	device:=new(Device) 
	_,err:=MysqlDB.Table("sxiot_device").Where("device_id=?",deviceID).Delete(device)
	return err
}