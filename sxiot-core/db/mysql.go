package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hashwing/sxiot/sxiot-core/config"
)

type AdminUser struct {
	UserID       string    `xorm:"admin_id" json:"admin_id"`
	UserRole     int    `xorm:"admin_role" json:"admin_role"`
	UserAccount  string `xorm:"admin_account" json:"admin_account"`
	UserPassword string `xorm:"admin_password" json:"-"`
	UserAlias    string `xorm:"admin_alias" json:"admin_alias"`
	UserEamil    string `xorm:"admin_email" json:"admin_email"`
	UserPhone    string `xorm:"admin_phone" josn:"admin_phone"`
}

type PersonUser struct {
	UserID       string    `xorm:"user_id" json:"user_id"`
	UserAccount  string `xorm:"user_account" json:"user_account"`
	UserPassword string `xorm:"user_password" json:"-"`
	UserAlias    string `xorm:"user_alias" json:"user_alias"`
	UserEmail    string `xorm:"user_email" json:"user_email"`
	UserPhone    string `xorm:"user_phone" json:"user_phone"`
}

type DeviceBrand struct {
	ID       string    `xorm:"brand_id" json:"brand_id"`
	Name   string `xorm:"brand_name" json:"brand_name"`
	Type string `xorm:"brand_type" json:"brand_type"`
	Metadata   string `xorm:"brand_metadata" json:"brand_metadata"`
}

type DeviceGateway struct {
	ID    string    `xorm:"gateway_id" json:"gateway_id"`
	AdminID string `xorm:"admin_id" json:"admin_id"`
	Name  string `xorm:"gateway_name" json:"gateway_name"`
}

type Device struct {
	ID    string    `xorm:"device_id" json:"device_id"`
	AdminID string `xorm:"admin_id" json:"admin_id"`
	BrandID string    `xorm:"brand_id" json:"brand_id"`
	Name  string `xorm:"device_alias" json:"device_alias"`
	Unit  string `xorm:"device_unit" json:"device_unit"`
}

type PersonDevice struct {
	ID string  `xorm:"id" json:"id"`
	UserID    string    `xorm:"user_id" json:"person_device_user"`
	DeviceID      string    `xorm:"gateway_id" json:"device_id"`
	Alias   string `xorm:"device_alias" json:"device_alias"`
	Status bool `xorm:"-" json:"device_status"`
}

type Message struct {
	ID      int       `xorm:"meaasge_id" json:"message_id"`
	UserID    int       `xorm:"user_id" json:"user_id"`
	Title   string    `xorm:"message_title" json:"message_title"`
	Content string    `xorm:"message_content" json:"massage_content"`
	Created time.Time `xorm:"created" json:"created"`
}

type News struct {
	ID      int       `xorm:"news_id" json:"news_id"`
	Title   string    `xorm:"news_title" json:"news_title"`
	Content string    `xorm:"news_content" json:"news_content"`
	Created time.Time `xorm:"created" json:"created"`
}

var MysqlDB *xorm.Engine

func NewDB() error {
	var err error
	url:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.CommonConfig.Mysql.UserName,
		config.CommonConfig.Mysql.Password,
		config.CommonConfig.Mysql.URL,
		config.CommonConfig.Mysql.DbName,
	)
	MysqlDB, err = xorm.NewEngine("mysql", url)
	return err
}
