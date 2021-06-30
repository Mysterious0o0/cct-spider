package config

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
)

func govConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/gov",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func minConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/min",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

var MinistryCf = minConfig()
