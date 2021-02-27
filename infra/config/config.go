package config

import (
	"github.com/spf13/viper"
)

// Config reprepare for type config app
type Config struct {
	Database struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"database"`
	Todo struct {
		MaxTaskes int `mapstructure:"max_taskes_in_a_day"`
	} `mapstructure:"todo"`
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
