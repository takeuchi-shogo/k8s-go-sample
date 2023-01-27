package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	// database
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
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
