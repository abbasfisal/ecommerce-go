package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppPort    string
	AppVersion string
	AppDebug   bool
	AppName    string
	Db         Db
}

type Db struct {
	Driver   string
	Host     string
	Name     string
	Port     string
	UserName string
	Password string
}

func configViper() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func NewConfig() *Config {

	configViper()

	return &Config{
		AppPort:    viper.GetString("APP_PORT"),
		AppVersion: viper.GetString("APP_VERSION"),
		AppDebug:   viper.GetBool("APP_DEBUG"),
		AppName:    viper.GetString("APP_NAME"),

		Db: Db{
			Driver:   viper.GetString("DB_DRIVER"),
			Host:     viper.GetString("DB_HOST"),
			Name:     viper.GetString("DB_NAME"),
			Port:     viper.GetString("DB_PORT"),
			UserName: viper.GetString("DB_USERNAME"),
			Password: viper.GetString("DB_PASSWORD"),
		},
	}
}
