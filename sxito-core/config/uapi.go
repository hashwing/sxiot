package config

const (
	UAPI_SERVICE_NAME="sxito-uapi"
	UAPI_SERVICE_DESC="sxito-uapi"
	UAPI_LOG_PATH="/var/log/sxito/sxito-uapi"
)

type PlatformConfig struct {
	JwtSecret string `ini:"jwt_secret"`
	WebPort   int    `ini:"web_port"`
}