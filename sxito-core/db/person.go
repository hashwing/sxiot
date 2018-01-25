package db

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func AuthUser(account, password string) (bool,*PersonUser) {
	h := md5.New()
	io.WriteString(h, password)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	var user PersonUser
	res, err := MysqlDB.Table("sxito_user").Where("user_account=? and user_password=?", account, passwdMD5).Get(&user)
	if !res || err != nil {
		return false,nil
	}
	return true,&user
}

func AddUser(user *PersonUser)error{
	h := md5.New()
	io.WriteString(h, user.UserPassword)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	user.UserPassword=passwdMD5
	_,err:=MysqlDB.Table("sxito_user").Insert(user)
	return err
}
