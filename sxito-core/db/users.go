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
	res, err := MysqlDB.Table("sxito_admin").Where("admin_account=? and admin_password=?", account, passwdMD5).Get(&user)
	if !res || err != nil {
		return false,nil
	}
	return true,&user
}
