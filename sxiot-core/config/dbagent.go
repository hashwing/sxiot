package config

const (
	DBAGENT_SERVICE_NAME="sxiot-dbagent"
	DBAGENT_SERVICE_DESC="sxiot-dbagent"
	DBAGENT_LOG_PATH="/var/log/sxiot/sxiot-dbagent"
)

type DBAgentService struct {
	MqttURL string `ini:"mqtt_url"`
	Topic string `ini:"mqtt_topic"`
	ClientID string `ini:"client_id"`
}