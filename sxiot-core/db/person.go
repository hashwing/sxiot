package db

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
)

func AuthUser(account, password string) (bool, *PersonUser) {
	h := md5.New()
	io.WriteString(h, password)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	var user PersonUser
	res, err := MysqlDB.Table("sxiot_user").Where("user_account=? and user_password=?", account, passwdMD5).Get(&user)
	if !res || err != nil {
		return false, nil
	}
	return true, &user
}

func AuthOpenID(openID string) (bool, *PersonUser) {
	var user PersonUser
	res, err := MysqlDB.Table("sxiot_user").Where("user_openid=?", openID).Get(&user)
	if !res || err != nil {
		return false, nil
	}
	return true, &user
}

func UpdateUserOpenID(openID, account string) error {
	user := new(PersonUser)
	user.OpenID = openID
	_, err := MysqlDB.Table("sxiot_user").Where("user_account=?", account).Update(user)
	return err
}

func AddUser(user *PersonUser) error {
	h := md5.New()
	io.WriteString(h, user.UserPassword)
	passwdMD5 := hex.EncodeToString(h.Sum(nil))
	user.UserPassword = passwdMD5
	has, err := MysqlDB.Table("sxiot_user").Where("user_account=?", user.UserAccount).Exist()
	if has {
		return errors.New("user account is exist")
	}
	if err != nil {
		return err
	}
	_, err = MysqlDB.Table("sxiot_user").Insert(user)
	return err
}

// GetUser get user info
func GetUser(uid string) (*PersonUser, error) {
	var user PersonUser
	_, err := MysqlDB.Table("sxiot_user").Where("user_id=?", uid).Get(&user)
	return &user, err
}

func CountUser() (int64, error) {
	user := new(PersonUser)
	return MysqlDB.Table("sxiot_user").Count(user)
}
