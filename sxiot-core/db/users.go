package db

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func AuthAdmin(account, password string) (bool,*AdminUser) {
	h := md5.New()
	io.WriteString(h, password)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	var user AdminUser
	res, err := MysqlDB.Table("sxiot_admin").Where("admin_account=? and admin_password=?", account, passwdMD5).Get(&user)
	if !res || err != nil {
		return false,nil
	}
	return true,&user
}

func AddAdmin(user *AdminUser)error{
	h := md5.New()
	io.WriteString(h, user.UserPassword)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	user.UserPassword=passwdMD5
	_,err:=MysqlDB.Table("sxiot_admin").Insert(user)
	return err
}

func UpdateAdmin(user *AdminUser)error{
	if user.UserPassword!=""{
		h := md5.New()
		io.WriteString(h, user.UserPassword)
		passwdMD5 := hex.EncodeToString(h.Sum(nil))
		user.UserPassword=passwdMD5
	}
	_,err:=MysqlDB.Table("sxiot_admin").Where("admin_id=?",user.UserID).Update(user)
	return err
}

func FindAdmins()(*[]AdminUser,error){
	var user []AdminUser
	err:=MysqlDB.Table("sxiot_admin").Find(&user)
	if err!=nil{
		return nil,err
	}
	return &user,nil
}


func GetAdmin(uid string)(*AdminUser,error){
	var user AdminUser
	res,err:=MysqlDB.Table("sxiot_admin").Where("admin_id=?",uid).Get(&user)
	if err!=nil{
		return nil,err
	}
	if !res{
		return nil,err
	}
	return &user,nil
}


func DelAdmin(uid string)error{
	user:=new(AdminUser)
	_,err:=MysqlDB.Table("sxiot_admin").Where("admin_id=?",uid).Delete(user)
	return err
}