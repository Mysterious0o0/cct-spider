package main

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/miit"
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

func main() {
	minV := minConfig()
	miitUrl := minV.GetString("工业和信息化部")
	miit.GetFirstUrl(miitUrl)
}