package config

const (
	MONITOR_SERVICE_NAME="sxito-monitor"
	MONITOR_SERVICE_DESC="sxito-monitor"
	MONITOR_LOG_PATH="/var/log/sxito/sxito-monitor"
)

type Monitor struct {
	Interval int `ini:"interval"`
}