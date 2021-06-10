package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConfigType string
	ConfigName string
	ConfigPath string
}

func (c *Config) NewConfig() (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigType(c.ConfigType)
	v.SetConfigName(c.ConfigName)
	v.AddConfigPath(c.ConfigPath)
	err = v.ReadInConfig()
	return
}
