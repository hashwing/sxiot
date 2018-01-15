package config

const (
	CONFIG_PATH="/etc/sxito/sxito.conf"
)


type Influxdb struct{
	URL string `ini:"url"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	DbName string `ini:"db_name"`
}

type Mysql struct {
	URL string `ini:"url"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	DbName string `ini:"db_name"`
}