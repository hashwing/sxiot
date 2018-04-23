package db

func AddPersonDevice( device *PersonDevice)error{
	res,err:=MysqlDB.Table("sxiot_user_device").Where("device_id=?",device.DeviceID).Get(device)
	if res{
		return nil
	}
	_,err =MysqlDB.Table("sxiot_user_device").Insert(device)
	return err
}

func FindPersonDevices(userID string)([]PersonDevice,error){
	var devcies []PersonDevice
	err:=MysqlDB.Table("sxiot_user_device").Where("user_id=?",userID).Find(&devcies)
	return devcies,err
}

// AuthPersonDevice auth device
func AuthPersonDevice(id,gatewayID string)(bool){
	var device PersonDevice
	res,_:=MysqlDB.Table("sxiot_user_device").Where("gateway_id=? and id=?",gatewayID,id).Get(&device)
	return res
}

func UpdatePersonDevice(device *PersonDevice)error{
	_,err:=MysqlDB.Table("sxiot_user_device").Where("id=?",device.ID).Update(device)
	return err
}

func DelPersonDevice(deviceID string)error{
	device:=new(Device) 
	_,err:=MysqlDB.Table("sxiot_user_device").Where("device_id=?",deviceID).Delete(device)
	return err
}