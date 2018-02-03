package config

const (
	UAPI_SERVICE_NAME="sxiot-uapi"
	UAPI_SERVICE_DESC="sxiot-uapi"
	UAPI_LOG_PATH="/var/log/sxiot/sxiot-uapi"
)

type PlatformConfig struct {
	JwtSecret string `ini:"jwt_secret"`
	WebPort   int    `ini:"web_port"`
}