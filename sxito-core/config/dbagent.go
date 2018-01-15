package config

const (
	DBAGENT_SERVICE_NAME="sxito-dbagent"
	DBAGENT_SERVICE_DESC="sxito-dbagent"
	DBAGENT_LOG_PATH="/var/log/sxito/sxito-dbagent"
)

type DBAgentService struct {
	MqttURL string `ini:"mqtt_url"`
	Topic string `ini:"mqtt_topic"`
	ClientID string `ini:"client_id"`
}