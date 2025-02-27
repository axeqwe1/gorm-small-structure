package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string
	DBSource string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read config", err)
	}

	return &Config{
		DBDriver: viper.GetString("DB_DRIVER"),
		DBSource: viper.GetString("DB_SOURCE"),
	}

}
