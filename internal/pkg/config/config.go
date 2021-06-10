package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	configType string
	configName string
	configPath string
}

func (c *Config) NewConfig() (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigType(c.configType)
	v.SetConfigName(c.configName)
	v.AddConfigPath(c.configPath)
	err = v.ReadInConfig()
	return
}
