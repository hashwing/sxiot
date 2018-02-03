package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
	"github.com/go-ini/ini"
)
var CommonConfig = &SxiotConfig{}

type SxiotConfig struct {
	DBAgent DBAgentService `ini:"dbagent"`
	Influxdb Influxdb `ini:"influxdb"`
	Mysql Mysql `ini:"mysql"`
	Platform PlatformConfig `ini:"platform"`
	Monitor Monitor `ini:"monitor"`
}

func NewCommonConfig() error {
	err := LoadConfig(CONFIG_PATH, CommonConfig)
	return err
}

func LoadConfig(file string, settings interface{}) error {

	if file != "" {

		absConfPath, err := filepath.Abs(file)
		if err != nil {
			logs.Debug(err)
			return err
		}

		if err := ini.MapTo(settings, absConfPath); err != nil {
			logs.Debug(err)
			return err
		}

		return nil
	}

	return errors.New("file is nil")
}

func WriteConfig(file string, settings interface{}) error {

	if file != "" {

		cfg := ini.Empty()
		err := ini.ReflectFrom(cfg, settings)
		if err != nil {
			return err
		}

		if file == "-" {
			cfg.WriteTo(os.Stdout)
		} else {
			err = cfg.SaveTo(file)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("file is nil")
}
