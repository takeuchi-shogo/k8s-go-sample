package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ApplicationName string

	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	// database
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	Cors struct {
		AllowOringins []string
	}

	Jwt struct {
		TokenExpireAt int
		SecretKey     string
	}
}

func NewConfig(path string) *Config {

	viper.SetConfigFile(path)
	viper.SetConfigFile("development.env")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read config: ", err)
	}

	c := &Config{}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("env file unmarshal: ", err)
	}

	if c.Environment == "devlopment" {
		log.Print("\n==================================================\n")
		log.Print("\n\nCurrently, it's a development environment!!!!!!\n\n")
		log.Print("\n==================================================\n")
	}
	// kubernetes環境の時は下記を反映したい
	// 他にもいい書き方はあるはず
	if os.Getenv("ENV") == "kubernetes" {
		c.DBHost = os.Getenv("DB_HOST")
	}

	c.ApplicationName = "sample-app"

	c.Cors.AllowOringins = []string{"http://localhost:3000"}

	return c
}
