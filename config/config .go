package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	DBName   string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}

var AppConfig Config

func InitConfig() (Config, error) {

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "mydb")
	viper.SetDefault("DB_PORT", "5432")

	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")

	viper.SetConfigFile("app.env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %v", err))
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}

	return AppConfig, nil
}
