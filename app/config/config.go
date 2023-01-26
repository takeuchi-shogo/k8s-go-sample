package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func NewConfig(path string) *Config {

	viper.SetConfigFile(path)
	viper.SetConfigFile("app.env")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read config: ", err)
	}

	c := &Config{}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("env file unmarshal: ", err)
	}

	return c
}
