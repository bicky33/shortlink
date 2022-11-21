package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	AppPort int `mapstructure:"APP_PORT"`
}

func LoadConfigApp(path string) *ConfigApp {
	var config ConfigApp
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigType("env")
	vp.SetConfigName(".env")

	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf("err %v", err.Error())
	}

	if err := vp.Unmarshal(&config); err != nil {
		log.Fatalln("unable to marshal env to viper")
	}
	return &config
}
