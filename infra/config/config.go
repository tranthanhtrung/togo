package config

import (
	"github.com/spf13/viper"
)

// Config reprepare for type config app
type Config struct {
	Database struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"database"`
	Task struct {
		MaxInTheDay int `mapstructure:"max_in_the_day"`
	} `mapstructure:"task"`
}

// LoadConfig load config from file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
