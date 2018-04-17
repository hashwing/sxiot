package config

const (
	CONFIG_PATH="/etc/sxiot/sxiot.conf"
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

type Emqtt struct {
	URL  	 	string 	`ini:"emq_url"`
	MqttURL 	string 	`ini:"mqtt_url"`
	User 	 	string 	`ini:"emq_user"`
	Password 	string 	`ini:"emq_passwd"`
	CUser		string	`ini:"mqtt_username"`
	CPassword	string	`ini:"mqtt_password"`
}

type Etcd struct {
	URL  	 []string `ini:"url"`
}