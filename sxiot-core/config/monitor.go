package config

const (
	MONITOR_SERVICE_NAME="sxiot-monitor"
	MONITOR_SERVICE_DESC="sxiot-monitor"
	MONITOR_LOG_PATH="/var/log/sxiot/sxiot-monitor"
)

type Monitor struct {
	Interval int `ini:"interval"`
}